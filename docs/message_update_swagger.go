package docs

// swagger:route POST /message/{id} message updateEndpoint
// Update the message by Id
// responses:
//   200: MessageResponse
//	 400: ErrorResponse
//	 401: ErrorResponse

// swagger:parameters updateEndpoint
type MessageUpdateParamsWrapper struct {
	// in:path
	ID string `json:"id"`
	// Add the string that needs palindrome check to the Body property 'text'
	// in:body
	Body struct {
		// Required: true
		// Example: Expected type string
		Text string
	}
}
