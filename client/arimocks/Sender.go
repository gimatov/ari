// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/gimatov/ari/v5"
	mock "github.com/stretchr/testify/mock"
)

// Sender is an autogenerated mock type for the Sender type
type Sender struct {
	mock.Mock
}

// Send provides a mock function with given fields: e
func (_m *Sender) Send(e ari.Event) {
	_m.Called(e)
}
