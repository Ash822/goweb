package main

import (
	"github.com/ash822/goweb/controller"
	"github.com/ash822/goweb/repository"
	"github.com/ash822/goweb/router"
	"github.com/ash822/goweb/service"
	"net/http"
	_ "github.com/ash822/goweb/docs"
)

var (
	msgRepo       = repository.GetInstance()
	msgSvc        = service.GetInstance(msgRepo)
	msgController = controller.GetInstance(msgSvc)
	httpRouter    = router.HttpRouter()
)

func main() {
	httpRouter.Get("/", func(res http.ResponseWriter, req *http.Request) {
		msg := "Request served the server"
		payload := []byte(`{"message": "` + msg + `"}`)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(payload)
	})

	httpRouter.Get("/message/{id}", msgController.GetMessageById)
	httpRouter.Get("/messages", msgController.GetAllMessages)
	httpRouter.Post("/message", msgController.CreateMessage)
	httpRouter.Post("/message/{id}", msgController.UpdateMessage)
	httpRouter.Delete("/message/{id}", msgController.DeleteMessage)

	httpRouter.CreateServer(PORT)
}
