package models

import "github.com/google/uuid"

type Users struct {
	UserId    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `jsom:"password"`
	UsersType string    `json:"userstype"`
}
