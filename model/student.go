package model

import (
	"log"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataStudent struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

type StudentModel []struct {
	ID    bson.ObjectId `bson:"_id" json:"_id"`
	Name  string        `bson:"name" json:"name"`
	Age   int           `bson:"age" json:"age"`
	Grade int           `bson:"grade" json:"grade"`
}

func (s *StudentModel) TableName() string {
	return "student"
}

type StudentCrud struct {
	session *mgo.Session
}

func (s *StudentCrud) ViewAll() StudentModel {
	var studentModel StudentModel
	var clone = s.session.Clone()
	c := clone.DB("hello").C(studentModel.TableName())
	err := c.Find(bson.M{}).All(&studentModel)
	if err != nil {
		panic(err.Error())
	}
	defer clone.Close()
	return studentModel
}

func (s *StudentCrud) Insert(data DataStudent) map[string]interface{} {
	var studentModel StudentModel
	var clone = s.session.Clone()
	c := clone.DB("hello").C(studentModel.TableName())
	var M map[string]interface{}
	err := c.Insert(&DataStudent{data.Name, data.Age, data.Grade})
	if err != nil {
		log.Fatal(err.Error())
		M = map[string]interface{}{
			"status": "error",
		}
	}
	M = map[string]interface{}{
		"status": "success",
	}
	defer clone.Close()
	return M
}

func (s *StudentCrud) Edit(id string) map[string]interface{} {
	var studentModel StudentModel
	var M map[string]interface{}
	for _, v := range studentModel {
		M = map[string]interface{}{
			"id":    v.ID,
			"name":  v.Name,
			"age":   v.Age,
			"grade": v.Grade,
		}
	}
	var clone = s.session.Clone()
	c := clone.DB("hello").C(studentModel.TableName())
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&M)
	if err != nil {
		panic(err.Error())
	}
	defer clone.Close()
	return M
}

func (s *StudentCrud) Update(data map[string]interface{}) map[string]interface{} {
	var studentModel StudentModel
	var result map[string]interface{}
	var clone = s.session.Clone()
	var c = clone.DB("hello").C(studentModel.TableName())
	ageConvert, _ := strconv.Atoi(data["age"].(string))
	gradeConvert, _ := strconv.Atoi(data["grade"].(string))
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(data["_id"].(string))}, bson.M{"$set": bson.M{"name": data["name"], "age": ageConvert, "grade": gradeConvert}})
	if err != nil {
		result = map[string]interface{}{
			"status": "error",
		}
		panic(err.Error())
	}
	result = map[string]interface{}{
		"status": "success",
	}
	defer clone.Close()
	return result
}
