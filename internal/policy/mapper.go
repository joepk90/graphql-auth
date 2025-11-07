package policy

func MapAuthenticationResponse(errorMessage *string) *AuthenticationResponse {
	if errorMessage == nil {
		return &AuthenticationResponse{
			Success: true,
		}
	}

	return &AuthenticationResponse{
		ErrorMessage: errorMessage,
	}
}
