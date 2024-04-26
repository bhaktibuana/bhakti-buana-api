package aboutService

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	aboutRequest "bhakti-buana-api/src/requests/about"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// About Update Service
/*
 * @param context *gin.Context
 * @param request *aboutRequest.S_UpdateRequest
 * @returns *models.About
 */
func Update(context *gin.Context, request *aboutRequest.S_UpdateRequest) *models.About {
	var about models.About

	_id, _ := primitive.ObjectIDFromHex(request.ID)
	filter := bson.M{"_id": _id}

	payload := bson.M{
		"$set": bson.M{
			"nick_name":  request.NickName,
			"role":       request.Role,
			"summary":    request.Summary,
			"email":      strings.ToLower(request.Email),
			"location":   request.Location,
			"updated_at": time.Now(),
		},
	}

	if _, err := database.About.UpdateOne(context, filter, payload); err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	about.ID = _id

	return &about
}
