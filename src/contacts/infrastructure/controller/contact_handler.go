package controller

import (
	"fmt"
	"github.com/fmcarrero/contacts-api/src/contacts/application/command"
	"github.com/fmcarrero/contacts-api/src/contacts/application/query"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/controller/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ContactHandler interface {
	GetContacts(ctx echo.Context) error
	EditContact(ctx echo.Context) error
	AddContact(ctx echo.Context) error
	RemoveContact(ctx echo.Context) error
}

type contactHandler struct {
	GetAllContactsQuery  query.GetAllContactsQuery
	EditContactCommand   command.EditContactCommand
	AddContactCommand    command.AddContactCommand
	RemoveContactCommand command.RemoveContactCommand
}

func NewContactHandler(getAllContactsQuery query.GetAllContactsQuery,
	editContactCommand command.EditContactCommand,
	addContactCommand command.AddContactCommand,
	removeContactCommand command.RemoveContactCommand) ContactHandler {
	return &contactHandler{
		GetAllContactsQuery:  getAllContactsQuery,
		EditContactCommand:   editContactCommand,
		AddContactCommand:    addContactCommand,
		RemoveContactCommand: removeContactCommand,
	}
}
func (c contactHandler) RemoveContact(ctx echo.Context) error {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid id %s, id should be numeric", ctx.Param("id")))
	}
	contact, err := c.RemoveContactCommand.Executess(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, contact)
}

func (c contactHandler) AddContact(ctx echo.Context) error {
	var addContactRequest command.AddContactRequest
	err := ctx.Bind(&addContactRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	contact, err := c.AddContactCommand.Executes(ctx.Request().Context(), addContactRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, contact)
}
func (c contactHandler) EditContact(ctx echo.Context) error {
	var editContactRequest command.EditContactRequest
	err := ctx.Bind(&editContactRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	contact, err := c.EditContactCommand.Execute(ctx.Request().Context(), editContactRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, contact)
}
func (c contactHandler) GetContacts(ctx echo.Context) error {
	contacts, err := c.GetAllContactsQuery.Execute(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, dto.ToContactsDTO(contacts))
}
