package skillService

import (
	"bhakti-buana-api/src/configs"
	"bhakti-buana-api/src/constants"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	skillRequest "bhakti-buana-api/src/requests/skill"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type S_IndexServiceResult struct {
	Data       []models.Skills
	Pagination helpers.S_Pagination
}

// Skill Store Service
/*
 * @param context *gin.Context
 * @param request *skillRequest.S_StoreRequest
 * @returns *models.Skills
 */
func Store(context *gin.Context, request *skillRequest.S_StoreRequest) *models.Skills {
	var skill models.Skills

	var dir string
	if configs.AppConfig().GIN_MODE == "release" {
		dir = "../public/skills/"
	} else {
		dir = "./public/skills/"
	}
	helpers.CheckDir(dir)

	newFileName := fmt.Sprintf("%v-%s", time.Now().Unix(), request.Image.Filename)

	if err := context.SaveUploadedFile(request.Image, dir+newFileName); err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	parsedUrl, _ := url.Parse(configs.AppConfig().BASE_URL + "/public/skills/" + newFileName)

	filter := bson.M{"deleted_at": bson.M{"$eq": nil}}

	var order int64

	if request.Order != 0 {
		order = request.Order

		filterEqual := bson.M{"$and": []bson.M{filter, {"order": bson.M{"$eq": order}}}}
		countEqual, _ := database.Skills.CountDocuments(context, filterEqual)

		if countEqual != 0 {
			filter = bson.M{"$and": []bson.M{filter, {"order": bson.M{"$gte": order}}}}
			count, _ := database.Skills.CountDocuments(context, filter)

			if count > 0 {
				opts := options.Find().SetSort(map[string]int{"order": 1})
				cursor, err := database.Skills.Find(context, filter, opts)
				if err != nil {
					helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
					return nil
				}
				defer cursor.Close(context)

				var result models.Skills
				for cursor.Next(context) {
					cursor.Decode(&result)

					updateFilter := bson.M{"_id": result.ID}

					payload := bson.M{
						"$set": bson.M{
							"order":      result.Order + 1,
							"updated_at": time.Now(),
						},
					}

					_, err := database.Skills.UpdateOne(context, updateFilter, payload)
					if err != nil {
						helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
						return nil
					}
				}
			}
		}
	} else {
		opts := options.Find().SetSort(map[string]int{"order": -1})
		cursor, err := database.Skills.Find(context, filter, opts)
		if err != nil {
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}
		defer cursor.Close(context)

		var result models.Skills
		if cursor.Next(context) {
			if err := cursor.Decode(&result); err != nil {
				order = 1
			} else {
				order = result.Order + 1
			}
		} else {
			order = 1
		}
	}

	skill = models.Skills{
		Code:      request.Code,
		Order:     order,
		Name:      request.Name,
		SourceUrl: request.SourceUrl,
		ImageUrl:  parsedUrl.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := database.Skills.InsertOne(context, &skill)
	if err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	skill.ID = result.InsertedID.(primitive.ObjectID)

	return &skill
}

// Skill Index Service
/*
 * @param context *gin.Context
 * @param request *skillRequest.S_IndexRequest
 * @returns *S_IndexServiceResult
 */
func Index(context *gin.Context, request *skillRequest.S_IndexRequest) *S_IndexServiceResult {
	var result S_IndexServiceResult

	paginateResult := helpers.Paginate(context, request.S_PaginationRequest, database.Skills, reflect.TypeOf(models.Skills{}), bson.M{}, "order")

	skills := make([]models.Skills, 0)
	for _, item := range paginateResult.Data {
		if skill, ok := item.(models.Skills); ok {
			skills = append(skills, skill)
		}
	}

	result = S_IndexServiceResult{
		Data:       skills,
		Pagination: paginateResult.Pagination,
	}

	return &result
}

// Skill Update Service
/*
 * @param context *gin.Context
 * @param request *skillRequest.S_UpdateRequest
 * @returns *models.Skills
 */
func Update(context *gin.Context, request *skillRequest.S_UpdateRequest) *models.Skills {
	var skill models.Skills

	_id, _ := primitive.ObjectIDFromHex(request.ID)
	filter := bson.M{"_id": _id}

	checkFilter := bson.M{"deleted_at": bson.M{"$eq": nil}}

	order := request.Order
	filterEqual := bson.M{"$and": []bson.M{checkFilter, {"order": bson.M{"$eq": order}}}}
	countEqual, _ := database.Skills.CountDocuments(context, filterEqual)

	if countEqual != 0 {
		checkFilter = bson.M{"$and": []bson.M{checkFilter, {"order": bson.M{"$gte": order}}}}
		count, _ := database.Skills.CountDocuments(context, checkFilter)

		if count > 0 {
			opts := options.Find().SetSort(map[string]int{"order": 1})
			cursor, err := database.Skills.Find(context, checkFilter, opts)
			if err != nil {
				helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
				return nil
			}
			defer cursor.Close(context)

			var result models.Skills
			for cursor.Next(context) {
				cursor.Decode(&result)

				updateFilter := bson.M{"_id": result.ID}

				payload := bson.M{
					"$set": bson.M{
						"order":      result.Order + 1,
						"updated_at": time.Now(),
					},
				}

				_, err := database.Skills.UpdateOne(context, updateFilter, payload)
				if err != nil {
					helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
					return nil
				}
			}
		}
	}

	var payload primitive.M

	if request.Image != nil {
		if err := database.Skills.FindOne(context, filter).Decode(&skill); err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
				return nil
			default:
				helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
				return nil
			}
		}

		var dir string
		if configs.AppConfig().GIN_MODE == "release" {
			dir = "../public/skills/"
		} else {
			dir = "./public/skills/"
		}
		helpers.CheckDir(dir)

		newFileName := fmt.Sprintf("%v-%s", time.Now().Unix(), request.Image.Filename)

		if err := context.SaveUploadedFile(request.Image, dir+newFileName); err != nil {
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}

		oldParts := strings.Split(skill.ImageUrl, "/")
		oldFileName := oldParts[len(oldParts)-1]
		if skill.ImageUrl != "" {
			os.Remove(dir + oldFileName)
		}

		parsedUrl, _ := url.Parse(configs.AppConfig().BASE_URL + "/public/skills/" + newFileName)

		payload = bson.M{
			"$set": bson.M{
				"code":       request.Code,
				"order":      order,
				"name":       request.Name,
				"source_url": request.SourceUrl,
				"image_url":  parsedUrl.String(),
				"updated_at": time.Now(),
			},
		}
	} else {
		payload = bson.M{
			"$set": bson.M{
				"code":       request.Code,
				"order":      order,
				"name":       request.Name,
				"source_url": request.SourceUrl,
				"updated_at": time.Now(),
			},
		}
	}

	result, err := database.Skills.UpdateOne(context, filter, payload)
	if err != nil {
		helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
		return nil
	}

	if result.ModifiedCount == 0 {
		helpers.HttpResponse(constants.ID_NOT_FOUND, http.StatusNotFound, context, nil)
		return nil
	}

	if err := database.Skills.FindOne(context, filter).Decode(&skill); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, err.Error())
			return nil
		}
	}

	return &skill
}
