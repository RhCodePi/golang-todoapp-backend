package dtos

import uuid "github.com/satori/go.uuid"

type CreateTodoMessageCommandResponse struct {
	TodoMessage uuid.UUID `json:"todomessageid"`
	TodolistId  uuid.UUID `json:"todolistid"`
}
