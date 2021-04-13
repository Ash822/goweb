package service

import (
	"errors"
	. "github.com/ash822/goweb/entity"
	"github.com/ash822/goweb/repository"
	"github.com/ash822/goweb/utils"
	"github.com/google/uuid"
)

type MessageService interface {
	Create(message *Message) (*Message, error)
	Update(id string, message *Message) (*Message, error)
	Delete(id string) error
	FindById(id string) (*Message, error)
	FindAll() ([]Message, error)
}

var repo repository.MessageRepository

type svc struct{}

func GetInstance(msgRepo repository.MessageRepository) MessageService {
	repo = msgRepo
	return &svc{}
}

func (*svc) Create(msg *Message) (*Message, error) {
	if msg.Text == "" {
		return nil, errors.New("the message text is invalid or not found")
	}

	msg.Id = uuid.New().String()
	msg.Palindrome = msgutils.IsPalindrome(msg.Text)
	return repo.Create(msg)
}

func (*svc) Update(id string, newMsg *Message) (*Message, error) {
	if id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	if newMsg.Text == "" {
		return nil, errors.New("the message text is invalid or not found")
	}

	newMsg.Id = id
	newMsg.Palindrome = msgutils.IsPalindrome(newMsg.Text)
	return repo.Update(newMsg)
}

func (*svc) Delete(id string) error {
	if id == "" {
		return errors.New("the id provided is invalid")
	}

	return repo.Delete(id)
}

func (*svc) FindById(id string) (*Message, error) {
	if id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	return repo.FindById(id)
}

func (*svc) FindAll() ([]Message, error) {
	return repo.FindAll()
}
