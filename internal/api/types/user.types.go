package types

type GetUser struct {
	Name string `json:"name"`
	UUID string `json:"identifier"`
}

type GetUserWithID struct {
	ID   uint
	Name string `json:"name"`
	UUID string `json:"identifier"`
}

type AddUser struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
