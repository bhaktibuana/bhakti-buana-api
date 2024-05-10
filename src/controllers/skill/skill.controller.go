package skillController

import (
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/helpers"
	skillRequest "bhakti-buana-api/src/requests/skill"
	skillResult "bhakti-buana-api/src/results/skill"
	skillService "bhakti-buana-api/src/services/skill"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Skill Store Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Store(context *gin.Context) {
	request := skillRequest.Store(context)
	if request == nil {
		return
	}

	skill := skillService.Store(context, request)
	if skill == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusCreated, context, skillResult.Store(skill))
}

// Skill Index Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Index(context *gin.Context) {
	request := skillRequest.Index(context)
	if request == nil {
		return
	}

	skills := skillService.Index(context, request)
	if skills == nil {
		return
	}

	helpers.HttpResponsePaginate(constants.REQUEST_SUCCESS, http.StatusOK, context, skillResult.Index(skills.Data), skills.Pagination)
}

// Skill Update Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Update(context *gin.Context) {
	request := skillRequest.Update(context)
	if request == nil {
		return
	}

	skill := skillService.Update(context, request)
	if skill == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, skillResult.Update(skill))
}

// Skill Index Controller
/*
 * @param context *gin.Context
 * @returns
 */
func ShowPublic(context *gin.Context) {
	var request skillRequest.S_IndexRequest
	request.S_PaginationRequest.Page = 1
	request.S_PaginationRequest.PerPage = 1000
	request.S_PaginationRequest.SortNumber = 1

	skills := skillService.Index(context, &request)
	if skills == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, skillResult.Index(skills.Data))
}
