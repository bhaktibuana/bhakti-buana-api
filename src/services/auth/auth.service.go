package authService

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	authRequest "bhakti-buana-api/src/requests/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_LoginRequest
 * @returns *models.Users
 */
func Login(context *gin.Context, request *authRequest.S_LoginRequest) *models.Users {
	var user models.Users

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	normalizedEmail := strings.ToLower(request.Email)

	filter := bson.M{"email": normalizedEmail, "password": request.Password}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.WRONG_MAIL_PASS, http.StatusBadRequest, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if !user.DeletedAt.IsZero() {
		helpers.HttpResponse(constants.WRONG_MAIL_PASS, http.StatusBadRequest, context, nil)
		return nil
	}

	if user.Status == models.USER_STATUS_UNVERIFIED {
		helpers.HttpResponse(constants.UNVERIFIED_MAIL, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}

// Me Service
/*
 * @param context *gin.Context
 * @returns *models.Users
 */
func Me(context *gin.Context) *models.Users {
	var user models.Users

	id, err := helpers.GetSelfID(context)
	if err != nil {
		helpers.HttpResponse(constants.INVALID_USER, http.StatusBadRequest, context, nil)
		return nil
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == models.USER_STATUS_UNVERIFIED {
		helpers.HttpResponse(constants.UNVERIFIED_MAIL, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}
