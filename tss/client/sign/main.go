package main

import (
	"context"
	"crypto/elliptic"
	"encoding/base64"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bnb-chain/tss-lib/v2/tss"
	pb "github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/utils"
	"google.golang.org/grpc"
)

func init() {
	tss.RegisterCurve("elliptic.p256Curve", elliptic.P256())
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <share_data_filename>")
	}

	filename := os.Args[1]

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMPCServiceClient(conn)

	signCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Extract session ID from filename
	sessionID := strings.TrimSuffix(strings.TrimPrefix(filename, "share_data_"), ".txt")

	shareData, err := utils.LoadFromJSON(filename)
	if err != nil {
		log.Fatalf("failed to load share data: %v", err)
	}

	encryptedShare, err := base64.StdEncoding.DecodeString(string(shareData))
	if err != nil {
		log.Fatalf("failed to decode base64 share: %v", err)
	}

	msgHash := []byte("hello")

	res, err := client.NotifyAction(signCtx, &pb.ActionRequest{
		SessionId: sessionID,
		Parties:   []uint32{1, 2, 3},
		Threshold: 2,
		MsgHash:   msgHash,
		ShareData: encryptedShare,
		Action:    pb.Action_INIT_SIGN,
	})
	if err != nil {
		log.Fatalf("failed to notify signing action: %v", err)
	}

	log.Printf("Sign init response: %v", res)
}
