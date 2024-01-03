package main

import (
	"fmt"
	"net"

	"github.com/go-exe/alerts"
	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	grpcServer := grpc.NewServer()

	alert := alerts.New()

	alerts.RegisterAlertServiceServer(grpcServer, alert)

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}

}
