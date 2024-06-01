package repository

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type contactRepository struct {
	conn   *pgx.Conn
	logger *zap.Logger
}

func NewContactRepository(conn *pgx.Conn, logger *zap.Logger) repository.ContactRepository {
	return contactRepository{
		conn:   conn,
		logger: logger,
	}
}
func (c contactRepository) GetAllContacts(ctx context.Context) ([]domain.Contact, error) {
	var result []contact
	query, err := c.conn.Query(ctx, "SELECT c.id, c.full_name, c.phone_number, c.email, c.created_at FROM contacts c")
	if err != nil {
		c.logger.Error("Error getting contacts", zap.Error(err))
		return nil, err
	}
	defer query.Close()
	for query.Next() {
		var contactResult contact
		err = query.Scan(&contactResult.ID, &contactResult.FullName, &contactResult.PhoneNumber, &contactResult.Email, &contactResult.CreatedAt)
		if err != nil {
			c.logger.Error("Error scanning contact", zap.Error(err))
			return nil, err
		}
		result = append(result, contactResult)
	}
	return c.mappers(result), nil
}
func (c contactRepository) mappers(source []contact) []domain.Contact {
	response := make([]domain.Contact, 0)
	for _, s := range source {
		response = append(response, c.mapper(s))
	}
	return response
}

func (c contactRepository) mapper(contactRepository contact) domain.Contact {
	return domain.Contact{
		ID:          contactRepository.ID,
		FullName:    contactRepository.FullName,
		PhoneNumber: contactRepository.PhoneNumber,
		Email:       contactRepository.Email,
		CreatedAt:   contactRepository.CreatedAt,
	}
}
