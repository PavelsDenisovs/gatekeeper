package db

import "context"

type CounterRepository interface {
	NextID(ctx context.Context, table TableName) (int, error)
}