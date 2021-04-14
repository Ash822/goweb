package controller

import (
	"bytes"
	"encoding/json"
	. "github.com/ash822/goweb/entity"
	"github.com/ash822/goweb/repository"
	"github.com/ash822/goweb/service"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	msgRepo = repository.GetInstance()
	msgSvc  = service.GetInstance(msgRepo)
	msgCtrl = GetInstance(msgSvc)
)

const (
	id1 = "1"
	id2 = "2"
	validInput1 = "abba"
	validInput2 = "kayak"
	invalidInput = "baby"
)

func Setup() string {
	var msg = MessageResponse{
		Id: id1,
		Text:  validInput1,
		Palindrome: true,
	}
	msgRepo.Create(&msg)
	return msg.Id
}

func TearDown(msg *MessageResponse) {
	msgRepo.Delete(msg.Id)
}

func TestController_CreateMessage(t *testing.T) {
	var body = []byte(`{"id": "` + id2 + `", "text": "` + validInput2 + `"}`)
	req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(body))

	handler := http.HandlerFunc(msgCtrl.CreateMessage)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusCreated
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var msg MessageResponse
	json.NewDecoder(io.Reader(response.Body)).Decode(&msg)

	// Assert HTTP response
	g := NewGomegaWithT(t)

	// Id will be auto generated and not be equal to the body
	g.Expect(msg.Id).ShouldNot(Equal(id2))
	g.Expect(msg.Id).ShouldNot(Equal(""))

	// Text must be equal to the body
	g.Expect(msg.Text).Should(Equal(validInput2))

	// And it should be a valid palindrome
	g.Expect(msg.Palindrome).To(BeTrue())

	// Clean up collection
	TearDown(&msg)
}

func TestController_CreateMessage2(t *testing.T) {
	var body = []byte(`{"text1": "` + validInput2 + `"}`)
	req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(body))

	handler := http.HandlerFunc(msgCtrl.CreateMessage)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("the message text is invalid or not found"))
}

func TestController_UpdateMessage(t *testing.T) {
	id := Setup()

	var body = []byte(`{"text": "` + invalidInput + `"}`)

	req, _ := http.NewRequest("POST", "/message/" + id, bytes.NewBuffer(body))

	handler := http.HandlerFunc(msgCtrl.UpdateMessage)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusOK
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var msg MessageResponse
	json.NewDecoder(io.Reader(response.Body)).Decode(&msg)

	g := NewGomegaWithT(t)

	// Id should not be modified
	g.Expect(msg.Id).Should(Equal(id))

	// Text must be equal to the body
	g.Expect(msg.Text).Should(Equal(invalidInput))

	// And it should be an invalid palindrome
	g.Expect(msg.Palindrome).NotTo(BeTrue())

	TearDown(&msg)
}

func TestController_UpdateMessage2(t *testing.T) {
	var body = []byte(`{"id": "` + id2 + `", "text": "` + invalidInput + `"}`)

	req, _ := http.NewRequest("POST", "/message/", bytes.NewBuffer(body))

	handler := http.HandlerFunc(msgCtrl.UpdateMessage)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("the id provided is invalid"))
}

func TestController_DeleteMessage(t *testing.T) {
	id := Setup()

	req, _ := http.NewRequest("DELETE", "/message/" + id, nil)

	handler := http.HandlerFunc(msgCtrl.DeleteMessage)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusNoContent
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}
}

func TestController_DeleteMessage2(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/message/11", nil)

	handler := http.HandlerFunc(msgCtrl.DeleteMessage)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("a message does not exists for the given Id: 11"))
}


func TestController_GetAllMessages(t *testing.T) {
	Setup()

	req, _ := http.NewRequest("GET", "/messages", nil)

	handler := http.HandlerFunc(msgCtrl.GetAllMessages)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusOK
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}
}

func TestController_GetMessageById(t *testing.T) {
	Setup()

	req, _ := http.NewRequest("GET", "/message/11", nil)

	handler := http.HandlerFunc(msgCtrl.GetMessageById)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("a message does not exists for the given Id: 11"))
}

func TestController_GetMessageById2(t *testing.T) {
	id := Setup()

	req, _ := http.NewRequest("GET", "/message/"+id, nil)

	handler := http.HandlerFunc(msgCtrl.GetMessageById)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	expected := http.StatusOK
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	var msg MessageResponse
	json.NewDecoder(io.Reader(response.Body)).Decode(&msg)

	g := NewGomegaWithT(t)

	// Id should not be modified
	g.Expect(msg.Id).Should(Equal(id))

	// Text must be equal to the body
	g.Expect(msg.Text).Should(Equal(validInput1))

	// And it should be an invalid palindrome
	g.Expect(msg.Palindrome).To(BeTrue())

	TearDown(&msg)
}
