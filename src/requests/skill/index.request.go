package skillRequest

import (
	"bhakti-buana-api/src/helpers"

	"github.com/gin-gonic/gin"
)

type S_IndexRequest struct {
	helpers.S_PaginationRequest
}

// Skill Index Request
/*
 * @param context *gin.Context
 * @returns *S_IndexRequest
 */
func Index(context *gin.Context) *S_IndexRequest {
	var request S_IndexRequest

	paginationRequest := helpers.PaginationRequest(context)
	if paginationRequest == nil {
		return nil
	}

	request.S_PaginationRequest = *paginationRequest

	return &request
}
