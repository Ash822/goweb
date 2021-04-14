package controller

import (
	"encoding/json"
	. "github.com/ash822/goweb/entity"
	"github.com/ash822/goweb/service"
	"net/http"
	"strings"
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

	id := strings.TrimPrefix(req.URL.Path, "/message/")

	msg, err := svc.FindById(id)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode(ServiceError{Error: err.Error()})
		return
	}

	resw.WriteHeader(http.StatusOK)
	json.NewEncoder(resw).Encode(msg)
}

func (*controller) GetAllMessages(resw http.ResponseWriter, _ *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	msgs, err := svc.FindAll()

	if err != nil {
		resw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resw).Encode(ServiceError{Error: "encountered error fetching the messages"})
	} else {
		resw.WriteHeader(http.StatusOK)
		json.NewEncoder(resw).Encode(msgs)
	}
}

func (*controller) CreateMessage(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	var msgReq MessageRequest
	var err error

	json.NewDecoder(req.Body).Decode(&msgReq)

	result, err := svc.Create(&msgReq)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode(ServiceError{Error: err.Error()})
		return
	}

	resw.WriteHeader(http.StatusCreated)
	json.NewEncoder(resw).Encode(result)
}

func (*controller) UpdateMessage(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	var msgReq MessageRequest

	id := strings.TrimPrefix(req.URL.Path, "/message/")

	json.NewDecoder(req.Body).Decode(&msgReq)

	result, err := svc.Update(id, &msgReq)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode(ServiceError{Error: err.Error()})
		return
	}

	resw.WriteHeader(http.StatusOK)
	json.NewEncoder(resw).Encode(result)
}

func (*controller) DeleteMessage(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-type", "application/json")

	id := strings.TrimPrefix(req.URL.Path, "/message/")

	err := svc.Delete(id)

	if err != nil {
		resw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resw).Encode(ServiceError{Error: err.Error()})
		return
	}

	resw.WriteHeader(http.StatusNoContent)
}
