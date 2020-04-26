package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	_doctorHandler "github.com/tamaApotek/tama-go-server/doctor/handler"
	_doctorUsecase "github.com/tamaApotek/tama-go-server/doctor/usecase"
	_userHandler "github.com/tamaApotek/tama-go-server/user/handler"
	_userRepo "github.com/tamaApotek/tama-go-server/user/repository"
	_userUsecase "github.com/tamaApotek/tama-go-server/user/usecase"
)

func init() {

}

func main() {
	clientOpt := options.Client().ApplyURI("mongodb://localhost/27107")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("tama")

	userRepo := _userRepo.NewUserMongo(db)

	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	doctorUsecase := _doctorUsecase.NewDoctorUsecase(userRepo)

	r := gin.Default()

	_userHandler.NewUserHandler(r.Group("/users"), userUsecase)
	_doctorHandler.NewDoctorHandler(r.Group("/doctors"), doctorUsecase)

	log.Fatal(r.Run())
}
