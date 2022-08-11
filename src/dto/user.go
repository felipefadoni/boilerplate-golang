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
