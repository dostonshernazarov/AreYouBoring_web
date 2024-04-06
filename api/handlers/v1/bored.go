package v1

import (
	"context"
	"github/NotBoring/api/handlers/models"
	boredapi "github/NotBoring/bored_api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	// models "github/NotBoring/api/handlers/models"
	l "github/NotBoring/pkg/logger"
)

// GetActivity gets activity
// @Summary GetActivity
// @Description Api for getting random activity
// @Tags bored
// @Accept json
// @Produce json
// @Success 200 {object} models.BoredRes
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/activity [get]
func (h *handlerV1) GetActivity(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true


	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.BoredService().GetTip(ctx, &boredapi.Bored{
		Ok: "ok",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.BoredRes{
		Activity: response.Resultt,
	})
}
