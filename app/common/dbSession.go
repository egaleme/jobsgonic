package common

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func createDbSession() {
	uri := "mongodb://egaleme:200owina07@ds021299.mlab.com:21299/jobeet"
	session, err := mgo.Dial(uri)

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
