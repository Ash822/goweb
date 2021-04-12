package repository

import (
	"errors"
	"github.com/ash822/goweb/entity"
)

type MessageRepository interface {
	FindById(id string) (*entity.Message, error)
	FindAll() ([]entity.Message, error)
	Create(msg *entity.Message) (*entity.Message, error)
	Update(newMsg *entity.Message) (*entity.Message, error)
	Delete(id string) error
}

var (
	msgs []entity.Message
)

type repo struct{}

func GetInstance() MessageRepository {
	return &repo{}
}

func (*repo) Create(msg *entity.Message) (*entity.Message, error) {
	msgs = append(msgs, entity.Message{
		Id:         msg.Id,
		Text:       msg.Text,
		Palindrome: msg.Palindrome,
	})

	return msg, nil
}

func (*repo) Update(newMsg *entity.Message) (*entity.Message, error) {
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
		msgs[index] = entity.Message{
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
		return errors.New("a message does not exists for the given Id:" + id)
	} else {
		msgs[index] = msgs[len(msgs)-1]
		msgs[len(msgs)-1] = entity.Message{}
		msgs = msgs[:len(msgs)-1]

		return nil
	}
}

func (*repo) FindById(id string) (*entity.Message, error) {
	if id == "" {
		return nil, errors.New("the id provided is invalid")
	}

	for _, msg := range msgs {
		if id == msg.Id {
			return &msg, nil
		}
	}

	return nil, errors.New("a message does not exists for the given Id:" + id)
}

func (*repo) FindAll() ([]entity.Message, error) {
	if len(msgs) == 0 {
		msgs = []entity.Message{}
	}
	return msgs, nil
}
