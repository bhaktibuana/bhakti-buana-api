package resumeController

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	resumeRequest "bhakti-buana-api/src/requests/resume"
	resumeResult "bhakti-buana-api/src/results/resume"
	resumeService "bhakti-buana-api/src/services/resume"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Resume Store Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Store(context *gin.Context) {
	request := resumeRequest.Store(context)
	if request == nil {
		return
	}

	resume := resumeService.Store(context, request)
	if resume == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, resumeResult.Store(resume))
}
