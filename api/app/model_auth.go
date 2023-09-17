package app

import "time"

type PostAuthRequest struct {
	// Token is the authentication token.
	Token string `json:"token"`
}

type PostAuthResponse struct {
	// User is the authenticated user.
	User UserInfo `json:"user"`
}

func NewPostAuthResponse(user UserInfo) PostAuthResponse {
	return PostAuthResponse{
		User: user,
	}
}

type PostAuthEmailRequest struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// Email is the email.
	Email string `json:"email"`

	// AccessCode is the access_code.
	AccessCode string `json:"access_code"`
}

type PostAuthEmailResponse struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// User is the authenticated user.
	User UserInfo `json:"user"`
}

func NewPostAuthEmailResponse(token string, user UserInfo) PostAuthEmailResponse {
	return PostAuthEmailResponse{
		Token: token,
		User:  user,
	}
}

type PostAuthKeplrRequest struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// Address is the address.
	Address string `json:"address"`

	// Signature is the signature.
	Signature string `json:"signature"`
}

type PostAuthKeplrResponse struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// User is the authenticated user.
	User UserInfo `json:"user"`
}

func NewPostAuthKeplrResponse(token string, user UserInfo) PostAuthKeplrResponse {
	return PostAuthKeplrResponse{
		Token: token,
		User:  user,
	}
}

type PostAuthLoginRequest struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// Username is the username.
	Username string `json:"username"`

	// Password is the password.
	Password string `json:"password"`
}

type PostAuthLoginResponse struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// User is the authenticated user.
	User UserInfo `json:"user"`
}

func NewPostAuthLoginResponse(token string, user UserInfo) PostAuthLoginResponse {
	return PostAuthLoginResponse{
		Token: token,
		User:  user,
	}
}

type UserInfo struct {
	// Id is the authenticated user id.
	Id string `json:"id"`

	// Email is the authenticated user email.
	Email string `json:"email"`

	// Address is the authenticated user address.
	Address string `json:"address"`

	// Username is the authenticated user username.
	Username string `json:"username"`

	// CreatedAt is when the authenticated user was created.
	CreatedAt time.Time `json:"created_at"`
}
