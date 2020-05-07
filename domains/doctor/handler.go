package doctor

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

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

	var doctor Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		err = fmt.Errorf("Invalid data: %w", err)
		query.HandleErrorResponse(c, err)
		return
	}

	fmt.Printf("[DEBUG] doctor %+v\n", doctor)

	id, err := d.doctorUsecase.Add(ctx, &doctor)
	if err != nil {
		query.HandleErrorResponse(c, err)
		return
	}

	query.HandleSuccessResponse(c, id)
}
