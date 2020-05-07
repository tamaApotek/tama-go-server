package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/config"
	"github.com/tamaApotek/tama-go-server/domains/doctor"
	"github.com/tamaApotek/tama-go-server/domains/user"
)

func main() {
	client, err := config.InitMongo("mongodb://localhost/27107")
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("tama")

	userRepo := user.NewRepoMongo(db)
	doctorRepo := doctor.NewRepoMongo(db)

	userUsecase := user.NewUsecase(userRepo)
	doctorUsecase := doctor.NewUsecase(doctorRepo, userRepo)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())

	user.NewHandler(r.Group("/users"), userUsecase)
	doctor.NewHandler(r.Group("/doctors"), doctorUsecase)

	log.Fatal(r.Run())
}
