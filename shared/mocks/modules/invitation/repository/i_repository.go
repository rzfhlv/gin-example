// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	mock "github.com/stretchr/testify/mock"

	param "github.com/rzfhlv/gin-example/pkg/param"

	sql "database/sql"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *IRepository) Count(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, invitation
func (_m *IRepository) Create(ctx context.Context, invitation model.Invitation) (sql.Result, error) {
	ret := _m.Called(ctx, invitation)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Invitation) (sql.Result, error)); ok {
		return rf(ctx, invitation)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Invitation) sql.Result); ok {
		r0 = rf(ctx, invitation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Invitation) error); ok {
		r1 = rf(ctx, invitation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAttendee provides a mock function with given fields: ctx, attendee
func (_m *IRepository) CreateAttendee(ctx context.Context, attendee model.Attendee) error {
	ret := _m.Called(ctx, attendee)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Attendee) error); ok {
		r0 = rf(ctx, attendee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, _a1
func (_m *IRepository) Get(ctx context.Context, _a1 param.Param) ([]model.Invitation, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []model.Invitation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) ([]model.Invitation, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) []model.Invitation); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Invitation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, param.Param) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IRepository) GetByID(ctx context.Context, id int64) (model.Invitation, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Invitation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.Invitation, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Invitation); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Invitation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMemberID provides a mock function with given fields: ctx, memberID
func (_m *IRepository) GetByMemberID(ctx context.Context, memberID int64) ([]model.InvitationDetail, error) {
	ret := _m.Called(ctx, memberID)

	var r0 []model.InvitationDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]model.InvitationDetail, error)); ok {
		return rf(ctx, memberID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []model.InvitationDetail); ok {
		r0 = rf(ctx, memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.InvitationDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, invitation, id
func (_m *IRepository) Update(ctx context.Context, invitation model.Invitation, id int64) (sql.Result, error) {
	ret := _m.Called(ctx, invitation, id)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Invitation, int64) (sql.Result, error)); ok {
		return rf(ctx, invitation, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Invitation, int64) sql.Result); ok {
		r0 = rf(ctx, invitation, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Invitation, int64) error); ok {
		r1 = rf(ctx, invitation, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
