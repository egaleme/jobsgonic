package user

import (
	"errors"

	"github.com/egaleme/jobsgonic/app/common"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user User
	c.BindJSON(&user)
	err := validateUser(&user)
	if err != nil {
		common.DisplayAppError(c.Writer, err, "An unexpected error occurred", 500)
		return
	}
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("users")
	repo := &UserRepository{collection}
	repo.CreateUser(&user)
	c.JSON(201, MessageResource{Message: "success"})
}

func Login(c *gin.Context) {

	var role string
	loginUser := LoginUser{}
	var token string
	c.BindJSON(&loginUser)
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("users")
	repo := &UserRepository{collection}
	if user, err := repo.Loginuser(loginUser); err != nil {
		common.DisplayAppError(c.Writer, err, "Invalid credentails", 401)
		return
	} else {
		//if login successful, generate token
		//userid := user.Id.String()
		if user.Email == common.AppConfig.Email {
			role = common.AppConfig.Role
		} else {
			role = "member"
		}

		token, err = common.GenerateJWT(user.Email, role)
		if err != nil {
			common.DisplayAppError(c.Writer, err, "Error generating token", 500)
			return
		}

		auth := AuthUser{Message: "success", Firstname: user.Firstname, Email: user.Email, Token: token}

		c.JSON(201, auth)

	}
}

func validateUser(user *User) error {
	ds := common.NewSession()
	defer ds.Close()
	collection := ds.DB(common.AppConfig.Database).C("users")
	repo := &UserRepository{collection}
	userStore, _ := repo.GetAll()
	for _, c := range userStore {
		if c.Email == user.Email {
			return errors.New("user already exits")
		}

	}
	return nil
}
