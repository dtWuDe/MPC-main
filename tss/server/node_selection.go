package server

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type NodeInfo struct {
	ID          uint32
	Address     string
	Load        float64   // Current load (0.0 - 1.0)
	LastSeen    time.Time // Last heartbeat
	Status      string    // "online", "offline", "maintenance"
	ActiveTasks int       // Number of active signing sessions
}

type NodeSelectionStrategy int

const (
	RoundRobin NodeSelectionStrategy = iota
	LoadBased
	AvailabilityBased
	Random
)

type NodeSelector struct {
	nodes    map[uint32]*NodeInfo
	mutex    sync.RWMutex
	strategy NodeSelectionStrategy
}

func NewNodeSelector(strategy NodeSelectionStrategy) *NodeSelector {
	return &NodeSelector{
		nodes:    make(map[uint32]*NodeInfo),
		strategy: strategy,
	}
}

// RegisterNode adds or updates a node in the selector
func (ns *NodeSelector) RegisterNode(id uint32, address string) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()

	ns.nodes[id] = &NodeInfo{
		ID:       id,
		Address:  address,
		Load:     0.0,
		LastSeen: time.Now(),
		Status:   "online",
	}
}

// UpdateNodeStatus updates node status and load
func (ns *NodeSelector) UpdateNodeStatus(id uint32, load float64, status string) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()

	if node, exists := ns.nodes[id]; exists {
		node.Load = load
		node.Status = status
		node.LastSeen = time.Now()
	}
}

// SelectNodes chooses which nodes should participate in TSS operation
func (ns *NodeSelector) SelectNodes(totalNodes, threshold int, sessionID string) ([]uint32, error) {
	ns.mutex.RLock()
	defer ns.mutex.RUnlock()

	if len(ns.nodes) < threshold {
		return nil, fmt.Errorf("insufficient nodes available: %d < %d", len(ns.nodes), threshold)
	}

	switch ns.strategy {
	case RoundRobin:
		return ns.selectNodesRoundRobin(totalNodes, threshold, sessionID)
	case LoadBased:
		return ns.selectNodesByLoad(threshold)
	case AvailabilityBased:
		return ns.selectNodesByAvailability(threshold)
	case Random:
		return ns.selectNodesRandom(threshold)
	default:
		return ns.selectNodesRoundRobin(totalNodes, threshold, sessionID)
	}
}

// Round-robin selection based on session ID
func (ns *NodeSelector) selectNodesRoundRobin(totalNodes, threshold int, sessionID string) ([]uint32, error) {
	// Create deterministic hash from session ID
	hash := sha256.Sum256([]byte(sessionID))
	startNode := int(hash[0]) % totalNodes

	parties := make([]uint32, threshold)
	for i := 0; i < threshold; i++ {
		nodeID := (startNode + i) % totalNodes
		parties[i] = uint32(nodeID + 1) // Node IDs start from 1
	}

	return parties, nil
}

// Load-based selection (prefer nodes with lower load)
func (ns *NodeSelector) selectNodesByLoad(threshold int) ([]uint32, error) {
	var availableNodes []*NodeInfo

	// Filter online nodes
	for _, node := range ns.nodes {
		if node.Status == "online" {
			availableNodes = append(availableNodes, node)
		}
	}

	if len(availableNodes) < threshold {
		return nil, fmt.Errorf("insufficient online nodes: %d < %d", len(availableNodes), threshold)
	}

	// Sort by load (ascending)
	sort.Slice(availableNodes, func(i, j int) bool {
		return availableNodes[i].Load < availableNodes[j].Load
	})

	// Select first 'threshold' nodes
	parties := make([]uint32, threshold)
	for i := 0; i < threshold; i++ {
		parties[i] = availableNodes[i].ID
	}

	return parties, nil
}

// Availability-based selection (prefer nodes with fewer active tasks)
func (ns *NodeSelector) selectNodesByAvailability(threshold int) ([]uint32, error) {
	var availableNodes []*NodeInfo

	// Filter online nodes
	for _, node := range ns.nodes {
		if node.Status == "online" {
			availableNodes = append(availableNodes, node)
		}
	}

	if len(availableNodes) < threshold {
		return nil, fmt.Errorf("insufficient online nodes: %d < %d", len(availableNodes), threshold)
	}

	// Sort by active tasks (ascending)
	sort.Slice(availableNodes, func(i, j int) bool {
		return availableNodes[i].ActiveTasks < availableNodes[j].ActiveTasks
	})

	// Select first 'threshold' nodes
	parties := make([]uint32, threshold)
	for i := 0; i < threshold; i++ {
		parties[i] = availableNodes[i].ID
	}

	return parties, nil
}

// Random selection from available nodes
func (ns *NodeSelector) selectNodesRandom(threshold int) ([]uint32, error) {
	var availableNodes []*NodeInfo

	// Filter online nodes
	for _, node := range ns.nodes {
		if node.Status == "online" {
			availableNodes = append(availableNodes, node)
		}
	}

	if len(availableNodes) < threshold {
		return nil, fmt.Errorf("insufficient online nodes: %d < %d", len(availableNodes), threshold)
	}

	// Random selection
	parties := make([]uint32, 0, threshold)
	rand.Seed(time.Now().UnixNano())

	for len(parties) < threshold && len(availableNodes) > 0 {
		index := rand.Intn(len(availableNodes))
		parties = append(parties, availableNodes[index].ID)
		availableNodes = append(availableNodes[:index], availableNodes[index+1:]...)
	}

	return parties, nil
}

// GetNodeInfo returns information about a specific node
func (ns *NodeSelector) GetNodeInfo(id uint32) (*NodeInfo, bool) {
	ns.mutex.RLock()
	defer ns.mutex.RUnlock()

	node, exists := ns.nodes[id]
	return node, exists
}

// GetAllNodes returns all registered nodes
func (ns *NodeSelector) GetAllNodes() map[uint32]*NodeInfo {
	ns.mutex.RLock()
	defer ns.mutex.RUnlock()

	result := make(map[uint32]*NodeInfo)
	for id, node := range ns.nodes {
		result[id] = &NodeInfo{
			ID:          node.ID,
			Address:     node.Address,
			Load:        node.Load,
			LastSeen:    node.LastSeen,
			Status:      node.Status,
			ActiveTasks: node.ActiveTasks,
		}
	}

	return result
}

// MarkNodeBusy increments active tasks for a node
func (ns *NodeSelector) MarkNodeBusy(id uint32) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()

	if node, exists := ns.nodes[id]; exists {
		node.ActiveTasks++
	}
}

// MarkNodeIdle decrements active tasks for a node
func (ns *NodeSelector) MarkNodeIdle(id uint32) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()

	if node, exists := ns.nodes[id]; exists && node.ActiveTasks > 0 {
		node.ActiveTasks--
	}
}
