// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dmosyan/Learning-Go/apis/banking/domain (interfaces: AccountRepository)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/domain/account_repository_mock.go -package=domain github.com/dmosyan/Learning-Go/apis/banking/domain AccountRepository
//

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	domain "github.com/dmosyan/Learning-Go/apis/banking/domain"
	errs "github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
	isgomock struct{}
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// FindBy mocks base method.
func (m *MockAccountRepository) FindBy(accountId string) (*domain.Account, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBy", accountId)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindBy indicates an expected call of FindBy.
func (mr *MockAccountRepositoryMockRecorder) FindBy(accountId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBy", reflect.TypeOf((*MockAccountRepository)(nil).FindBy), accountId)
}

// Save mocks base method.
func (m *MockAccountRepository) Save(arg0 domain.Account) (*domain.Account, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockAccountRepositoryMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAccountRepository)(nil).Save), arg0)
}

// SaveTransaction mocks base method.
func (m *MockAccountRepository) SaveTransaction(transaction domain.Transaction) (*domain.Transaction, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTransaction", transaction)
	ret0, _ := ret[0].(*domain.Transaction)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// SaveTransaction indicates an expected call of SaveTransaction.
func (mr *MockAccountRepositoryMockRecorder) SaveTransaction(transaction any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTransaction", reflect.TypeOf((*MockAccountRepository)(nil).SaveTransaction), transaction)
}
