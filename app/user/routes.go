package user

import (
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(a gin.IRouter) {

	a.POST("/users/register/", Register)
	a.POST("/users/login/", Login)

}
