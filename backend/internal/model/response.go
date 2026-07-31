package model

// APIResponse is the standard structure for all API response bodies.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// HelloData represents the structure of the data returned by the hello handler.
type HelloData struct {
	Message string   `json:"message"`
	Topics  []string `json:"topics"`
	Version string   `json:"version"`
}
