package authController

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	authRequest "bhakti-buana-api/src/requests/auth"
	authResult "bhakti-buana-api/src/results/auth"
	authService "bhakti-buana-api/src/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Login(context *gin.Context) {
	request := authRequest.Login(context)
	if request == nil {
		return
	}

	user := authService.Login(context, request)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.LOGIN_SUCCESS, http.StatusOK, context, authResult.Login(user))
}
