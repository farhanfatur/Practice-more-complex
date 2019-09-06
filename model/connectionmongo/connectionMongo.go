package connectionmongo

import (
	"TheLast/model/baseconnection"
	"TheLast/service"
	"log"

	"gopkg.in/mgo.v2"
)

type mongoStruct struct {
}

var takeConnectionMongo *mgo.Session

// Connection Mongo
func (m mongoStruct) Connection() {
	var session, err = mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err.Error())
	}
	takeConnectionMongo = session
}

func GetConnectionMongo() *mgo.Session {
	return takeConnectionMongo.Clone()
}

func NewMongo() service.ConnectInterface {
	return &mongoStruct{}
}

func init() {
	baseconnection.RegisterConnection(NewMongo())
}
