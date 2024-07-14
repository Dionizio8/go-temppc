// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entity "github.com/Dionizio8/go-temppc/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockAddressRepository is an autogenerated mock type for the AddressRepositoryInterface type
type MockAddressRepository struct {
	mock.Mock
}

type MockAddressRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAddressRepository) EXPECT() *MockAddressRepository_Expecter {
	return &MockAddressRepository_Expecter{mock: &_m.Mock}
}

// GetAddress provides a mock function with given fields: zipCode
func (_m *MockAddressRepository) GetAddress(zipCode string) (entity.Address, error) {
	ret := _m.Called(zipCode)

	if len(ret) == 0 {
		panic("no return value specified for GetAddress")
	}

	var r0 entity.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Address, error)); ok {
		return rf(zipCode)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Address); ok {
		r0 = rf(zipCode)
	} else {
		r0 = ret.Get(0).(entity.Address)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(zipCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAddressRepository_GetAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddress'
type MockAddressRepository_GetAddress_Call struct {
	*mock.Call
}

// GetAddress is a helper method to define mock.On call
//   - zipCode string
func (_e *MockAddressRepository_Expecter) GetAddress(zipCode interface{}) *MockAddressRepository_GetAddress_Call {
	return &MockAddressRepository_GetAddress_Call{Call: _e.mock.On("GetAddress", zipCode)}
}

func (_c *MockAddressRepository_GetAddress_Call) Run(run func(zipCode string)) *MockAddressRepository_GetAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAddressRepository_GetAddress_Call) Return(_a0 entity.Address, _a1 error) *MockAddressRepository_GetAddress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAddressRepository_GetAddress_Call) RunAndReturn(run func(string) (entity.Address, error)) *MockAddressRepository_GetAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAddressRepository creates a new instance of MockAddressRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAddressRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAddressRepository {
	mock := &MockAddressRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}