package db

import (
	"context"
	"database/sql"

	"github.com/jaumecapdevila/go-cqrs-microservice/schema"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertMessage(ctx context.Context, message schema.Message) error {
	_, err := r.db.Exec("INSERT INTO messages(id, body, created_at) VALUES($1, $2, $3)", message.ID, message.Body, message.CreatedAt)
	return err
}

func (r *PostgresRepository) ListMessages(ctx context.Context, skip uint64, take uint64) ([]schema.Message, error) {
	rows, err := r.db.Query("SELECT * FROM messages ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Messages
	messages := []schema.Message{}
	for rows.Next() {
		message := schema.Message{}
		if err = rows.Scan(&message.ID, &message.Body, &message.CreatedAt); err == nil {
			messages = append(messages, message)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
