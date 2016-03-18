package user

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *User) error {
	user.Id = bson.NewObjectId()
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	user.Password = ""
	err = r.C.Insert(&user)
	return err

}
func (r *UserRepository) Loginuser(loginUser LoginUser) (u User, err error) {
	err = r.C.Find(bson.M{"email": loginUser.Email}).One(&u)
	if err != nil {
		return
	}
	//validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(loginUser.Password))
	if err != nil {
		u = User{}
	}
	return

}

func (r *UserRepository) GetAll() ([]User, error) {
	var result []User
	err := r.C.Find(nil).Iter().All(&result)
	return result, err

}
