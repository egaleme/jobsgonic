package common

import (
	"gopkg.in/mgo.v2"
	"log"
)

var session *mgo.Session

func createDbSession() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

}

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatalf("[GetSession]:%s\n", err)
		}
	}
	return session

}
