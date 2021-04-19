package docs

import . "github.com/ash822/goweb/entity"

// swagger:route GET /messages message
// Get all messages
// responses:
//   200: MessagesResponse
//	 401: ErrorResponse

// swagger:response MessagesResponse
type MessagesResponseWrapper struct {
	// in:body
	Body []MessageResponse
}