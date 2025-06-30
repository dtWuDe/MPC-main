package main

import (
	"context"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	pb "github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/utils"
	"google.golang.org/grpc"
)

func init() {
	tss.RegisterCurve("elliptic.p256Curve", elliptic.P256())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMPCServiceClient(conn)
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	sessionID := uuid.NewString()
	notifyCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	_, err = client.NotifyAction(notifyCtx, &pb.ActionRequest{
		SessionId: sessionID,
		Parties:   []uint32{1, 2, 3},
		Threshold: 2,
		Action:    pb.Action_INIT_KEYGEN,
	})
	if err != nil {
		log.Fatalf("failed to notify keygen action: %v", err)
	}

	log.Println("Waiting for keygen result...")

	pubsub := redisClient.Subscribe(context.Background(), "keygen:"+sessionID)
	defer pubsub.Close()

	jsonPayload := waitForKeygenResult(pubsub.Channel())

	// Parse JSON để lấy phần "share_data"
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonPayload), &result); err != nil {
		log.Fatalf("failed to parse result JSON: %v", err)
	}

	shareData := result["share_data"].(string)

	err = utils.SaveToTextFile([]byte(shareData), fmt.Sprintf("share_data_%s.txt", sessionID))
	if err != nil {
		log.Fatalf("failed to save share data: %v", err)
	}

	log.Printf("Keygen complete. Share data saved to share_data_%s.txt\n", sessionID)
}

func waitForKeygenResult(ch <-chan *redis.Message) string {
	for msg := range ch {
		return msg.Payload
	}
	return ""
}
