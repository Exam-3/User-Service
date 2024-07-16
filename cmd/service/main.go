package main

import (
	"fmt"
	"log"
	"net"
	"user_service/pkg"

	"google.golang.org/grpc"

	"user_service/config"
	"user_service/service"
	"user_service/storage/postgres"

	pb1 "user_service/genproto/authentication"
	pb "user_service/genproto/user"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Starting server...")
	cfg := config.Load()
	fmt.Println(cfg.Server.USER_PORT)
	lis, err := net.Listen("tcp", cfg.Server.USER_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	itemClient := pkg.CreateItemClient(*cfg)
	userService := service.NewUserService(db, itemClient)
	authService := service.NewAuthService(db)
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userService)
	pb1.RegisterAuthenticationServer(server, authService)

	log.Printf("server listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
