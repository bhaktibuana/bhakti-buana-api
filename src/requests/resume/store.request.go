package resumeRequest

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type S_StoreRequest struct {
	Title string                `form:"title" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

// Resume Store Request
/*
 * @param context *gin.Context
 * @returns *S_StoreRequest
 */
func Store(context *gin.Context) *S_StoreRequest {
	var request S_StoreRequest

	maxSize := 5 << 20

	if err := context.ShouldBind(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if ext := strings.ToLower(filepath.Ext(request.File.Filename)); ext != ".pdf" {
		helpers.HttpResponse(constants.FILE_ACCEPT_PDF, http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if request.File.Size > int64(maxSize) {
		helpers.HttpResponse(constants.FILE_LIMIT_5MB, http.StatusRequestEntityTooLarge, context, nil)
		return nil
	}

	return &request
}
