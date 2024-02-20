package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	desc "github.com/vadskev/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create: username: %v", req.Usernames)

	id := genRandomID()
	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	log.Printf("Delete: id: %v", req.Id)
	return &desc.DeleteResponse{}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	log.Printf("SendMessage: From: %v, Text: %v, Timestamp: %v ", req.From, req.Text, req.Timestamp)
	return &desc.SendMessageResponse{}, nil
}

func genRandomID() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
