package main

import (
	"context"
	"testing"

	pb "grpc-user-service/proto"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	s := &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		},
	}

	req := &pb.GetUserRequest{Id: 1}
	res, err := s.GetUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, req.Id, res.User.Id)
}

func TestGetUsers(t *testing.T) {
	s := &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	req := &pb.GetUsersRequest{Ids: []int32{1, 2}}
	res, err := s.GetUsers(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Users, 2)
}

func TestSearchUser(t *testing.T) {
	s := &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	// Adjust the search criteria to match the test data
	req := &pb.SearchUserRequest{City: "LA", Phone: 1234567890, Married: true}
	res, err := s.SearchUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Users, 1)
	assert.Equal(t, "LA", res.Users[0].City)
	assert.Equal(t, int64(1234567890), res.Users[0].Phone)
	assert.True(t, res.Users[0].Married)
}
