package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	users "./users_service"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	PORT = 9090
)

func main() {
	// create new gRPC server
	grpcServer := grpc.NewServer()

	// register and add user service to gRPC server
	users.RegisterUserServiceServer(grpcServer, &users.UserService{})

	// set logging for gRPC server
	grpclog.SetLogger(log.New(os.Stdout, "GRPC:", log.LstdFlags))

	// wrap gRPC server with the gRPC Web package
	wrappedServer := grpcweb.WrapServer(grpcServer)

	// handle requests to server with ServeHTTP handler in gRPC Web package
	handler := func(res http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(res, req)
	}

	// declare and configure http server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: http.HandlerFunc(handler),
	}

	// initialize server and listen for requests
	grpclog.Println("Starting server...")
	log.Fatalln(httpServer.ListenAndServe())
}
