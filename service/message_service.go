package service

import (
	"errors"
	"github.com/ash822/goweb/entity"
	"github.com/ash822/goweb/repository"
	"github.com/google/uuid"
)

type MessageService interface {
	IsPalindrome(text string) (bool, error)
	Create(message *entity.Message) (*entity.Message, error)
	Update(id string, message *entity.Message) (*entity.Message, error)
	Delete(id string) error
	FindById(id string) (*entity.Message, error)
	FindAll() ([]entity.Message, error)
}

var repo repository.MessageRepository

type svc struct{}

func GetInstance(msgRepo repository.MessageRepository) MessageService {
	repo = msgRepo
	return &svc{}
}

func (*svc) IsPalindrome(text string) (bool, error) {
	if text == "" {
		err := errors.New("text provided is empty")
		return false, err
	} else {
		// TODO: Check if palindrome
		return false, nil
	}
}

func (*svc) Create(msg *entity.Message) (*entity.Message, error) {
	msg.Id = uuid.New().String()
	return repo.Create(msg)
}

func (*svc) Update(id string, newMsg *entity.Message) (*entity.Message, error) {
	newMsg.Id = id
	newMsg.Palindrome = false // TODO: Check for palindrome
	return repo.Update(newMsg)
}

func (*svc) Delete(id string) error {
	return repo.Delete(id)
}

func (*svc) FindById(id string) (*entity.Message, error) {
	return repo.FindById(id)
}

func (*svc) FindAll() ([]entity.Message, error) {
	return repo.FindAll()
}
