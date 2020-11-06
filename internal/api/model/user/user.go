package user

import (
	"time"

	"github.com/errntry/errntry/internal/api/conn"
	"gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name,omitempty"`
	Username   string        `json:"username" bson:"username,omitempty"`
	Email      string        `json:"email" bson:"email,omitempty"`
	Avatar     string        `json:"avatar" bson:"avatar,omitempty"`
	Password   string        `json:"password" bson:"password,omitempty"`
	IsVerified bool          `json:"is_verified" bson:"is_verified,omitempty"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
}

// Users list
type Users []User

// UserInfo model function
func Info(id bson.ObjectId, userCollection string) (User, error) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}
