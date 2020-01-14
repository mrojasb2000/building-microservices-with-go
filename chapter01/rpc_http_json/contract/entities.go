package contract

// HelloWorldRequest request type
type HelloWorldRequest struct {
	Name string
}

// HelloWorldResponse reponse type
type HelloWorldResponse struct {
	Message string `json:"message"`
}
