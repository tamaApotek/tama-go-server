package doctor

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/delivery"
)

type doctorHandler struct {
	doctorUsecase Usecase
}

// NewHandler construct "/doctors" route handler
func NewHandler(r *gin.RouterGroup, delivery delivery.Delivery, doctorUsecase Usecase) {
	h := &doctorHandler{doctorUsecase}

	r.POST("/", delivery.Handle(h.Add))
}

func (d *doctorHandler) Add(c *gin.Context) (interface{}, error) {
	ctx, _ := context.WithTimeout(c, 3*time.Second)

	var doctor Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		err = fmt.Errorf("Invalid data: %w", err)
		return nil, err
	}

	id, err := d.doctorUsecase.Add(ctx, &doctor)
	if err != nil {
		return nil, err
	}

	return id, nil
}
