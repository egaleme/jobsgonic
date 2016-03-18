package common

import (
	"io/ioutil"
	"log"
)

const (
	privKeyPath = "app/keys/app.rsa"
	pubKeyPath  = "app/keys/app.rsa.pub"
)

var (
	VerifyKey, SignKey []byte
)

func initKeys() {
	var err error
	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys] : %s\n", err)
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys] : %s\n", err)
		panic(err)
	}

}
