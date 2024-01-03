package main

import (
	"context"
	"fmt"

	"github.com/go-exe/alerts"
	"google.golang.org/grpc"
)

func main() {

	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	defer connection.Close()

	alert := alerts.NewAlertServiceClient(connection)

	msg := alerts.AlertRequest{
		Body: "Hello World",
	}

	response, err := alert.Send(context.Background(), &msg)
	if err != nil {
		fmt.Println("send failed, err:", err)
		return
	}

	fmt.Println("Response: ", response.Response)
}
