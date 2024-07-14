package main

import (
	"log"
	"user_service/api"
	"user_service/api/handler"
	"user_service/genproto/authentication"
	l "user_service/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	hand := Newhandler()
	router := api.Router(hand)
	log.Println("server is running")
	log.Fatal(router.Run(":8085"))
}
func Newhandler() *handler.Handler {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	authservice := authentication.NewAuthenticationClient(conn)
	loggers := l.NewLogger()
	return &handler.Handler{Auth: authservice, Log: loggers}
}
