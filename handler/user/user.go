package user

import (
	"bookroom/models"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	Username string  `json:"username"`
	Password string `json:"password"`
	Repassword string
	IsAdmin bool `json:"isadmin"`
	CreateAt string `json:"create_at"`
}

type ListRequest struct {
	Username string `form:"username"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

type ListResponse struct {
	TotalCount int64            `json:"totalCount"`
	UserList   []*models.User `json:"userList"`
}