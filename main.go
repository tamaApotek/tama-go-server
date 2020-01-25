package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/user/handler"
	_userRepo "github.com/tamaApotek/tama-go-server/user/repository"
	_userUsecase "github.com/tamaApotek/tama-go-server/user/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	clientOpt := options.Client().ApplyURI("mongodb://localhost/27107")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	_ = client.Database("tama")

	userRepo := _userRepo.NewUserMongo(client)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)

	r := gin.Default()

	handler.NewUserHandler(r)

	log.Fatal(r.Run())
}
