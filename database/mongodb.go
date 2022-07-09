package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

var Mgo *mgo.Collection

func ConnectMongo() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	Mgo = session.DB("GinLearning").C("TestMongoDB")
}
