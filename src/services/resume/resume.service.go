package resumeService

import (
	"bhakti-buana-api/src/configs"
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	resumeRequest "bhakti-buana-api/src/requests/resume"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Resume Store Service
/*
 * @param context *gin.Context
 * @param request *resumeRequest.S_StoreRequest
 * @returns *models.Resumes
 */
func Store(context *gin.Context, request *resumeRequest.S_StoreRequest) *models.Resumes {
	var resume models.Resumes

	dir := "./public/resumes/"
	helpers.CheckDir(dir)

	newFileName := fmt.Sprintf("%v-%s", time.Now().Unix(), request.File.Filename)

	if err := context.SaveUploadedFile(request.File, dir+newFileName); err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
	}

	parsedUrl, _ := url.Parse(configs.AppConfig().BASE_URL + "/public/resumes/" + newFileName)

	resume = models.Resumes{
		Title:      request.Title,
		URL:        parsedUrl.String(),
		Downloaded: 0,
		Status:     models.RESUME_STATUS_INACTIVE,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	result, err := database.Resumes.InsertOne(context, &resume)
	if err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	resume.ID = result.InsertedID.(primitive.ObjectID)

	return &resume
}
