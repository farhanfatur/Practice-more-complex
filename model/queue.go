package model

import (
	"fmt"
	"log"
	"runtime"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	. "TheLast/model/backgroundservice"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/mgo.v2"
)

type QueueRedis struct {
	ConnDial *redis.Pool
}

func (c *QueueRedis) Push(key interface{}) bool {
	con := c.ConnDial.Get()
	redis.Strings(con.Do("RPUSH", "list", key))
	fmt.Println(runtime.NumGoroutine())
	defer con.Close()
	FuncName <- "Push"
	Val <- true
	return true
}

func (c *QueueRedis) Pop() interface{} {
	con := c.ConnDial.Get()
	var count, _ = con.Do("LLEN", "list")
	var index, err = redis.String(con.Do("LINDEX", "list", "0"))
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
	if _, err2 := con.Do("LPOP", "list"); err2 != nil {
		panic(err2)
	}
	defer con.Close()
	FuncName <- "Pop"
	Val <- index
	return index
}

func (c *QueueRedis) Keys() []interface{} {
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

func (c *QueueRedis) Len() int {
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

func (c *QueueRedis) Contains(key interface{}) bool {
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
			gConvert, _ := strconv.Atoi(g)
			if key == gConvert {
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

// ==================================================================================================================================================

type QueueModelMongo struct {
	ID     bson.ObjectId `bson:"_id" json:"_id"`
	Number string        `bson:"number" json:"number"`
}

func (q *QueueModelMongo) TableName() string {
	return "Queue"
}

type QueueMongo struct {
	connMongo *mgo.Session
}

func (q *QueueMongo) Push(key interface{}) bool {
	var model QueueModelMongo
	var conn = q.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	err := c.Insert(&QueueModelMongo{ID: bson.NewObjectId(), Number: key.(string)})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	FuncName <- "Push"
	Val <- true
	return true
}

func (q *QueueMongo) Keys() []interface{} {
	var model QueueModelMongo
	var getData = []QueueModelMongo{}
	var conn = q.connMongo.Clone()
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

func (q *QueueMongo) Len() int {
	var model QueueModelMongo
	var conn = q.connMongo.Clone()
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

func (q *QueueMongo) Pop() interface{} {
	var model QueueModelMongo
	var conn = q.connMongo.Clone()
	c := conn.DB("hello").C(model.TableName())
	err := c.Find(bson.M{}).One(&model)
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

func (q *QueueMongo) Contains(key interface{}) bool {
	var exist bool
	var model QueueModelMongo
	var getData = []QueueModelMongo{}
	var conn = q.connMongo.Clone()
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
	defer conn.Close()
	FuncName <- "Contains"
	Val <- exist
	return exist
}
