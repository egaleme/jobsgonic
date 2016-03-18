package category

import (
	"errors"

	"github.com/egaleme/jobsgonic/app/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func CreateCategory(c *gin.Context) {

	var category Category
	c.BindJSON(&category)
	err := validateCategory(&category)
	if err != nil {
		common.DisplayAppError(c.Writer, err, "An unexpected error occurred", 500)
		return
	}
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("categories")
	repo := &CategoryRepository{collection}
	repo.Create(&category)
	c.JSON(201, &category)

}

func GetCategories(c *gin.Context) {

	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("categories")
	repo := &CategoryRepository{collection}
	categories, err := repo.GetAll()
	if err != nil {
		common.DisplayAppError(c.Writer, err, "An unexpected error occurred", 500)
		return
	}
	c.JSON(200, &categories)
}

func GetCategoryById(c *gin.Context) {

	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("categories")
	id := c.Param("id")
	repo := &CategoryRepository{collection}
	category, err := repo.GetById(id)
	if err != nil {
		common.DisplayAppError(c.Writer, err, "An unexpected error occurred", 500)
		return
	}
	c.JSON(200, &category)
}

func UpdateCategory(c *gin.Context) {

	id := c.Param("id")
	ds := common.NewSession()
	defer ds.Close()
	var category Category
	c.BindJSON(&category)
	category.Id = bson.ObjectIdHex(id)
	collection := ds.DB(common.AppConfig.Database).C("categories")
	repo := &CategoryRepository{collection}
	if err := repo.Update(&category); err != nil {
		common.DisplayAppError(
			c.Writer,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	c.JSON(201, &CategoryResource{Message: "successfully updated"})
}

func DeleteCategory(c *gin.Context) {

	id := c.Param("id")
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("categories")
	repo := &CategoryRepository{collection}
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			c.Writer,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	c.JSON(201, &CategoryResource{Message: "successfully deleted"})

}

func validateCategory(category *Category) error {
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("categories")
	repo := &CategoryRepository{collection}
	categoryStore, _ := repo.GetAll()
	for _, c := range categoryStore {
		if c.Name == category.Name {
			return errors.New("category already exits")
		}

	}
	return nil
}
