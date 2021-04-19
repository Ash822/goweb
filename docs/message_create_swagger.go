package docs

import . "github.com/ash822/goweb/entity"

// swagger:route POST /message message bodyOfMessageEndpoint
// Check if the given text is palindrome
// responses:
//   201: MessageResponse
//	 400: ErrorResponse
//	 401: ErrorResponse

// swagger:response MessageResponse
type MessageResponseWrapper struct {
	// in:body
	Body MessageResponse
}

// swagger:response ErrorResponse
type ErrorResponseWrapper struct {
	// in:body
	Body ServiceError
}

// swagger:parameters bodyOfMessageEndpoint
type MessageParamsWrapper struct {
	// Add the string that needs palindrome check to the Body property 'text'
	// in:body
	Body struct {
		// Required: true
		// Example: Expected type string
		Text string
	}
}
