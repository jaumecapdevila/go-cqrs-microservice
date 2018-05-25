package db

import (
	"context"

	"github.com/jaumecapdevila/go-cqrs-microservice/schema"
)

type Repository interface {
	Close()
	InsertMessage(ctx context.Context, meow schema.Message) error
	ListMessages(ctx context.Context, skip uint64, take uint64) ([]schema.Message, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertMeow(ctx context.Context, message schema.Message) error {
	return impl.InsertMessage(ctx, message)
}

func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Message, error) {
	return impl.ListMessages(ctx, skip, take)
}
