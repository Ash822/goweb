package docs

// swagger:route GET /message/{id} message idOfMessageEndpoint
// Get a message by id
// responses:
//   200: MessageResponse
//	 400: ErrorResponse
//	 401: ErrorResponse

// swagger:parameters idOfMessageEndpoint
type MessageIdParamsWrapper struct {
	// in:path
	ID string `json:"id"`
}

