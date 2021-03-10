package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
	"github.com/sato-takumi20-fixer/gin-training-api/database"
	"github.com/sato-takumi20-fixer/gin-training-api/database/model"
)

var dbContext = database.CreateInMemoryDbContext()

type PostRequest struct {
	UserId int `json:"userId"`
	Formula string `json:"formula"`
}

type PutRequest PostRequest

func DeclareCalcApi(router *gin.Engine) {
	router.GET("/calc/:userId", func(context *gin.Context) {
		Get(context)
	})
	router.POST("/calc", func(context *gin.Context) {
		Post(context)
	})
	router.PUT("/calc/:userId/:id", func(context *gin.Context) {
		Put(context)
	})
	router.DELETE("/calc/:userId/:id", func(context *gin.Context) {
		Delete(context)
	})
}

func Get(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("userId"))
	if fatalIfError(context, err) {
		return
	}
	var records []model.Formula
	if fatalIfError(
		context, 
		dbContext.Order("ID asc").Where(&model.Formula{UserId: userId}).Find(&records).Error,
	) {
		return
	}
	context.JSON(http.StatusOK, gin.H{"records" :  records})
}

func Post(context *gin.Context) {
	var postCalcRequest PostRequest
	if fatalIfError(context, context.ShouldBindJSON(&postCalcRequest)) {
		return
	}
	expression, err := govaluate.NewEvaluableExpression(postCalcRequest.Formula);
	if fatalIfError(context, err) {
		return
	}
	result, err := expression.Evaluate(nil);
	if fatalIfError(context, err) {
		return
	}
	fatalIfError(
		context,
		dbContext.Create(&model.Formula{
			UserId: postCalcRequest.UserId,
			Formula: postCalcRequest.Formula,
			Result: fmt.Sprintf("%v", result),
		}).Error,
	)
	
	context.JSON(
		http.StatusOK, 
		gin.H{
			"formula": postCalcRequest.Formula,
			"result" :  result,
		},
	)
}

func Put(context *gin.Context) {
	var putCalcRequest PutRequest
	if fatalIfError(context, context.ShouldBindJSON(&putCalcRequest)) {
		return
	}
	userId, err := strconv.Atoi(context.Param("userId"))
	if fatalIfError(context, err) {
		return
	}
	id, err := strconv.Atoi(context.Param("id"))
	if fatalIfError(context, err) {
		return
	}
	expression, err := govaluate.NewEvaluableExpression(putCalcRequest.Formula);
	if fatalIfError(context, err) {
		return
	}
	result, err := expression.Evaluate(nil);
	if fatalIfError(context, err) {
		return
	}
	var record model.Formula
	if fatalIfError(
		context, 
		dbContext.Where(&model.Formula{ID: id, UserId: userId}).First(&record).Error,
	) {
		return
	}
	record.Formula = putCalcRequest.Formula
	record.Result = fmt.Sprintf("%v", result)
	fatalIfError(
		context,
		dbContext.Save(&record).Error,
	)
	context.JSON(http.StatusOK, gin.H{"formula": putCalcRequest.Formula,"result" :  result})
}

func Delete(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("userId"))
	if fatalIfError(context, err) {
		return
	}
	id, err := strconv.Atoi(context.Param("id"))
	if fatalIfError(context, err) {
		return
	}
	var record model.Formula
	if fatalIfError(
		context, 
		dbContext.Where(&model.Formula{ID: id, UserId: userId}).First(&record).Error,
	) {
		return
	}
	fatalIfError(
		context,
		dbContext.Delete(&record).Error,
	)
}

func fatalIfError(context *gin.Context, err error) bool {
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
    }
	return false
}