package controller

import (
	"encoding/json"
	"github.com/ash822/goweb/entity"
	"github.com/ash822/goweb/service"
	"github.com/gorilla/mux"
	"net/http"
)

type MessageController interface {
	GetMessageById(resw http.ResponseWriter, req *http.Request)
	GetAllMessages(resw http.ResponseWriter, req *http.Request)
	CreateMessage(resw http.ResponseWriter, req *http.Request)
	UpdateMessage(resw http.ResponseWriter, req *http.Request)
	DeleteMessage(resw http.ResponseWriter, req *http.Request)
}

var svc service.MessageService

type controller struct{}

func GetInstance(msgSvc service.MessageService) MessageController {
	svc = msgSvc
	return &controller{}
}

func (*controller) GetMessageById(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	msg, err := svc.FindById(id)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode([]byte(`{"error": "` + "err.Error()" + `"}`))
	}

	resw.WriteHeader(http.StatusOK)
	json.NewEncoder(resw).Encode(msg)
}

func (*controller) GetAllMessages(resw http.ResponseWriter, _ *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	msgs, err := svc.FindAll()

	if err != nil {
		resw.WriteHeader(http.StatusInternalServerError)
		resw.Write([]byte(`{"error": "Encountered error fetching the messages"`))
	} else {
		resw.WriteHeader(http.StatusOK)
		json.NewEncoder(resw).Encode(msgs)
	}
}

func (*controller) CreateMessage(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	var msg entity.Message
	var err error

	json.NewDecoder(req.Body).Decode(&msg)

	result, err := svc.Create(&msg)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	resw.WriteHeader(http.StatusCreated)
	json.NewEncoder(resw).Encode(result)
}

func (*controller) UpdateMessage(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	var msg entity.Message

	params := mux.Vars(req)
	id := params["id"]

	json.NewDecoder(req.Body).Decode(&msg)

	result, err := svc.Update(id, &msg)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	resw.WriteHeader(http.StatusOK)
	json.NewEncoder(resw).Encode(result)
}

func (*controller) DeleteMessage(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	err := svc.Delete(id)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	resw.WriteHeader(http.StatusNoContent)
}
