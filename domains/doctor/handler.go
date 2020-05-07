package doctor

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/domains/apperror"
	"github.com/tamaApotek/tama-go-server/domains/query"
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
		err = apperror.New("Invalid data", apperror.ErrInvalid, err)
		query.HandleErrorResponse(c, err)
		return
	}

	id, err := d.doctorUsecase.Add(ctx, doctor)
	if err != nil {
		query.HandleErrorResponse(c, err)
	}

	query.HandleSuccessResponse(c, id)
}
