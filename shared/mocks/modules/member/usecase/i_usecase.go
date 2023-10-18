// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/rzfhlv/gin-example/internal/modules/member/model"
	mock "github.com/stretchr/testify/mock"

	param "github.com/rzfhlv/gin-example/pkg/param"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, memberPayload
func (_m *IUsecase) Create(ctx context.Context, memberPayload model.Member) (model.Member, error) {
	ret := _m.Called(ctx, memberPayload)

	var r0 model.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Member) (model.Member, error)); ok {
		return rf(ctx, memberPayload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Member) model.Member); ok {
		r0 = rf(ctx, memberPayload)
	} else {
		r0 = ret.Get(0).(model.Member)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Member) error); ok {
		r1 = rf(ctx, memberPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, _a1
func (_m *IUsecase) Get(ctx context.Context, _a1 param.Param) ([]model.Member, int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []model.Member
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) ([]model.Member, int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) []model.Member); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, param.Param) int64); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, param.Param) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IUsecase) GetByID(ctx context.Context, id int64) (model.Member, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.Member, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Member); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Member)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUsecase creates a new instance of IUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUsecase {
	mock := &IUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
