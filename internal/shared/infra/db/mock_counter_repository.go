package db

import "context"

type MockCounterRepository struct {
	next int
}

func (m *MockCounterRepository) NextID(ctx context.Context, table TableName) (int, error) {
	m.next++
	return m.next, nil
}