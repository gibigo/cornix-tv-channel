package types

type AddUser struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name     string    `json:"username"`
	Channels []Channel `json:"channels"`
}
