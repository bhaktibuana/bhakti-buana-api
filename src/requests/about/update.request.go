package aboutRequest

import (
	"bhakti-buana-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_UpdateParam struct {
	ID string `form:"id"`
}

type S_AboutLocation struct {
	City     string `json:"city" binding:"required"`
	Province string `json:"province" binding:"required"`
	Country  string `json:"country" binding:"required"`
}

type S_UpdateBody struct {
	NickName string          `json:"nick_name" binding:"required"`
	Role     string          `json:"role" binding:"required"`
	Email    string          `json:"email" binding:"required"`
	Summary  string          `json:"summary" binding:"required"`
	Location S_AboutLocation `json:"location" binding:"required"`
}

type S_UpdateRequest struct {
	S_UpdateParam
	S_UpdateBody
}

// About Update Request
/*
 * @param context *gin.Context
 * @returns *S_UpdateRequest
 */
func Update(context *gin.Context) *S_UpdateRequest {
	var param S_UpdateParam
	var body S_UpdateBody
	var request S_UpdateRequest

	param.ID = context.Param("id")

	if param.ID == ":id" {
		helpers.HttpResponse("Param 'id' is required", http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	request = S_UpdateRequest{
		S_UpdateParam: param,
		S_UpdateBody:  body,
	}

	return &request
}
