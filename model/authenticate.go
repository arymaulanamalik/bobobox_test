package model

type Error struct {
	Error   error
	Message string
	Code    string
}

type AuthenticateRequest struct {
	Username string
	Password string
}
type AuthenticateResponse struct {
	Token        string
	ExpiredAt    int
	RefreshToken string
}
