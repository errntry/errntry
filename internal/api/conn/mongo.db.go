package conn

import (
	"github.com/errntry/errntry/internal/config"
	"github.com/errntry/errntry/internal/errors"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	cfg := config.Load()

	host := cfg.Get("MONGO_HOST")
	dbName := cfg.Get("MONGO_DB_NAME")

	session, err := mgo.Dial(host)
	if err != nil {
		errors.Fatal(3, "session err:", err)
	}

	cred := &mgo.Credential{
		Username: cfg.Get("MONGO_DB_USER"),
		Password: cfg.Get("MONGO_DB_PASS"),
		Source:   cfg.Get("MONGO_AUTH_DB"),
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
