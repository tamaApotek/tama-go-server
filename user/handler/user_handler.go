package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/user"
)

// UserHandler represent http handler for user
type UserHandler struct {
	userUsecase user.Usecase
}

// NewUserHandler will initiate user/ resources endpoint
func NewUserHandler(r *gin.Engine, userUsecase user.Usecase) {
	handler := &UserHandler{userUsecase}

	r.GET("/users/:uid", handler.FindByUID)
}

func (u *UserHandler) FindByUID(c *gin.Context) {
	uid := c.Param("uid")

	ctx, _ := context.WithTimeout(c, 3*time.Second)

	user, err := u.userUsecase.FindByUID(ctx, uid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, user)
}
