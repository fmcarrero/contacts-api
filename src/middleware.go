package src

import (
	"errors"
	customError "github.com/fmcarrero/contacts-api/src/contacts/domain/error"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Middleware func(*Server)

// Middlewares build the middlewares of the server
func (s *Server) Middlewares(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		middleware(s)
	}
}
func WithRecover() Middleware {
	return func(s *Server) {
		s.Server.Use(echoMiddleware.Recover())
	}
}
func WithRequestID() Middleware {
	return func(s *Server) {
		s.Server.Use(
			echoMiddleware.RequestIDWithConfig(echoMiddleware.RequestIDConfig{
				Generator: func() string {
					return uuid.New().String()
				},
			}),
		)
	}
}
func WithErrorHandler() Middleware {
	return func(s *Server) {
		s.Server.HTTPErrorHandler = func(err error, ctx echo.Context) {
			var contactNotFoundError customError.ContactNotFoundError
			var contactValidationError customError.ContactValidationError
			var echoError *echo.HTTPError
			switch {
			case errors.As(err, &contactValidationError):
				returned := NewError(http.StatusBadRequest, contactValidationError.Type, err)
				_ = ctx.JSON(returned.StatusCode, returned)
				return
			case errors.As(err, &contactNotFoundError):
				returned := NewError(http.StatusNotFound, contactNotFoundError.Type, err)
				_ = ctx.JSON(returned.StatusCode, returned)
				return

			case errors.As(err, &echoError):
				returned := NewError(echoError.Code, "echo.validations", echoError)
				_ = ctx.JSON(returned.StatusCode, returned)
				return
			default:
				_ = ctx.JSON(internalServerError.StatusCode, internalServerError)
				return
			}
		}
	}
}

var internalServerError = Error{StatusCode: http.StatusInternalServerError, Type: "internal_server_error", Message: "Internal server error"}

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
}

func NewError(statusCode int, errType string, err error) Error {
	return Error{StatusCode: statusCode, Type: errType, Message: err.Error()}
}
