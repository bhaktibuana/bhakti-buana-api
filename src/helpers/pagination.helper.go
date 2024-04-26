package helpers

import (
	"bhakti-buana-api/src/constants"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type S_PaginationRequest struct {
	Page       int64  `form:"page" binding:"required"`
	PerPage    int64  `form:"per_page" binding:"required"`
	Sort       string `form:"sort" binding:"required"`
	SortNumber int64  `json:"sort_number"`
}

type S_PaginateResult struct {
	Data       []interface{}
	Pagination S_Pagination
}

type S_Pagination struct {
	PerPage      int64            `json:"per_page"`
	CurrentPage  int64            `json:"current_page"`
	LastPage     int64            `json:"last_page"`
	FirstPageURL *string          `json:"first_page_url"`
	NextPageURL  *string          `json:"next_page_url"`
	PrevPageURL  *string          `json:"prev_page_url"`
	LastPageURL  *string          `json:"last_page_url"`
	Path         string           `json:"path"`
	From         int64            `json:"from"`
	To           int64            `json:"to"`
	Total        int64            `json:"total"`
	Links        []PaginationLink `json:"links"`
}

type PaginationLink struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

// Pagination Request
/*
 * @param context *gin.Context
 * @returns *S_PaginationRequest
 */
func PaginationRequest(context *gin.Context) *S_PaginationRequest {
	var request S_PaginationRequest

	if err := context.ShouldBindQuery(&request); err != nil {
		HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	if request.Sort == "asc" {
		request.SortNumber = 1
	} else if request.Sort == "desc" {
		request.SortNumber = -1
	} else {
		HttpResponse(constants.INVALID_SORT, http.StatusBadRequest, context, nil)
		return nil
	}

	return &request
}

// Paginate Options
/*
 * @param context *gin.Context
 * @param request S_PaginationRequest
 * @param collection *mongo.Collection
 * @param modelType reflect.Type
 * @param queryOptions ...interface{} (filter interface{}, sortBy string)
 * @returns *S_PaginateResult
 */
func Paginate(context *gin.Context, request S_PaginationRequest, collection *mongo.Collection, modelType reflect.Type, queryOptions ...interface{}) *S_PaginateResult {
	var result S_PaginateResult
	var resultSlicePtr interface{}
	var dataSlice []interface{}
	var filter interface{}

	resultSlicePtr = &dataSlice

	sortBy := "created_at"
	filter = bson.M{}

	if len(queryOptions) > 0 {
		filter = queryOptions[0]
		if val, ok := queryOptions[1].(string); ok {
			sortBy = val
		}
	}

	filter = bson.M{"$and": []bson.M{filter.(bson.M), {"deleted_at": bson.M{"$eq": nil}}}}

	opts := options.Find().
		SetSort(map[string]int{sortBy: int(request.SortNumber)}).
		SetLimit(request.PerPage).
		SetSkip((request.Page - 1) * request.PerPage)

	count, err := collection.CountDocuments(context, filter)
	if err != nil {
		HttpResponse(err.Error(), http.StatusInternalServerError, context, nil)
		return nil
	}

	cursor, err := collection.Find(context, filter, opts)
	if err != nil {
		HttpResponse(err.Error(), http.StatusInternalServerError, context, nil)
		return nil
	}
	defer cursor.Close(context)

	reflect.ValueOf(resultSlicePtr).Elem().Set(reflect.MakeSlice(reflect.TypeOf(resultSlicePtr).Elem(), 0, 0))

	for cursor.Next(context) {
		elem := reflect.New(modelType).Interface()
		if err := cursor.Decode(elem); err != nil {
			HttpResponse(err.Error(), http.StatusInternalServerError, context, nil)
			return nil
		}
		reflect.ValueOf(resultSlicePtr).Elem().Set(reflect.Append(reflect.ValueOf(resultSlicePtr).Elem(), reflect.ValueOf(elem).Elem()))
	}

	if err := cursor.Err(); err != nil {
		HttpResponse(err.Error(), http.StatusInternalServerError, context, nil)
		return nil
	}

	baseUrl, _ := context.Get("baseUrl")
	path := fmt.Sprintf("%s%s", baseUrl, context.Request.URL.Path)

	from := 1
	to := request.Page * request.PerPage
	totalPages := int(math.Ceil(float64(count) / float64(request.PerPage)))

	var firstPageUrl *string
	var prevPageUrl *string
	var nextPageUrl *string
	var lastPageUrl *string

	firstPageUrl = parseUrlRawQuery(context, path, int64(from), request.PerPage)
	lastPageUrl = parseUrlRawQuery(context, path, int64(totalPages), request.PerPage)

	if to > count {
		to = count
	}

	if count == 0 {
		from = 0
		firstPageUrl = nil
		lastPageUrl = nil
	}

	if request.Page-1 <= 0 {
		prevPageUrl = nil
	} else {
		prevPageUrl = parseUrlRawQuery(context, path, int64(request.Page-1), request.PerPage)
	}

	if request.Page+1 > int64(totalPages) {
		nextPageUrl = nil
	} else {
		nextPageUrl = parseUrlRawQuery(context, path, int64(request.Page+1), request.PerPage)
	}

	var links []PaginationLink
	for i := 1; i <= totalPages; i++ {
		url := parseUrlRawQuery(context, path, int64(i), request.PerPage)
		label := strconv.Itoa(i)
		active := i == int(request.Page)
		links = append(links, PaginationLink{
			URL:    *url,
			Label:  label,
			Active: active,
		})
	}

	result = S_PaginateResult{
		Data: *resultSlicePtr.(*[]interface{}),
		Pagination: S_Pagination{
			PerPage:      request.PerPage,
			CurrentPage:  request.Page,
			LastPage:     int64(totalPages),
			FirstPageURL: firstPageUrl,
			NextPageURL:  nextPageUrl,
			PrevPageURL:  prevPageUrl,
			LastPageURL:  lastPageUrl,
			Path:         path,
			From:         int64(from),
			To:           to,
			Total:        count,
			Links:        links,
		},
	}

	return &result
}

func parseUrlRawQuery(context *gin.Context, path string, page, perPage int64) *string {
	rawQuery := context.Request.URL.RawQuery
	query, _ := url.ParseQuery(rawQuery)

	query.Set("page", strconv.Itoa(int(page)))
	query.Set("per_page", strconv.Itoa(int(perPage)))
	rawQuery = query.Encode()

	url := fmt.Sprintf("%s%s", path, rawQuery)

	return &url
}
