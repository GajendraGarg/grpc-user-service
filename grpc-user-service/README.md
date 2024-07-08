# gRPC User Service in Golang

This project implements a gRPC service for managing user details with search functionality. The service includes endpoints for fetching user details based on a user ID, retrieving a list of user details based on a list of user IDs, and searching for user details based on specific criteria.

## Prerequisites

- Go 1.22.4 or higher
- Docker
- Docker Compose
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Project Structure
```
.
├── grpc-user-service/
├── ├── proto/
├── │ └── user.proto
├── ├── server/
├── │ ├── Dockerfile
├── │ ├── main.go
├── │ ├── server.go
├── │ ├── server_test.go
├── │ └── validation.go
├── ├── docker-compose.yml
├── ├── go.mod
└── └── go.sum
```

## Setup

### 1. Install Go

Install Go 1.22.4 or higher from the [official Go website](https://golang.org/dl/).

### 2. Install Protocol Buffers

Follow the instructions to install Protocol Buffers for your operating system from the [official documentation](https://grpc.io/docs/protoc-installation/).

### 3. Install `protoc-gen-go` and `protoc-gen-go-grpc` Plugins

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure that $GOPATH/bin is in your PATH.

### 4. Generate Go Code from Proto File
Navigate to the proto directory and run:

```
protoc --go_out=. --go-grpc_out=. user.proto
```

### 5. Build and Run with Docker
To build and run the application using Docker, follow these steps:
Build the Docker images:
```
docker-compose build
```
Run the Docker containers:
```
docker-compose up
```

The server will start and listen on localhost:50051.


### Running Server Locally (Without Docker)
If you prefer to run the server locally without Docker, follow these steps:

Navigate to Server Directory: Go to the server directory of your project.
```
cd server
go run main.go
```
This will start the server locally, and it will also listen on localhost:50051.

### 7. Running the Client
To access the gRPC service endpoints using the client, you can either run the client from within the Docker container or locally:

Within Docker: The client will automatically run and interact with the server when you use docker-compose up.

Locally: Run the client locally while the server is running in Docker.

```
cd client
go run .
```

Code Overview

User Model

```
message User {
    int32 id = 1;
    string fname = 2;
    string city = 3;
    int64 phone = 4;
    double height = 5;
    bool married = 6;
}

```

gRPC Service

```
service UserService {
    rpc GetUser (GetUserRequest) returns (UserResponse);
    rpc GetUsers (GetUsersRequest) returns (UsersResponse);
    rpc SearchUser (SearchUserRequest) returns (UsersResponse);
}
```

Validation Functions
```
ValidateUserID(id int32) error
ValidateGetUsersRequest(req *pb.GetUsersRequest) error
ValidateSearchUserRequest(req *pb.SearchUserRequest) error
```
Server Endpoints
```
GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error)
GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.UsersResponse, error)
SearchUser(ctx context.Context, req *pb.SearchUserRequest) (*pb.UsersResponse, error)
```

## Edge Cases Handled

- Invalid user ID in `GetUser` and `GetUsers`
- Empty list of user IDs in `GetUsers`
- Invalid phone number in `SearchUser`
- User not found scenarios
