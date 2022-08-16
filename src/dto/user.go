package dto

type GetAllUserDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Accepted  bool   `json:"accepted"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Accepted bool   `json:"accepted"`
}

type CreateUserReturnDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Accepted  bool   `json:"accepted"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
