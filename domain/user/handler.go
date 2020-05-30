package user

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/common"
	"github.com/tamaApotek/tama-go-server/delivery"
)

type Handler struct {
	userUsecase Usecase
}

// NewHandler will initiate user/ resources endpoint
func NewHandler(r *gin.RouterGroup, userUsecase Usecase, delivery delivery.Delivery) {
	handler := &Handler{userUsecase}

	r.GET("/users/:user_id", handler.FindByID)
}

func (u *Handler) FindByID(c *gin.Context) {
	uid := c.Param("user_id")

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	user, err := u.userUsecase.FindByID(ctx, uid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, user)
}

func (u *Handler) Create(c *gin.Context) {
	user := new(User)
	if err := c.ShouldBindJSON(user); err != nil {
		delivery.HandleErrorResponse(c.Writer, common.ErrInvalidBody)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	uid, err := u.userUsecase.Create(ctx, user)
	if err != nil {
		delivery.HandleErrorResponse(c.Writer, err)
		return
	}

	delivery.HandleSuccessResponse(c.Writer, uid)
}
