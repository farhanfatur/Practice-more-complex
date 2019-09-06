package model

import (
	. "TheLast/model/backgroundservice"
	"TheLast/model/baseconnection"
	"TheLast/model/connectionmongo"
	"TheLast/model/connectionredis"
	"TheLast/service"
)

/* Register is calling connection and register it in the same time
#db = get choose the database use ex: redis, mongo
opt = get the method waht we want ex: stackredis, queueredis, stackmongo, queuemongo#*/
func Register(db string, opt string) service.ServiceAction {
	var svc service.ServiceAction
	baseconnection.DoConnect(db)
	mgo := connectionmongo.GetConnectionMongo()
	redis := connectionredis.GetConnectionRedis()
	if opt == "stack" && db == "mongo" {
		svc = &StackMongo{mgo}
	} else if opt == "queue" && db == "mongo" {
		svc = &QueueMongo{mgo}
	}

	if opt == "stack" && db == "redis" {
		svc = &StackRedis{redis}
	} else if opt == "queue" && db == "redis" {
		svc = &QueueRedis{redis}
	}
	DbName <- db
	return svc
}
