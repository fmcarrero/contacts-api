package controller

import (
	"github.com/fmcarrero/contacts-api/src/contacts/application/command"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/controller/dto"
	"github.com/labstack/echo/v4"
)

type ContactHandler interface {
	GetContacts(ctx echo.Context) error
}

type contactHandler struct {
	GetAllContactsCommand command.GetAllContactsCommand
}

func NewContactHandler(getAllContactsCommand command.GetAllContactsCommand) ContactHandler {
	return &contactHandler{
		GetAllContactsCommand: getAllContactsCommand,
	}
}
func (c contactHandler) GetContacts(ctx echo.Context) error {
	contacts, err := c.GetAllContactsCommand.Execute(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, dto.ToContactsDTO(contacts))
}
