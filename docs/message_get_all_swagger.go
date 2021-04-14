package docs

import . "github.com/ash822/goweb/entity"

// swagger:route GET /messages message
// Get all messages
// responses:
//   200: MessagesResponse

// swagger:response MessagesResponse
type MessagesResponseWrapper struct {
	// in:body
	Body []MessageResponse
}