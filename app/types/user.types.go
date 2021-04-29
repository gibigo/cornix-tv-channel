package types

type GetUser struct {
	ID   uint
	Name string `json:"name"`
}

type AddUser struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
