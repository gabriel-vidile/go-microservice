package controllers

import (
	"net/http"

	"github.com/gabriel-vidile/go-microservice/internal/repositories"
	"github.com/gabriel-vidile/go-microservice/internal/services"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {

	services := services.NewGetCategoryService(repository)
	categories, err := services.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"success": true, "categories": categories})
}
