package model

import (
	"fmt"
	"log"

	. "TheLast/model/backgroundservice"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type StackRedis struct {
	ConnDial *redis.Pool
}

func (c *StackRedis) Push(key interface{}) bool {
	con := c.ConnDial.Get()
	redis.Strings(con.Do("RPUSH", "list", key))
	defer con.Close()
	FuncName <- "Push"
	Val <- true
	return true
}

func (c *StackRedis) Pop() interface{} {
	con := c.ConnDial.Get()
	var count, _ = con.Do("LLEN", "list")
	var index, err = redis.String(con.Do("LINDEX", "list", "-1"))
	if err != nil {
		fmt.Println("Data is empty")
	}
	if count.(int64) <= 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Data is empty")
			}
		}()
	}
	if _, err2 := con.Do("RPOP", "list"); err2 != nil {
		panic(err2)
	}
	FuncName <- "Pop"
	Val <- index
	defer con.Close()
	return index
}

func (c *StackRedis) Keys() []interface{} {
	con := c.ConnDial.Get()
	var result, err = redis.Strings(con.Do("LRANGE", "list", "0", "-1"))
	if err != nil {
		panic(err)
	}
	if len(result) <= 0 {
		fmt.Println("Data is empty")
	}
	resultParse := make([]interface{}, 0)
	for _, each := range result {
		resultParse = append(resultParse, each)
	}
	defer con.Close()
	FuncName <- "Keys"
	Val <- resultParse
	return resultParse

}

func (c *StackRedis) Len() int {
	con := c.ConnDial.Get()
	var count, _ = redis.Int(con.Do("LLEN", "list"))
	if count <= 0 {
		fmt.Println("Data is empty")
	}
	defer con.Close()
	FuncName <- "Len"
	Val <- count
	return count
}

func (c *StackRedis) Contains(key interface{}) bool {
	con := c.ConnDial.Get()
	var exist bool
	var result, err = redis.Strings(con.Do("LRANGE", "list", "0", "-1"))
	if err != nil {
		panic(err)
	}
	if len(result) == 1 {
		exist = true
	} else {
		for _, g := range result[:len(result)-1] {
			if key == g {
				if _, errD := con.Do("RPOP", "list"); errD != nil {
					panic(err)
				}

				exist = false
			} else {
				exist = true
			}
		}
	}
	defer con.Close()
	FuncName <- "Contains"
	Val <- exist
	return exist
}

// ====================================================================================================================================================

type StackModelMongo struct {
	ID     bson.ObjectId `bson:"_id" json:"_id"`
	Number string        `bson:"number" json:"number"`
}

func (s *StackModelMongo) TableName() string {
	return "Stack"
}

type StackMongo struct {
	connMongo *mgo.Session
}

func (s *StackMongo) Push(key interface{}) bool {
	var model QueueModelMongo
	var conn = s.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	err := c.Insert(&QueueModelMongo{ID: bson.NewObjectId(), Number: key.(string)})
	if err != nil {
		log.Fatal(err.Error())
	}
	FuncName <- "Push"
	Val <- true
	defer conn.Close()
	return true
}

func (s *StackMongo) Keys() []interface{} {
	var model QueueModelMongo
	var getData = []QueueModelMongo{}
	var conn = s.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	err := c.Find(bson.M{}).All(&getData)
	if err != nil {
		panic(err.Error())
	}
	var viewAll = make([]interface{}, 0)
	for _, val := range getData {
		viewAll = append(viewAll, val.Number)
	}
	defer conn.Close()
	FuncName <- "Keys"
	Val <- viewAll
	return viewAll
}

func (s *StackMongo) Len() int {
	var model QueueModelMongo
	var conn = s.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	num, err := c.Count()
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	FuncName <- "Len"
	Val <- num
	return num
}

func (s *StackMongo) Pop() interface{} {
	var model QueueModelMongo
	var conn = s.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	err := c.Find(bson.M{}).Limit(1).Sort("-$natural").One(&model)
	if err != nil {
		panic(err.Error())
	}
	err = c.Remove(bson.M{"_id": model.ID})
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	FuncName <- "Pop"
	Val <- model.Number
	return model.Number
}

func (s *StackMongo) Contains(key interface{}) bool {
	var exist bool
	var model QueueModelMongo
	var getData = []QueueModelMongo{}
	var conn = s.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	err := c.Find(bson.M{"number": key.(string)}).All(&getData)
	if err != nil {
		panic("Find error")
	}
	var dataSlice = make([]interface{}, 0)
	for _, val := range getData {
		dataSlice = append(dataSlice, val.Number)
	}
	if len(dataSlice) == 1 {
		fmt.Println("First")
		exist = true
	} else {
		for _, g := range dataSlice[:len(dataSlice)-1] {
			if key == g {
				err := conn.DB("hello").C(model.TableName()).Remove(bson.M{"number": key})
				if err != nil {
					panic(err.Error())
				}
				exist = false
			} else {
				exist = true
			}
		}
	}
	FuncName <- "Contains"
	Val <- exist
	defer conn.Close()
	return exist
}
