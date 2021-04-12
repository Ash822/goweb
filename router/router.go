package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Router interface {
	CreateServer(port string)
	Get(path string, fn func(resw http.ResponseWriter, req *http.Request))
	Post(path string, fn func(resw http.ResponseWriter, req *http.Request))
	Delete(path string, fn func(resw http.ResponseWriter, req *http.Request))
}

type httpRouter struct {}

var muxRouter = mux.NewRouter()

func HttpRouter() Router {
	return &httpRouter{}
}

func (*httpRouter) Get(path string, fn func(resw http.ResponseWriter, req *http.Request)) {
	muxRouter.HandleFunc(path, fn).Methods("GET")
}

func (*httpRouter) Post(path string, fn func(resw http.ResponseWriter, req *http.Request)) {
	muxRouter.HandleFunc(path, fn).Methods("POST")
}

func (*httpRouter) Delete(path string, fn func(resw http.ResponseWriter, req *http.Request)) {
	muxRouter.HandleFunc(path, fn).Methods("DELETE")
}

func (*httpRouter) CreateServer(port string) {
	log.Printf("Server is listening at port:%s", port)
	http.ListenAndServe(":"+port, muxRouter)
}
