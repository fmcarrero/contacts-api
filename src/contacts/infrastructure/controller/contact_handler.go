package controller

import (
	"github.com/fmcarrero/contacts-api/src/contacts/application/command"
	"github.com/fmcarrero/contacts-api/src/contacts/application/query"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/controller/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactHandler interface {
	GetContacts(ctx echo.Context) error
	EditContact(ctx echo.Context) error
	AddContact(ctx echo.Context) error
}

type contactHandler struct {
	GetAllContactsQuery query.GetAllContactsQuery
	EditContactCommand  command.EditContactCommand
	AddContactCommand   command.AddContactCommand
}

func NewContactHandler(getAllContactsQuery query.GetAllContactsQuery,
	editContactCommand command.EditContactCommand,
	addContactCommand command.AddContactCommand) ContactHandler {
	return &contactHandler{
		GetAllContactsQuery: getAllContactsQuery,
		EditContactCommand:  editContactCommand,
		AddContactCommand:   addContactCommand,
	}
}

func (c contactHandler) AddContact(ctx echo.Context) error {
	var addContactRequest command.AddContactRequest
	err := ctx.Bind(&addContactRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	contact, err := c.AddContactCommand.Executes(ctx.Request().Context(), addContactRequest)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, contact)
}
func (c contactHandler) EditContact(ctx echo.Context) error {
	var editContactRequest command.EditContactRequest
	err := ctx.Bind(&editContactRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	contact, err := c.EditContactCommand.Execute(ctx.Request().Context(), editContactRequest)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, contact)
}
func (c contactHandler) GetContacts(ctx echo.Context) error {
	contacts, err := c.GetAllContactsQuery.Execute(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, dto.ToContactsDTO(contacts))
}
