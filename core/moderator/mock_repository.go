// Code generated by mockery v2.14.0. DO NOT EDIT.

package moderator

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// ApplyPromotion provides a mock function with given fields: userId, topicId
func (_m *MockRepository) ApplyPromotion(userId uint, topicId uint) (Domain, error) {
	ret := _m.Called(userId, topicId)

	var r0 Domain
	if rf, ok := ret.Get(0).(func(uint, uint) Domain); ok {
		r0 = rf(userId, topicId)
	} else {
		r0 = ret.Get(0).(Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userId, topicId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApprovePromotion provides a mock function with given fields: promotionId
func (_m *MockRepository) ApprovePromotion(promotionId uint) error {
	ret := _m.Called(promotionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(promotionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPromotionRequest provides a mock function with given fields:
func (_m *MockRepository) GetPromotionRequest() ([]Domain, error) {
	ret := _m.Called()

	var r0 []Domain
	if rf, ok := ret.Get(0).(func() []Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectPromotion provides a mock function with given fields: promotionId
func (_m *MockRepository) RejectPromotion(promotionId uint) error {
	ret := _m.Called(promotionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(promotionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveModerator provides a mock function with given fields: userId, topicId
func (_m *MockRepository) RemoveModerator(userId uint, topicId uint) error {
	ret := _m.Called(userId, topicId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userId, topicId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
