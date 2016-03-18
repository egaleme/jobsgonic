package common

import (
	"gopkg.in/mgo.v2"
)

func NewSession() *mgo.Session {
	sess := GetSession().Copy()
	return sess

}
