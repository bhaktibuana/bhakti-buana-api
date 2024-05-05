package aboutService

import (
	"bhakti-buana-api/src/configs"
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	aboutRequest "bhakti-buana-api/src/requests/about"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// About UpdatePhoto Service
/*
 * @param context *gin.Context
 * @param request *aboutRequest.S_UpdatePhotoRequest
 * @returns *models.About
 */
func UpdatePhoto(context *gin.Context, request *aboutRequest.S_UpdatePhotoRequest) *models.About {
	var about models.About

	_id, _ := primitive.ObjectIDFromHex(request.ID)
	filter := bson.M{"_id": _id}

	dir := "./public/about/"
	helpers.CheckDir(dir)

	newFileName := fmt.Sprintf("%v-%s", time.Now().Unix(), request.Photo.Filename)

	if err := context.SaveUploadedFile(request.Photo, dir+newFileName); err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	if err := database.About.FindOne(context, filter).Decode(&about); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}
	}

	oldParts := strings.Split(about.Photo, "/")
	oldFileName := oldParts[len(oldParts)-1]
	if about.Photo != "" {
		os.Remove(dir + oldFileName)
	}

	parsedUrl, _ := url.Parse(configs.AppConfig().BASE_URL + "/public/about/" + newFileName)

	payload := bson.M{
		"$set": bson.M{
			"photo":      parsedUrl,
			"updated_at": time.Now(),
		},
	}

	if _, err := database.About.UpdateOne(context, filter, payload); err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	return &about
}

// About Show Service
/*
 * @param context *gin.Context
 * @returns *models.About
 */
func Show(context *gin.Context) *models.About {
	var about models.About

	filter := bson.M{
		"deleted_at": bson.M{"$eq": nil},
	}

	findOptions := options.Find().SetLimit(1)

	cursor, err := database.About.Find(context, filter, findOptions)
	if err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}
	defer cursor.Close(context)

	for cursor.Next(context) {
		if err := cursor.Decode(&about); err != nil {
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}
	}

	if err := cursor.Err(); err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	return &about
}
