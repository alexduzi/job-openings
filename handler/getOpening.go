package handler

import (
	"fmt"
	"net/http"

	"github.com/alexduzi/job-openings/model"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		opening := model.Opening{}
		if err := db.First(&opening, id).Error; err != nil {
			sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
			return
		}
		sendSuccess(ctx, "get-opening", opening)
		return
	}

	openings := []model.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("openings with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "get-openings", openings)
}
