package dtos

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
