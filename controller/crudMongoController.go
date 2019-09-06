package controller

// import (
// 	"TheLast/model"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"path"
// )

// type CrudController struct {
// }

// func (c *CrudController) ViewAll(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		var newConn = model.NewCrud()
// 		data := newConn.ViewAll()
// 		w.Header().Set("content-type", "application/json")
// 		jsonInByte, err := json.Marshal(data)
// 		if err != nil {
// 			panic(err)
// 		}
// 		w.Write(jsonInByte)
// 	} else {
// 		w.Write([]byte("Method is not valid!"))
// 	}
// }

// func (c *CrudController) Insert(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		// studentModel := model.StudentModel{}
// 		var data model.DataStudent
// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&data); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		var newConn = model.NewCrud()
// 		M := newConn.Insert(data)
// 		w.Header().Set("content-type", "application/json")
// 		jsonInByte, err := json.Marshal(M)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		w.Write(jsonInByte)
// 	} else {
// 		w.Write([]byte("Method is not valid!"))
// 	}
// }

// func (c *CrudController) Edit(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		var pathString = r.URL.Path
// 		var pathParse, _ = url.Parse(pathString)
// 		ID := path.Base(pathParse.Path)
// 		var newConn = model.NewCrud()
// 		var data = newConn.Edit(ID)
// 		w.Header().Set("content-type", "application/json")
// 		jsonInByte, err := json.Marshal(data)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		w.Write(jsonInByte)
// 	} else {
// 		w.Write([]byte("Method is not valid!"))
// 	}
// }

// func (c *CrudController) Update(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "PUT" {
// 		var data model.StudentModel
// 		var M map[string]interface{}
// 		for _, v := range data {
// 			M = map[string]interface{}{
// 				"_id":   v.ID,
// 				"name":  v.Name,
// 				"age":   v.Age,
// 				"grade": v.Grade,
// 			}
// 		}
// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&M); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		var newConn = model.NewCrud()
// 		res := newConn.Update(M)
// 		w.Header().Set("content-type", "applicaion/json")
// 		jsonInByte, err := json.Marshal(res)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		w.Write(jsonInByte)
// 	} else {
// 		w.Write([]byte("Method is not valid!"))
// 	}
// }
