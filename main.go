package main

import (
	"TheLast/controller"
	"TheLast/model"
	. "TheLast/model/backgroundservice"
	"net/http"
)

func main() {
	var redisConn = model.Register("redis", "queue")
	// var wg sync.WaitGroup
	Controller := controller.BaseController{}
	Controller.RedisDo = redisConn
	router := new(CustomMux)
	router.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("asset"))))
	router.HandleFunc("/", Controller.View)
	// Stack & Queue(Redis + Mongo)
	// ------------------------------------------------
	router.HandleFunc("/api/keys", Controller.QueueController.Keys)
	router.HandleFunc("/api/lens", Controller.QueueController.Lens)
	router.HandleFunc("/api/pops", Controller.QueueController.Pop)
	router.HandleFunc("/api/pushs", Controller.QueueController.Push)
	// ------------------------------------------------

	// Crud(Mongo)
	// ------------------------------------------------
	// router.HandleFunc("/api/viewalls", Controller.CrudController.ViewAll)
	// router.HandleFunc("/api/inserts", Controller.CrudController.Insert)
	// router.HandleFunc("/api/edits/", Controller.CrudController.Edit)
	// router.HandleFunc("/api/updates", Controller.CrudController.Update)
	// ------------------------------------------------

	go DBService(DbName)
	go ModelService(FuncName, Val)
	server := ConfigService(router)
	server.ListenAndServe()

}
