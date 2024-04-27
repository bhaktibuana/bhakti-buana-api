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
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type S_IndexServiceResult struct {
	Data       []models.Resumes
	Pagination helpers.S_Pagination
}

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
		return nil
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

// Resume Index Service
/*
 * @param context *gin.Context
 * @param request *resumeRequest.S_IndexRequest
 * @returns *S_IndexServiceResult
 */
func Index(context *gin.Context, request *resumeRequest.S_IndexRequest) *S_IndexServiceResult {
	var result S_IndexServiceResult

	paginateResult := helpers.Paginate(context, request.S_PaginationRequest, database.Resumes, reflect.TypeOf(models.Resumes{}))

	resumes := make([]models.Resumes, 0)
	for _, item := range paginateResult.Data {
		if resume, ok := item.(models.Resumes); ok {
			resumes = append(resumes, resume)
		}
	}

	result = S_IndexServiceResult{
		Data:       resumes,
		Pagination: paginateResult.Pagination,
	}

	return &result
}

// Resume UpdateStatus Service
/*
 * @param context *gin.Context
 * @param request *resumeRequest.S_UpdateStatusRequest
 * @returns *models.Resumes
 */
func UpdateStatus(context *gin.Context, request *resumeRequest.S_UpdateStatusRequest) *models.Resumes {
	var resume models.Resumes

	_id, _ := primitive.ObjectIDFromHex(request.ID)
	filter := bson.M{"_id": _id}

	payload := bson.M{
		"$set": bson.M{
			"status":     request.Status,
			"updated_at": time.Now(),
		},
	}

	if request.Status == models.RESUME_STATUS_ACTIVE {
		filterMany := bson.M{
			"deleted_at": bson.M{"$eq": nil},
		}

		payloadMany := bson.M{
			"$set": bson.M{
				"status":     models.RESUME_STATUS_INACTIVE,
				"updated_at": time.Now(),
			},
		}

		if _, err := database.Resumes.UpdateMany(context, filterMany, payloadMany); err != nil {
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}
	}

	result, err := database.Resumes.UpdateOne(context, filter, payload)
	if err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	if result.ModifiedCount == 0 {
		helpers.HttpResponse(constants.ID_NOT_FOUND, http.StatusNotFound, context, nil)
		return nil
	}

	if err := database.Resumes.FindOne(context, filter).Decode(&resume); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}
	}

	var aboutPayload primitive.M

	aboutPayload = bson.M{
		"$set": bson.M{
			"resume":     models.S_AboutResume{},
			"updated_at": time.Now(),
		},
	}

	if request.Status == models.RESUME_STATUS_ACTIVE {
		aboutPayload = bson.M{
			"$set": bson.M{
				"resume": models.S_AboutResume{
					ID:  resume.ID,
					URL: resume.URL,
				},
				"updated_at": time.Now(),
			},
		}
	}

	database.About.UpdateMany(context, bson.M{}, aboutPayload)

	return &resume
}