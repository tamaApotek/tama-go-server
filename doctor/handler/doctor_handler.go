package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/doctor"
	"github.com/tamaApotek/tama-go-server/helpers"
	"github.com/tamaApotek/tama-go-server/models"
)

type doctorHandler struct {
	doctorUsecase doctor.Usecase
}

func NewDoctorHandler(r *gin.Engine, doctorUsecase doctor.Usecase) {
	handler := &doctorHandler{doctorUsecase}

	r.POST("/", handler.Add)
}

func (d *doctorHandler) Add(c *gin.Context) {
	ctx, _ := context.WithTimeout(c, 3*time.Second)

	doctor := new(models.Doctor)
	if err := c.ShouldBindJSON(doctor); err != nil {
		helpers.HandleErrorResponse(c, err)
		return
	}

	id, err := d.doctorUsecase.Add(ctx, *doctor)
	if err != nil {
		helpers.HandleErrorResponse(c, err)
	}

	helpers.HandleSuccessResponse(c, id)
}
