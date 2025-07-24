package handler

import (
	"fmt"
	"net/http"

	"github.com/alexduzi/job-openings/model"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	opening := model.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	request.Merge(&opening)

	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error update opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
