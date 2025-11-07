package policy

type AuthenticationResponse struct {
	Success      bool    `json:"success"`
	ErrorMessage *string `json:"error_message,omitempty"`
}
