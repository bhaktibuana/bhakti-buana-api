package skillRequest

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
	Code      string                `form:"code" binding:"required"`
	Order     int64                 `form:"order"`
	Name      string                `form:"name" binding:"required"`
	SourceUrl string                `form:"source_url" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
}

// Skill Store Request
/*
 * @param context *gin.Context
 * @returns *S_StoreRequest
 */
func Store(context *gin.Context) *S_StoreRequest {
	var request S_StoreRequest

	maxSize := 2 << 20

	if err := context.ShouldBind(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if ext := strings.ToLower(filepath.Ext(request.Image.Filename)); !(ext == ".jpg" || ext == ".jpeg" || ext == ".png") {
		helpers.HttpResponse(constants.FILE_ACCEPT_IMAGE, http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if request.Image.Size > int64(maxSize) {
		helpers.HttpResponse(constants.FILE_LIMIT_2MB, http.StatusRequestEntityTooLarge, context, nil)
		return nil
	}

	return &request
}
