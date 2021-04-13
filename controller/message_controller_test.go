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

func Setup() {
	var msg = Message{
		Id:    id1,
		Text:  validInput1,
		Palindrome: true,
	}
	msgRepo.Create(&msg)
}

func TearDown(msg *Message) {
	msgRepo.Delete(msg.Id)
}

func TestController_CreateMessage(t *testing.T) {
	//Create a new HTTP POST request
	var body = []byte(`{"id": "` + id2 + `", "text": "` + validInput2 + `"}`)
	req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(body))

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.CreateMessage)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusCreated
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var msg Message
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
	//Create a new HTTP POST request
	var body = []byte(`{"id": "` + id2 + `", "text1": "` + validInput2 + `"}`)
	req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(body))

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.CreateMessage)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	// Assert HTTP response
	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("the message text is invalid or not found"))
}

func TestController_UpdateMessage(t *testing.T) {
	// Provision the msg in the collection
	Setup()

	//Create a new HTTP POST request
	var body = []byte(`{"id": "` + id1 + `", "text": "` + invalidInput + `"}`)

	req, _ := http.NewRequest("POST", "/message/1", bytes.NewBuffer(body))

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.UpdateMessage)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusOK
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var msg Message
	json.NewDecoder(io.Reader(response.Body)).Decode(&msg)

	// Assert HTTP response
	g := NewGomegaWithT(t)

	// Id should not be modified
	g.Expect(msg.Id).Should(Equal(id1))

	// Text must be equal to the body
	g.Expect(msg.Text).Should(Equal(invalidInput))

	// And it should be an invalid palindrome
	g.Expect(msg.Palindrome).NotTo(BeTrue())

	// Clean up collection
	TearDown(&msg)
}

func TestController_UpdateMessage2(t *testing.T) {
	//Create a new HTTP POST request
	var body = []byte(`{"id": "` + id2 + `", "text": "` + invalidInput + `"}`)

	req, _ := http.NewRequest("POST", "/message/", bytes.NewBuffer(body))

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.UpdateMessage)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	// Assert HTTP response
	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("the id provided is invalid"))
}

func TestController_DeleteMessage(t *testing.T) {
	// Provision the msg in the collection
	Setup()

	//Create a new HTTP POST request
	req, _ := http.NewRequest("DELETE", "/message/1", nil)

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.DeleteMessage)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusNoContent
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}
}

func TestController_DeleteMessage2(t *testing.T) {
	//Create a new HTTP POST request
	req, _ := http.NewRequest("DELETE", "/message/11", nil)

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.DeleteMessage)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	// Assert HTTP response
	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("a message does not exists for the given Id: 11"))
}


func TestController_GetAllMessages(t *testing.T) {
	// Provision the msg in the collection
	Setup()

	//Create a new HTTP POST request
	req, _ := http.NewRequest("GET", "/messages", nil)

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.GetAllMessages)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusOK
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}
}

func TestController_GetMessageById(t *testing.T) {
	// Provision the msg in the collection
	Setup()

	//Create a new HTTP POST request
	req, _ := http.NewRequest("GET", "/message/11", nil)

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.GetMessageById)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusBadRequest
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var err ServiceError
	json.NewDecoder(io.Reader(response.Body)).Decode(&err)

	// Assert HTTP response
	g := NewGomegaWithT(t)
	g.Expect(err.Error).Should(Equal("a message does not exists for the given Id: 11"))
}

func TestController_GetMessageById2(t *testing.T) {
	// Provision the msg in the collection
	Setup()

	//Create a new HTTP POST request
	req, _ := http.NewRequest("GET", "/message/1", nil)

	//Assign HTTP Handler function to controller
	handler := http.HandlerFunc(msgCtrl.GetMessageById)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	expected := http.StatusOK
	status := response.Code

	if status != expected {
		t.Errorf("Handler returned a wrong status code. Expected: %v Actual: %v", expected, status)
	}

	// Decode the HTTP response
	var msg Message
	json.NewDecoder(io.Reader(response.Body)).Decode(&msg)

	// Assert HTTP response
	g := NewGomegaWithT(t)

	// Id should not be modified
	g.Expect(msg.Id).Should(Equal(id1))

	// Text must be equal to the body
	g.Expect(msg.Text).Should(Equal(validInput1))

	// And it should be an invalid palindrome
	g.Expect(msg.Palindrome).To(BeTrue())

	// Clean up collection
	TearDown(&msg)
}
