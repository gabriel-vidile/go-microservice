package controllers

import (
	"net/http"

	"github.com/gabriel-vidile/go-microservice/internal/repositories"
	"github.com/gabriel-vidile/go-microservice/internal/services"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}
	services := services.NewCreateCategoryService(repository)
	err := services.Execute(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
