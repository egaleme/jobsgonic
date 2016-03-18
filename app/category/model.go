package category

import (
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	Id   bson.ObjectId ` bson:"_id" json:"id"`
	Name string        `json:"name"`
}

type CategoryResource struct {
	Message string `json:"message"`
}
