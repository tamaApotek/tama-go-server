package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/config"
	_doctorHandler "github.com/tamaApotek/tama-go-server/doctor/handler"
	_doctorUsecase "github.com/tamaApotek/tama-go-server/doctor/usecase"
	_userHandler "github.com/tamaApotek/tama-go-server/user/handler"
	_userRepo "github.com/tamaApotek/tama-go-server/user/repository"
	_userUsecase "github.com/tamaApotek/tama-go-server/user/usecase"
)

func main() {
	client, err := config.InitMongo("mongodb://localhost/27107")
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
