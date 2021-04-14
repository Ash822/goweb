package entity

type MessageRequest struct {
	Text       string `json:"text"`
}

type MessageResponse struct {
	Id         string `json:"id"`
	Text       string `json:"text"`
	Palindrome bool   `json:"palindrome"`
}
