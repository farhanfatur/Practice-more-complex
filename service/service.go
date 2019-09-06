package service

type ConnectInterface interface {
	Connection()
}

type ServiceAction interface {
	Push(key interface{}) bool
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}
