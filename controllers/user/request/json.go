package request

import "fgd/core/user"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenRequest struct {
	RefreshToken string `JSON:"refresh_token"`
}

func (r *User) ToDomain() *user.Domain {
	return &user.Domain{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}
