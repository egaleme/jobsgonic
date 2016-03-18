package common

func Startup() {
	InitConfig()
	createDbSession()
	initKeys()

}
