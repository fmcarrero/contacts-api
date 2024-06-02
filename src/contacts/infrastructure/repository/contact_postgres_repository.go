package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	customError "github.com/fmcarrero/contacts-api/src/contacts/domain/error"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type contactRepository struct {
	conn   *pgxpool.Pool
	logger *zap.Logger
}

func NewContactRepository(conn *pgxpool.Pool, logger *zap.Logger) repository.ContactRepository {
	return contactRepository{
		conn:   conn,
		logger: logger,
	}
}

func (c contactRepository) AddContact(ctx context.Context, contact domain.Contact) error {
	_, err := c.conn.Exec(ctx,
		"insert into contacts(id, full_name, phone_number, email, created_at, update_at) values($1, $2, $3, $4, $5, $6)",
		contact.ID, contact.FullName, contact.PhoneNumber, contact.Email, contact.CreatedAt, contact.UpdateAt,
	)
	if err != nil {
		c.logger.Error("Error adding contact", zap.Error(err), zap.Any("contact", contact))
		return err
	}
	return nil
}

func (c contactRepository) RemoveContact(ctx context.Context, id int64) error {
	_, err := c.conn.Exec(ctx, "delete from contacts where id=$1", id)
	if err != nil {
		c.logger.Error("Error removing contact", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}
func (c contactRepository) GetContactByID(ctx context.Context, id int64) (domain.Contact, error) {
	var contactResult contact
	err := c.conn.QueryRow(ctx, "SELECT c.id, c.full_name, c.phone_number, c.email, c.created_at FROM contacts c WHERE c.id=$1", id).Scan(
		&contactResult.ID, &contactResult.FullName, &contactResult.PhoneNumber, &contactResult.Email, &contactResult.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Contact{}, customError.NewContactNotFoundError(fmt.Sprintf("Contact with id %d not found", id), "contact.not_found")
		}
		c.logger.Error("Error getting contact", zap.Error(err), zap.Int64("id", id))
		return domain.Contact{}, err
	}
	return c.mapper(contactResult), nil

}

func (c contactRepository) EditContact(ctx context.Context, contact domain.Contact) error {
	_, err := c.conn.Exec(ctx,
		"update contacts set full_name=$2, phone_number=$3, email=$4, update_at=$5 where id=$1",
		contact.ID, contact.FullName, contact.PhoneNumber, contact.Email, contact.UpdateAt,
	)
	if err != nil {
		c.logger.Error("Error updating contact", zap.Error(err), zap.Int64("id", contact.ID))
		return err
	}
	return nil
}
func (c contactRepository) GetAllContacts(ctx context.Context) ([]domain.Contact, error) {
	var result []contact
	query, err := c.conn.Query(ctx, `SELECT c.id, c.full_name, c.phone_number, c.email, 
       										c.created_at , c.update_at
											FROM contacts c
											order by c.full_name asc`)
	if err != nil {
		c.logger.Error("Error getting contacts", zap.Error(err))
		return nil, err
	}
	defer query.Close()
	for query.Next() {
		var contactResult contact
		err = query.Scan(&contactResult.ID, &contactResult.FullName, &contactResult.PhoneNumber, &contactResult.Email,
			&contactResult.CreatedAt, &contactResult.UpdateAt)
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
		UpdateAt:    contactRepository.UpdateAt,
	}
}
