package main

import (
	"errors"
	"fmt"
	pb "grpc-user-service/proto"
)

// ValidateUserID checks if the user ID is valid
func ValidateUserID(id int32) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	return nil
}

// ValidateGetUsersRequest checks if the request contains valid user IDs
func ValidateGetUsersRequest(req *pb.GetUsersRequest) error {
	if len(req.Ids) == 0 {
		return errors.New("no user IDs provided")
	}
	for _, id := range req.Ids {
		if err := ValidateUserID(id); err != nil {
			return err
		}
	}
	return nil
}

// ValidateSearchUserRequest validates the SearchUserRequest.
func ValidateSearchUserRequest(req *pb.SearchUserRequest) error {
	if req.City == "" && req.Phone == 0 && req.Married == false {
		return fmt.Errorf("at least one search criteria must be specified")
	}
	return nil
}
