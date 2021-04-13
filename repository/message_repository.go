package repository

//go:generate mockgen -destination=../mocks/repository/message_repository.go -package=mocks github.com/ash822/goweb/repository MessageRepository

import (
	"errors"
	. "github.com/ash822/goweb/entity"
)

type MessageRepository interface {
	FindById(id string) (*Message, error)
	FindAll() ([]Message, error)
	Create(msg *Message) (*Message, error)
	Update(newMsg *Message) (*Message, error)
	Delete(id string) error
}

var (
	msgs []Message
)

type repo struct{}

func GetInstance() MessageRepository {
	return &repo{}
}

func (*repo) Create(msg *Message) (*Message, error) {
	msgs = append(msgs, Message{
		Id:         msg.Id,
		Text:       msg.Text,
		Palindrome: msg.Palindrome,
	})

	return msg, nil
}

func (*repo) Update(newMsg *Message) (*Message, error) {
	if newMsg.Id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	var index = -1
	for i, msg := range msgs {
		if newMsg.Id == msg.Id {
			index = i
		}
	}

	if index == -1 {
		return nil, errors.New("a message does not exists for the given Id:" + newMsg.Id)
	} else {
		msgs[index] = Message{
			Id:         newMsg.Id,
			Text:       newMsg.Text,
			Palindrome: newMsg.Palindrome,
		}

		return newMsg, nil
	}
}

func (*repo) Delete(id string) error {
	if id == "" {
		return errors.New("the id provided is invalid")
	}

	var index = -1
	for i, msg := range msgs {
		if id == msg.Id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("a message does not exists for the given Id: " + id)
	} else {
		msgs[index] = msgs[len(msgs)-1]
		msgs[len(msgs)-1] = Message{}
		msgs = msgs[:len(msgs)-1]

		return nil
	}
}

func (*repo) FindById(id string) (*Message, error) {
	if id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	for _, msg := range msgs {
		if id == msg.Id {
			return &msg, nil
		}
	}

	return nil, errors.New("a message does not exists for the given Id: " + id)
}

func (*repo) FindAll() ([]Message, error) {
	if len(msgs) == 0 {
		msgs = []Message{}
	}
	return msgs, nil
}
