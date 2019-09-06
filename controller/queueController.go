package controller

import (
	"TheLast/model"
	"TheLast/service"
	"encoding/json"
	"log"
	"net/http"
)

type QueueController struct {
	RedisDo service.ServiceAction
}

func (l *QueueController) Push(w http.ResponseWriter, r *http.Request) {
	var pushModel = model.QueueModelMongo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pushModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	l.RedisDo.Push(pushModel.Number)
	l.RedisDo.Contains(pushModel.Number)
	w.Header().Set("content-type", "application/json")
	var response = map[string]interface{}{
		"alert": "success",
	}
	jsonInByte, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(jsonInByte)
}

func (l *QueueController) Keys(w http.ResponseWriter, r *http.Request) {

	var jsonInByte, err = json.Marshal(l.RedisDo.Keys())
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonInByte)
}

func (l *QueueController) Lens(w http.ResponseWriter, r *http.Request) {
	var jsonInByte, err = json.Marshal(l.RedisDo.Len())
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonInByte)
}

func (l *QueueController) Pop(w http.ResponseWriter, r *http.Request) {
	var jsonInByte, err = json.Marshal(l.RedisDo.Pop())
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonInByte)
}
