package schedule

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/bootstrap"
)

type handler struct {
	*bootstrap.App

	schUC Usecase
}

func NewHandler(r *gin.RouterGroup, schUC Usecase, app *bootstrap.App) {
	h := &handler{App: app, schUC: schUC}

	r.POST("/", h.Create)
	r.GET("/doctors/:doctor_id", h.FindByDoctorID)
}

func (h *handler) Create(c *gin.Context) {
	sch := new(Schedule)

	if err := c.ShouldBindJSON(sch); err != nil {
		h.Handler.HandleErrorResponse(c.Writer, err)
		return
	}

	ctx := context.Background()
	schID, err := h.schUC.Create(ctx, sch)
	if err != nil {
		h.Handler.HandleErrorResponse(c.Writer, err)
		return
	}

	h.Handler.HandleSuccessResponse(c.Writer, schID)
}

func (h *handler) FindByDoctorID(c *gin.Context) {
	doctorID := c.Param("doctor_id")

	ctx := context.Background()
	schs, err := h.schUC.FindByDoctorID(ctx, doctorID)
	if err != nil {
		h.Handler.HandleErrorResponse(c.Writer, err)
		return
	}

	h.Handler.HandleSuccessResponse(c.Writer, schs)
}
