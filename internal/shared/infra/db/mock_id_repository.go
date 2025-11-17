package db

import "context"

type MockIDRepository struct {
	next    int
	AvailErr error
}

func (m *MockIDRepository) AvailableID(ctx context.Context, table TableName) (int, error) {
	if m.AvailErr != nil {
		return 0, m.AvailErr
	}
	m.next++
	return m.next, nil
}