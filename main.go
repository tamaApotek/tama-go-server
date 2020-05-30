package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/tamaApotek/tama-go-server/config"
	"github.com/tamaApotek/tama-go-server/delivery"
	"github.com/tamaApotek/tama-go-server/domain/doctor"
	"github.com/tamaApotek/tama-go-server/domain/queue"
	"github.com/tamaApotek/tama-go-server/domain/user"
	"github.com/tamaApotek/tama-go-server/internal"
)

func main() {
	client, err := config.InitMongo("mongodb://localhost/27107")
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("tama")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())

	log.SetOutput(gin.DefaultWriter)

	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatal("Invalid validator")
	}

	internal.InitValidation(validator)

	d := delivery.Delivery{}

	userRepo := user.NewRepoMongo(db)
	doctorRepo := doctor.NewRepoMongo(db)
	queueRepo := queue.NewRepoMongo(db)

	tz, _ := time.LoadLocation("Asia/Jakarta")

	userUsecase := user.NewUsecase(userRepo)
	doctorUsecase := doctor.NewUsecase(doctorRepo, userRepo)
	queueUsecase := queue.NewUsecase(tz, queueRepo, doctorRepo)

	user.NewHandler(r.Group("/users"), userUsecase)
	doctor.NewHandler(r.Group("/doctors"), d, doctorUsecase)
	queue.NewHandler(r.Group("/queues"), d, queueUsecase)

	log.Fatal(r.Run())
}
