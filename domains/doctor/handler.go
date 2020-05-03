package doctor

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/helpers"
)

type doctorHandler struct {
	doctorUsecase Usecase
}

func NewHandler(r *gin.RouterGroup, doctorUsecase Usecase) {
	handler := &doctorHandler{doctorUsecase}

	r.POST("/", handler.Add)
}

func (d *doctorHandler) Add(c *gin.Context) {
	ctx, _ := context.WithTimeout(c, 3*time.Second)

	doctor := new(Doctor)
	if err := c.ShouldBindJSON(doctor); err != nil {
		helpers.HandleErrorResponse(c, err)
		return
	}

	id, err := d.doctorUsecase.Add(ctx, doctor)
	if err != nil {
		helpers.HandleErrorResponse(c, err)
	}

	helpers.HandleSuccessResponse(c, id)
}
