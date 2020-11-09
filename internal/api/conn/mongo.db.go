package conn

import (
	"github.com/errntry/errntry/internal/errors"
	"github.com/profclems/go-dotenv"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	host := dotenv.GetString("MONGO_HOST")
	dbName := dotenv.GetString("MONGO_DB_NAME")

	session, err := mgo.Dial(host)
	if err != nil {
		errors.Fatal(3, "session err:", err)
	}

	cred := &mgo.Credential{
		Username: dotenv.GetString("MONGO_DB_USER"),
		Password: dotenv.GetString("MONGO_DB_PASS"),
		Source:   dotenv.GetString("MONGO_AUTH_DB"),
	}
	err = session.Login(cred)
	db = session.DB(dbName)

	if err != nil {
		errors.Fatal(3, "session err:", err)
	}
}

// GetMongoDB function to return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
