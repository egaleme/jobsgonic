package category

import (
	"github.com/egaleme/jobsgonic/app/common"
	"github.com/gin-gonic/gin"
)

func SetCategoryRoutes(a gin.IRouter) {
	categoryRoutes := a.Group("/category")
	categoryRoutes.Use(common.Authenticate)
	categoryRoutes.Use(common.Authorize)
	categoryRoutes.POST("/", CreateCategory)
	categoryRoutes.PUT("/:id", UpdateCategory)
	categoryRoutes.DELETE("/:id", DeleteCategory)

}
