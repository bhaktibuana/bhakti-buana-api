package aboutController

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	aboutRequest "bhakti-buana-api/src/requests/about"
	aboutResult "bhakti-buana-api/src/results/about"
	aboutService "bhakti-buana-api/src/services/about"
	"net/http"

	"github.com/gin-gonic/gin"
)

// About Update Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Update(context *gin.Context) {
	request := aboutRequest.Update(context)
	if request == nil {
		return
	}

	about := aboutService.Update(context, request)
	if about == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, aboutResult.Update(about))
}

// About UpdatePhoto Controller
/*
 * @param context *gin.Context
 * @returns
 */
func UpdatePhoto(context *gin.Context) {
	request := aboutRequest.UpdatePhoto(context)
	if request == nil {
		return
	}

	about := aboutService.UpdatePhoto(context, request)
	if about == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, aboutResult.UpdatePhoto(about))
}

// About Show Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Show(context *gin.Context) {
	about := aboutService.Show(context)
	if about == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, aboutResult.Show(about))
}
