package resumeRequest

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_UpdateStatusParam struct {
	ID string `form:"id"`
}

type S_UpdateStatusBody struct {
	Status string `json:"status" binding:"required"`
}

type S_UpdateStatusRequest struct {
	S_UpdateStatusParam
	S_UpdateStatusBody
}

// Resume UpdateStatus Request
/*
 * @param context *gin.Context
 * @returns *S_UpdateStatusRequest
 */
func UpdateStatus(context *gin.Context) *S_UpdateStatusRequest {
	var param S_UpdateStatusParam
	var body S_UpdateStatusBody
	var request S_UpdateStatusRequest

	param.ID = context.Param("id")

	if param.ID == ":id" {
		helpers.HttpResponse("Param 'id' is required", http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if !(body.Status == models.RESUME_STATUS_ACTIVE || body.Status == models.RESUME_STATUS_INACTIVE) {
		helpers.HttpResponse(constants.INVALID_RESUME_STATUS, http.StatusBadRequest, context, nil)
		return nil
	}

	request = S_UpdateStatusRequest{
		S_UpdateStatusParam: param,
		S_UpdateStatusBody:  body,
	}

	return &request
}
