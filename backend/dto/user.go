package dto

type Login struct {
	Password string `json:"hashPassword"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	UserId   string `json:"userId"`
	Roles    string `json:"roles"`
	UserName string `json:"userName"`
	Password string `json:"hashedPassword"`
	Expired  int64  `json:"expired"`
}

type CreateUser struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
	Fullname       string `json:"fullname"`
}


