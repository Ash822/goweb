package docs

// swagger:route DELETE /message/{id} message deleteMessageId
// Deletes a message by Id
// responses:
//   204: EmptyResponse
//   400: ErrorResponse

// swagger:response EmptyResponse
type EmptyResponse struct {

}

// swagger:parameters deleteMessageId
type DeleteMessageIdParamsWrapper struct {
	// in:path
	ID string `json:"id"`
}
