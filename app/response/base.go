package response

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}