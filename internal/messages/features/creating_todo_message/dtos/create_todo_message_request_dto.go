package dtos

import uuid "github.com/satori/go.uuid"

type CreateTodoMessageRequestDto struct {
	TodolistId uuid.UUID `json:"todolistid"`
	Content    string    `json:"content"`
}
