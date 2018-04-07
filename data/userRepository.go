package data

import (
	"log"

	"github.com/ezechidc/taskmanager/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

//CreateUser ctreates and stores user object to the database
func (r *UserRepository) CreateUser(user *models.User) error {
	objID := bson.NewObjectId()
	user.ID = objID
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	//clear the incoming text password
	user.Password = ""
	log.Print(&user)
	err = r.C.Insert(&user)
	return err
}

//Login validates user password with that stored on the database.
func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	//Validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}
