package pkg

import (
	"errors"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	
	"user_service/config"

	pb "user_service/genproto/item"
)
func CreateItemClient(cfg config.Config) pb.ItemServiceClient {
	conn, err := grpc.NewClient(cfg.Server.ITEM_PORT,

		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pb.NewItemServiceClient(conn)
}
