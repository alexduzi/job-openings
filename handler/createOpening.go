package handler

import (
	"net/http"

	"github.com/alexduzi/job-openings/model"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error when creating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("body %+v", request)

	opening := model.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error when creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
