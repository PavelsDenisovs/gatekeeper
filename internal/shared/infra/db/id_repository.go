package db

import "context"

type IDRepository interface {
	AvailableID(ctx context.Context, table TableName) (int, error)
}