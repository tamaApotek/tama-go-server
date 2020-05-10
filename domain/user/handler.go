package user

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase Usecase
}

// NewHandler will initiate user/ resources endpoint
func NewHandler(r *gin.RouterGroup, userUsecase Usecase) {
	handler := &Handler{userUsecase}

	r.GET("/users/:user_id", handler.FindByID)
}

func (u *Handler) FindByID(c *gin.Context) {
	uid := c.Param("user_id")

	ctx, _ := context.WithTimeout(c, 3*time.Second)

	user, err := u.userUsecase.FindByID(ctx, uid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, user)
}
