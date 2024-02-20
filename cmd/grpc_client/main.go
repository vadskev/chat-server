package main

import (
	"context"
	"github.com/fatih/color"
	desc "github.com/vadskev/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed toconnect to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewChatV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var userNames = []string{"Test1", "Test2"}
	r, err := c.Create(ctx, &desc.CreateRequest{Usernames: userNames})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	chat := r
	log.Printf(color.RedString("The chat was created with an ID: "), color.GreenString("%+v", chat.Id))
}
