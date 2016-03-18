package user

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId `bson:"_id" json:"id"`
	Firstname    string        `json:"firstname"`
	Lastname     string        `json:"lastname"`
	Password     string        `json:"password"`
	HashPassword []byte        `json:"hashpassword"`
	Email        string        `json:"email"`
}

type LoginUser struct {
	Email    string
	Password string
}

type AuthUser struct {
	Message   string `json:"message"`
	Firstname string `json:"firstname"`
	Token     string `json:"token"`
	Email     string `json:"email"`
}

type MessageResource struct {
	Message string `json:"message"`
}
