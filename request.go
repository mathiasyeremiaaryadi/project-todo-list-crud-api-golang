package main

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
