package router

import (
	"github.com/egaleme/jobsgonic/app/category"
	"github.com/egaleme/jobsgonic/app/job"
	"github.com/egaleme/jobsgonic/app/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(a *gin.Engine) {

	a.GET("/api/jobs/", job.GetJobs)
	a.GET("/api/jobs/:id", job.GetJobById)
	a.GET("/api/category/", category.GetCategories)
	a.GET("/api/category/:id", category.GetCategoryById)
	job.SetJobRoutes(a)
	user.SetUserRoutes(a)
	category.SetCategoryRoutes(a)

}
