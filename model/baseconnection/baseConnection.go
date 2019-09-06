package baseconnection

import (
	"TheLast/service"
	"time"
)

var registerLocal []service.ConnectInterface

func RegisterConnection(cn service.ConnectInterface) bool {
	registerLocal = append(registerLocal, cn)
	time.Sleep(100 * time.Millisecond)
	return true
}

func DoConnect(opt string) {
	for _, val := range registerLocal {
		val.Connection()
	}
}
