// Code generated by mockery v2.46.0. DO NOT EDIT.

package identitymocks

import mock "github.com/stretchr/testify/mock"

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// NewCallbacks creates a new instance of Callbacks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCallbacks(t interface {
	mock.TestingT
	Cleanup(func())
}) *Callbacks {
	mock := &Callbacks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
