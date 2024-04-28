package aboutRequest

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type S_UpdatePhotoParam struct {
	ID string `form:"id"`
}

type S_UpdatePhotoFormData struct {
	Photo *multipart.FileHeader `form:"photo" binding:"required"`
}

type S_UpdatePhotoRequest struct {
	S_UpdatePhotoParam
	S_UpdatePhotoFormData
}

// Resume UpdatePhoto Request
/*
 * @param context *gin.Context
 * @returns *S_UpdatePhotoRequest
 */
func UpdatePhoto(context *gin.Context) *S_UpdatePhotoRequest {
	var request S_UpdatePhotoRequest

	request.S_UpdatePhotoParam.ID = context.Param("id")

	if request.ID == ":id" {
		helpers.HttpResponse("Param 'id' is required", http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	maxSize := 2 << 20

	if err := context.ShouldBind(&request.S_UpdatePhotoFormData); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if ext := strings.ToLower(filepath.Ext(request.Photo.Filename)); !(ext == ".jpg" || ext == ".jpeg" || ext == ".png") {
		helpers.HttpResponse(constants.FILE_ACCEPT_IMAGE, http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if request.Photo.Size > int64(maxSize) {
		helpers.HttpResponse(constants.FILE_LIMIT_2MB, http.StatusRequestEntityTooLarge, context, nil)
		return nil
	}

	return &request
}
