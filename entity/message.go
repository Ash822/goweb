package entity

type Message struct {
	Id         string `json:"id"`
	Text       string `json:"text"`
	Palindrome bool   `json:"palindrome"`
}
