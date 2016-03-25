package user

import (
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(a *gin.Engine) {

	a.POST("/users/register/", Register)
	a.POST("/users/login/", Login)

}
