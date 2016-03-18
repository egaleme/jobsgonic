package job

import (
	"github.com/egaleme/jobsgonic/app/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func CreateJob(c *gin.Context) {

	userEmail, _ := c.Get("userEmail")
	//posterEmail := userEmail.(string)
	var job Job
	c.BindJSON(&job)
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("jobs")
	repo := &JobRepository{collection}
	//job.UserId = bson.ObjectIdHex(userid)
	job.PosterEmail = userEmail
	repo.Create(&job)
	c.JSON(201, job)

}

func GetJobs(c *gin.Context) {

	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("jobs")
	repo := &JobRepository{collection}
	jobs := repo.GetAll()
	c.JSON(200, jobs)
}

func GetJobById(c *gin.Context) {

	id := c.Param("id")
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("jobs")
	repo := &JobRepository{collection}
	job, _ := repo.GetById(id)
	c.JSON(200, job)

}

func UpdateJob(c *gin.Context) {

	userEmail, _ := c.Get("userEmail")
	posterEmail := userEmail.(string)
	id := c.Param("id")
	ds := common.NewSession()
	defer ds.Close()
	var job Job
	c.BindJSON(&job)
	job.Id = bson.ObjectIdHex(id)
	collection := ds.DB(common.AppConfig.Database).C("jobs")
	repo := &JobRepository{collection}
	if err := repo.Update(&job, posterEmail); err != nil {
		common.DisplayAppError(
			c.Writer,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	c.JSON(200, &JobResource{Message: "successfully updated"})

}

func DeleteJob(c *gin.Context) {

	userEmail, _ := c.Get("userEmail")
	posterEmail := userEmail.(string)
	id := c.Param("id")
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("jobs")
	repo := &JobRepository{collection}
	err := repo.Delete(id, posterEmail)
	if err != nil {
		common.DisplayAppError(
			c.Writer,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	c.JSON(200, &JobResource{Message: "successfully deleted"})

}

func GetJobsByUser(c *gin.Context) {

	userEmail, _ := c.Get("userEmail")
	posterEmail := userEmail.(string)
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("jobs")
	repo := &JobRepository{collection}
	jobs := repo.GetByUser(posterEmail)
	c.JSON(200, &jobs)

}
