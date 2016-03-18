package job

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	Id          bson.ObjectId ` bson:"_id" json:"id"`
	Description string        `json:"description"`
	CategoryId  bson.ObjectId `json:"categoryid"`
	Company     string        `json:"company"`
	Position    string        `json:"position"`
	Location    string        `json:"location"`
	ExpiresAt   time.Time     `json:"expiresat"`
	CreatedAt   time.Time     `json:"createdat"`
	UpdatedAt   time.Time     `json:"updatedat"`
	//UserId      bson.ObjectId `json:"userid"`
	PosterEmail interface{} `json:"poster"`
}

type JobResource struct {
	Message string `json:"message"`
}
