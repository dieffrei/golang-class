package service

import (
	"gopkg.in/mgo.v2"
	"os"
)

func GetSession() *mgo.Session {
	dbUrl := os.Getenv("DB_URL")
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session
}

func GetCollection(collectionName string) *mgo.Collection {
	dbName := os.Getenv("DB_NAME")
	return GetSession().DB(dbName).C(collectionName)
}