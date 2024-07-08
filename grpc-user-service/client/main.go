package main

import (
	"context"
	"log"
	"time"

	pb "grpc-user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// GetUser example
	getUserRequest := &pb.GetUserRequest{Id: 1}
	getUserResponse, err := client.GetUser(ctx, getUserRequest)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("GetUser Response: %v", getUserResponse.User)

	// GetUsers example
	getUsersRequest := &pb.GetUsersRequest{Ids: []int32{1, 2}}
	getUsersResponse, err := client.GetUsers(ctx, getUsersRequest)
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Printf("GetUsers Response: %v", getUsersResponse.Users)

	// SearchUser example
	searchUserRequest := &pb.SearchUserRequest{City: "LA", Phone: 1234567890, Married: true}
	searchUserResponse, err := client.SearchUser(ctx, searchUserRequest)
	if err != nil {
		log.Fatalf("could not search user: %v", err)
	}
	log.Printf("SearchUser Response: %v", searchUserResponse.Users)
}
