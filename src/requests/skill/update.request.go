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

type S_UpdateParam struct {
	ID string `form:"id"`
}

type S_UpdateFormData struct {
	Code      string                `form:"code" binding:"required"`
	Order     int64                 `form:"order" binding:"required"`
	Name      string                `form:"name" binding:"required"`
	SourceUrl string                `form:"source_url" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
}

type S_UpdateRequest struct {
	S_UpdateParam
	S_UpdateFormData
}

// Skill Update Request
/*
 * @param context *gin.Context
 * @returns *S_UpdateRequest
 */
func Update(context *gin.Context) *S_UpdateRequest {
	var request S_UpdateRequest

	request.S_UpdateParam.ID = context.Param("id")

	if request.ID == ":id" {
		helpers.HttpResponse("Param 'id' is required", http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	maxSize := 2 << 20

	if err := context.ShouldBind(&request.S_UpdateFormData); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if request.Image != nil {
		if ext := strings.ToLower(filepath.Ext(request.Image.Filename)); !(ext == ".jpg" || ext == ".jpeg" || ext == ".png") {
			helpers.HttpResponse(constants.FILE_ACCEPT_IMAGE, http.StatusUnprocessableEntity, context, nil)
			return nil
		}

		if request.Image.Size > int64(maxSize) {
			helpers.HttpResponse(constants.FILE_LIMIT_2MB, http.StatusRequestEntityTooLarge, context, nil)
			return nil
		}
	}

	return &request
}
