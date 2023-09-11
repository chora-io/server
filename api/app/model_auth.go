package app

type PostAuthRequest struct {
	// Address is the address.
	Address string `json:"address"`

	// Token is the authentication token.
	Token string `json:"token"`
}

type PostAuthResponse struct {
	// Authenticated is true if authenticated.
	Authenticated bool `json:"authenticated"`
}

func NewPostAuthResponse(authenticated bool) PostAuthResponse {
	return PostAuthResponse{
		Authenticated: authenticated,
	}
}

type PostAuthKeplrRequest struct {
	// Address is the address.
	Address string `json:"address"`

	// Signature is the signature.
	Signature string `json:"signature"`
}

type PostAuthKeplrResponse struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// Authenticated is true if authenticated.
	Authenticated bool `json:"authenticated"`
}

func NewPostAuthKeplrResponse(token string, authenticated bool) PostAuthKeplrResponse {
	return PostAuthKeplrResponse{
		Token:         token,
		Authenticated: authenticated,
	}
}

type PostAuthLoginRequest struct {
	// Username is the username.
	Username string `json:"username"`

	// Password is the password.
	Password string `json:"password"`
}

type PostAuthLoginResponse struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// Authenticated is true if authenticated.
	Authenticated bool `json:"authenticated"`
}

func NewPostAuthLoginResponse(token string, authenticated bool) PostAuthLoginResponse {
	return PostAuthLoginResponse{
		Token:         token,
		Authenticated: authenticated,
	}
}
