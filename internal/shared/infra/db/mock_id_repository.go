package db

import "context"

type MockIDRepository struct {
	next    int
	nextErr error
}

func (m *MockIDRepository) AvailableID(ctx context.Context, table TableName) (int, error) {
	if m.nextErr != nil {
		return 0, m.nextErr
	}
	m.next++
	return m.next, nil
}