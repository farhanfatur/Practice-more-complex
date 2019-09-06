package backgroundservice

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var DbName = make(chan string, 100)
var FuncName = make(chan string, 1000)
var Val = make(chan interface{}, 1000)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
	c.ServeMux.ServeHTTP(w, r)
}

func DBService(DbName <-chan string) {
	log.Println("Database", <-DbName, "is connected")
}

func ModelService(FuncName chan string, Val chan interface{}) {
	var mtx sync.Mutex
	mtx.Lock()
	if <-Val == true {
		close(Val)
		Val <- "true"
	}
	for {
		log.Println("Action", <-FuncName, "is success by parsing", <-Val)
	}

	mtx.Unlock()
}

func ConfigService(r *CustomMux) *http.Server {
	server := new(http.Server)
	server.Handler = r
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 5 * time.Second
	server.Addr = ":8000"
	log.Println("Server started at localhost:8000")
	return server
}
