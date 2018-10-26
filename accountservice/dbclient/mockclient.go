package dbclient

import (
	"servicesdemo/accountservice/model"

	"github.com/stretchr/testify/mock"
)

// MockBoltClient is a mock implementation of a datastore client for testing purposes
type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {}

func (m *MockBoltClient) Seed() {}
