package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"mpc/internal/model"
	"mpc/internal/repository"
	"mpc/pkg/errors"
	"mpc/pkg/logger"
	"strings"
	"time"

	"github.com/google/uuid"
)

type OrganizationService struct {
	orgRepo *repository.OrganizationRepository
}

func NewOrganizationService(orgRepo *repository.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		orgRepo: orgRepo,
	}
}

// CreateOrganization creates a new organization
func (s *OrganizationService) CreateOrganization(ctx context.Context, req *model.CreateOrganizationRequest) (*model.Organization, error) {
	// Generate unique API key prefix
	prefix, err := s.generateUniquePrefix()
	if err != nil {
		return nil, err
	}

	org := &model.Organization{
		ID:                     uuid.New(),
		Name:                   req.Name,
		Domain:                 req.Domain,
		APIKeyPrefix:           prefix,
		Status:                 "active",
		Plan:                   req.Plan,
		MaxUsers:               req.MaxUsers,
		MaxAPIRequestsPerMonth: req.MaxAPIRequestsPerMonth,
		Settings:               req.Settings,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	err = s.orgRepo.CreateOrganization(ctx, org)
	if err != nil {
		logger.Error("Failed to create organization", err)
		return nil, err
	}

	return org, nil
}

// GetOrganizationByID retrieves an organization by ID
func (s *OrganizationService) GetOrganizationByID(ctx context.Context, orgID uuid.UUID) (*model.Organization, error) {
	org, err := s.orgRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		logger.Error("Failed to get organization", err)
		return nil, err
	}
	if org.ID == uuid.Nil {
		return nil, errors.ErrOrganizationNotFound
	}
	return org, nil
}

// CreateAPIKey creates a new API key for an organization
func (s *OrganizationService) CreateAPIKey(ctx context.Context, orgID uuid.UUID, req *model.CreateAPIKeyRequest) (*model.APIKey, string, error) {
	// Generate API key
	apiKey := s.generateAPIKey()
	keyHash := s.hashAPIKey(apiKey)

	apiKeyData := &model.APIKey{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           req.Name,
		KeyHash:        keyHash,
		Permissions:    req.Permissions,
		Status:         "active",
		ExpiresAt:      req.ExpiresAt,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.orgRepo.CreateAPIKey(ctx, apiKeyData)
	if err != nil {
		logger.Error("Failed to create API key", err)
		return nil, "", err
	}

	return apiKeyData, apiKey, nil
}

// ValidateAPIKey validates an API key and returns organization and key data
func (s *OrganizationService) ValidateAPIKey(ctx context.Context, apiKey string) (*model.Organization, *model.APIKey, error) {
	keyHash := s.hashAPIKey(apiKey)

	apiKeyData, err := s.orgRepo.GetAPIKeyByHash(ctx, keyHash)
	if err != nil {
		logger.Error("Failed to get API key", err)
		return nil, nil, errors.ErrInvalidAPIKey
	}
	if apiKeyData.ID == uuid.Nil {
		return nil, nil, errors.ErrInvalidAPIKey
	}

	// Check if API key is active
	if apiKeyData.Status != "active" {
		return nil, nil, errors.ErrAPIKeyInactive
	}

	// Check expiration
	if apiKeyData.ExpiresAt != nil && time.Now().After(*apiKeyData.ExpiresAt) {
		return nil, nil, errors.ErrAPIKeyExpired
	}

	// Get organization
	org, err := s.orgRepo.GetOrganizationByID(ctx, apiKeyData.OrganizationID)
	if err != nil {
		logger.Error("Failed to get organization", err)
		return nil, nil, errors.ErrOrganizationNotFound
	}

	// Update last used
	go s.updateAPIKeyLastUsed(ctx, apiKeyData.ID)

	return org, apiKeyData, nil
}

// GetAPIKeysByOrganization gets all API keys for an organization
func (s *OrganizationService) GetAPIKeysByOrganization(ctx context.Context, orgID uuid.UUID) ([]*model.APIKey, error) {
	apiKeys, err := s.orgRepo.GetAPIKeysByOrganization(ctx, orgID)
	if err != nil {
		logger.Error("Failed to get API keys", err)
		return nil, err
	}
	return apiKeys, nil
}

// DeleteAPIKey deletes an API key
func (s *OrganizationService) DeleteAPIKey(ctx context.Context, orgID uuid.UUID, apiKeyID uuid.UUID) error {
	// Verify ownership
	apiKey, err := s.orgRepo.GetAPIKeyByID(ctx, apiKeyID)
	if err != nil {
		return err
	}
	if apiKey.OrganizationID != orgID {
		return errors.ErrInvalidAPIKey
	}

	err = s.orgRepo.DeleteAPIKey(ctx, apiKeyID)
	if err != nil {
		logger.Error("Failed to delete API key", err)
		return err
	}
	return nil
}

// CheckRateLimit checks if organization has exceeded rate limits
func (s *OrganizationService) CheckRateLimit(ctx context.Context, orgID uuid.UUID) error {
	// Get current month usage
	startOfMonth := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	usage, err := s.orgRepo.GetAPIUsageByOrganization(ctx, orgID, startOfMonth, time.Now())
	if err != nil {
		logger.Error("Failed to get API usage", err)
		return err
	}

	// Get organization limits
	org, err := s.orgRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		logger.Error("Failed to get organization", err)
		return err
	}

	if usage >= org.MaxAPIRequestsPerMonth {
		return errors.ErrRateLimitExceeded
	}

	return nil
}

// RecordAPIUsage records API usage for analytics
func (s *OrganizationService) RecordAPIUsage(ctx context.Context, orgID uuid.UUID, apiKeyID *uuid.UUID, endpoint, method string, statusCode int, responseTimeMs *int) error {
	usage := &model.APIUsage{
		ID:             uuid.New(),
		OrganizationID: orgID,
		APIKeyID:       apiKeyID,
		Endpoint:       endpoint,
		Method:         method,
		StatusCode:     statusCode,
		ResponseTimeMs: responseTimeMs,
		CreatedAt:      time.Now(),
	}

	err := s.orgRepo.CreateAPIUsage(ctx, usage)
	if err != nil {
		logger.Error("Failed to record API usage", err)
		return err
	}
	return nil
}

// GetUsageStats gets usage statistics for an organization
func (s *OrganizationService) GetUsageStats(ctx context.Context, orgID uuid.UUID) (*model.UsageStats, error) {
	// Get current month usage
	startOfMonth := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	monthlyUsage, err := s.orgRepo.GetAPIUsageByOrganization(ctx, orgID, startOfMonth, time.Now())
	if err != nil {
		logger.Error("Failed to get monthly usage", err)
		return nil, err
	}

	// Get total usage (last 12 months)
	startOfYear := time.Now().AddDate(-1, 0, 0)
	totalUsage, err := s.orgRepo.GetAPIUsageByOrganization(ctx, orgID, startOfYear, time.Now())
	if err != nil {
		logger.Error("Failed to get total usage", err)
		return nil, err
	}

	// Get average response time
	avgResponseTime, err := s.orgRepo.GetAverageResponseTime(ctx, orgID)
	if err != nil {
		logger.Error("Failed to get average response time", err)
		avgResponseTime = 0
	}

	return &model.UsageStats{
		TotalRequests:       totalUsage,
		RequestsThisMonth:   monthlyUsage,
		AverageResponseTime: avgResponseTime,
	}, nil
}

// Helper methods
func (s *OrganizationService) generateAPIKey() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return "sk_" + hex.EncodeToString(bytes)
}

func (s *OrganizationService) hashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}

func (s *OrganizationService) generateUniquePrefix() (string, error) {
	for i := 0; i < 10; i++ {
		bytes := make([]byte, 4)
		rand.Read(bytes)
		prefix := strings.ToUpper(hex.EncodeToString(bytes)[:8])

		// Check if prefix exists
		exists, err := s.orgRepo.CheckPrefixExists(context.Background(), prefix)
		if err != nil {
			continue
		}
		if !exists {
			return prefix, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique prefix")
}

func (s *OrganizationService) updateAPIKeyLastUsed(ctx context.Context, apiKeyID uuid.UUID) {
	now := time.Now()
	err := s.orgRepo.UpdateAPIKeyLastUsed(ctx, apiKeyID, &now)
	if err != nil {
		logger.Error("Failed to update API key last used", err)
	}
}

// GetOrganizations gets all organizations with pagination
func (s *OrganizationService) GetOrganizations(ctx context.Context, limit, offset int) ([]*model.Organization, error) {
	organizations, err := s.orgRepo.GetOrganizations(ctx, limit, offset)
	if err != nil {
		logger.Error("Failed to get organizations", err)
		return nil, err
	}
	return organizations, nil
}
