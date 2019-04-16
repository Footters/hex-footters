package auth

import (
	"fmt"

	"google.golang.org/grpc"
)

// NewAuthServiceProviderConnection func
func NewAuthServiceProviderConnection() *grpc.ClientConn {
	conn, err := grpc.Dial("auth:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error to create GRPC Dial")
		panic(err)
	}

	fmt.Println("State connection GRPC", conn.GetState())
	return conn
}
