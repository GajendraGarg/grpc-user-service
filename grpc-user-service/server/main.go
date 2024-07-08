package main

import (
	"context"
	"log"
	"net"

	pb "grpc-user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []pb.User
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	if err := ValidateUserID(req.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.UserResponse{User: &user}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.UsersResponse, error) {
	if err := ValidateGetUsersRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	var result []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				result = append(result, &user)
				break
			}
		}
	}
	return &pb.UsersResponse{Users: result}, nil
}

// SearchUser searches for users based on the given criteria.
func (s *server) SearchUser(ctx context.Context, req *pb.SearchUserRequest) (*pb.UsersResponse, error) {
	if err := ValidateSearchUserRequest(req); err != nil {
		return nil, err
	}

	var matchedUsers []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == false || user.Married == req.Married) {
			matchedUsers = append(matchedUsers, &user)
		}
	}

	return &pb.UsersResponse{Users: matchedUsers}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			// Add more sample users here
		},
	})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
