package common

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil

	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					c.Writer,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return
			default:
				DisplayAppError(c.Writer,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}
		default:
			DisplayAppError(c.Writer,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}
	}
	if token.Valid {

		c.Set("userEmail", token.Claims["email"])
		c.Set("userRole", token.Claims["userRole"])

		c.Next()
	} else {
		DisplayAppError(c.Writer,
			err,
			"Invalid Access Token",
			401,
		)
	}
}

func Authorize(c *gin.Context) {

	role, _ := c.Get("userRole")
	posterRole := role.(string)

	email, _ := c.Get("userEmail")
	posterEmail := email.(string)

	if posterRole == AppConfig.Role && posterEmail == AppConfig.Email {
		c.Next()

	} else {
		DisplayAppError(c.Writer,
			errors.New("not allowed"),
			"Not Authorized",
			401,
		)
	}
}

func GenerateJWT(email, role string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["iss"] = "admin"
	t.Claims["email"] = email
	t.Claims["userRole"] = role

	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tokenString, err := t.SignedString(SignKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
