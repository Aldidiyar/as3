package contact

import (
	"arch/services/contact/internal/domain"
	"context"
	"database/sql"
)

func NewContactRepository(db *sql.DB) *contactRepository {
	return &contactRepository{
		db: db,
	}
}

type contactRepository struct {
	db *sql.DB
}

func (r *contactRepository) Create(ctx context.Context, contact domain.Contact) error {
	stmt, err := r.db.Prepare("INSERT INTO contacts (name, email, phone) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, contact.FullName, contact.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (r *contactRepository) Get(ctx context.Context, id int) (domain.Contact, error) {
	stmt, err := r.db.Prepare("SELECT id, name, email, phone FROM contacts WHERE id = $1")
	if err != nil {
		return domain.Contact{}, err
	}

	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, id)

	var contact domain.Contact
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.PhoneNumber)
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *contactRepository) Update(ctx context.Context, contact domain.Contact) error {
	stmt, err := r.db.Prepare("UPDATE contacts SET name = $1, email = $2, phone = $3 WHERE id = $4")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, contact.FullName, contact.PhoneNumber, contact.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *contactRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare("DELETE FROM contacts WHERE id = $1")
	if err != nil {
		return err
	}

	// Выполнение запроса
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
