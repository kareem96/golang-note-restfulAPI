package web

type UserResponse struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

type VerifyRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	Name string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	ID int`json:"-" validate:"required,max=100"`
	Password string`json:"password,omitempty" validate:"max=100"`
	Name string`json:"name,omitempty" validate:"max=100"`
}

type LoginRequest struct {
	ID int`json:"id" validate:"required,max=100"`
	Password string`json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID int`json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID int`json:"id" validate:"required,max=100"`
}