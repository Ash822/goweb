package service

import (
	"errors"
	. "github.com/ash822/goweb/entity"
	"github.com/ash822/goweb/repository"
	"github.com/ash822/goweb/utils"
)

type MessageService interface {
	Create(message *MessageRequest) (*MessageResponse, error)
	Update(id string, message *MessageRequest) (*MessageResponse, error)
	Delete(id string) error
	FindById(id string) (*MessageResponse, error)
	FindAll() ([]MessageResponse, error)
}

var repo repository.MessageRepository

type svc struct{}

func GetInstance(msgRepo repository.MessageRepository) MessageService {
	repo = msgRepo
	return &svc{}
}

func (*svc) Create(msg *MessageRequest) (*MessageResponse, error) {
	if msg.Text == "" {
		return nil, errors.New("the message text is invalid or not found")
	}

	var msgRes MessageResponse

	msgRes.Text = msg.Text
	msgRes.Palindrome = msgutils.IsPalindrome(msg.Text)
	return repo.Create(&msgRes)
}

func (*svc) Update(id string, msg *MessageRequest) (*MessageResponse, error) {
	if id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	if msg.Text == "" {
		return nil, errors.New("the message text is invalid or not found")
	}

	var newMsg MessageResponse

	newMsg.Id = id
	newMsg.Text = msg.Text
	newMsg.Palindrome = msgutils.IsPalindrome(newMsg.Text)
	return repo.Update(&newMsg)
}

func (*svc) Delete(id string) error {
	if id == "" {
		return errors.New("the id provided is invalid")
	}

	return repo.Delete(id)
}

func (*svc) FindById(id string) (*MessageResponse, error) {
	if id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	return repo.FindById(id)
}

func (*svc) FindAll() ([]MessageResponse, error) {
	return repo.FindAll()
}
