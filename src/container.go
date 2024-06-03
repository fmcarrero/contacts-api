package src

import (
	"context"

	"github.com/fmcarrero/contacts-api/src/contacts/application/command"
	"github.com/fmcarrero/contacts-api/src/contacts/application/query"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/controller"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/repository"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/zap"
)

type Dependencies struct {
	ContactHandler controller.ContactHandler
	Config         Config
	Logger         *zap.Logger
	Conn           *pgxpool.Pool
}

func (d Dependencies) CloseDatabase() {
	if d.Conn != nil {
		d.Conn.Close()
	}
}

func Build() Dependencies {
	dependencies := Dependencies{}
	dependencies.Config = NewConfig()
	dependencies.Logger, _ = zap.NewProduction()
	dependencies.Conn = GetConn(dependencies.Config, dependencies.Logger)

	contactRepository := repository.NewContactRepository(dependencies.Conn, dependencies.Logger)

	getAllContacts := query.NewGetAllContacts(contactRepository)
	editContactCommand := command.NewEditContact(contactRepository)
	addContactCommand := command.NewAddContact(contactRepository)
	removeContactCommand := command.NewRemoveContact(contactRepository)
	dependencies.ContactHandler = controller.NewContactHandler(getAllContacts, editContactCommand,
		addContactCommand, removeContactCommand)
	return dependencies
}
func GetConn(cfg Config, logger *zap.Logger) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	return conn
}
