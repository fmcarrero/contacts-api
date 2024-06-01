package src

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/application/command"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/controller"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/repository"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type Dependencies struct {
	ContactHandler controller.ContactHandler
	Config         Config
	Logger         *zap.Logger
	conn           *pgx.Conn
}

func (d Dependencies) Close() {
	if d.conn != nil {
		err := d.conn.Close(context.Background())
		if err != nil {
			d.Logger.Error("Error closing connection", zap.Error(err))
		}
	}
}

func Build() Dependencies {
	dependencies := Dependencies{}
	dependencies.Config = NewConfig()
	dependencies.Logger, _ = zap.NewProduction()
	dependencies.conn = getConn(dependencies.Config, dependencies.Logger)

	contactRepository := repository.NewContactRepository(dependencies.conn, dependencies.Logger)

	getAllContacts := command.NewGetAllContacts(contactRepository)
	dependencies.ContactHandler = controller.NewContactHandler(getAllContacts)
	return dependencies
}
func getConn(cfg Config, logger *zap.Logger) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), cfg.Database.URL)
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	return conn
}
