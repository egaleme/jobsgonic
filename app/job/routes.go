package job

import (
	"github.com/egaleme/jobsgonic/app/common"
	"github.com/gin-gonic/gin"
)

func SetJobRoutes(a *gin.Engine) {
	jobRoutes := a.Group("/jobs", common.Authenticate)
	jobRoutes.POST("/", CreateJob)
	jobRoutes.PUT("/:id", UpdateJob)
	jobRoutes.DELETE("/:id", DeleteJob)
	jobRoutes.GET("/user", GetJobsByUser)

}
