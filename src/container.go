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
	conn           *pgxpool.Pool
}

func (d Dependencies) Close() {
	if d.conn != nil {
		d.conn.Close()
	}
}

func Build() Dependencies {
	dependencies := Dependencies{}
	dependencies.Config = NewConfig()
	dependencies.Logger, _ = zap.NewProduction()
	dependencies.conn = getConn(dependencies.Config, dependencies.Logger)

	contactRepository := repository.NewContactRepository(dependencies.conn, dependencies.Logger)

	getAllContacts := query.NewGetAllContacts(contactRepository)
	editContactCommand := command.NewEditContact(contactRepository)
	addContactCommand := command.NewAddContact(contactRepository)
	dependencies.ContactHandler = controller.NewContactHandler(getAllContacts, editContactCommand, addContactCommand)
	return dependencies
}
func getConn(cfg Config, logger *zap.Logger) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	return conn
}
