// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import todo_groups "github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"

// TodoGroupsServiceServer is an autogenerated mock type for the TodoGroupsServiceServer type
type TodoGroupsServiceServer struct {
	mock.Mock
}

// CreateTodoGroup provides a mock function with given fields: _a0, _a1
func (_m *TodoGroupsServiceServer) CreateTodoGroup(_a0 context.Context, _a1 *todo_groups.CreateTodoGroupRequest) (*todo_groups.CreateTodoGroupResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *todo_groups.CreateTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *todo_groups.CreateTodoGroupRequest) *todo_groups.CreateTodoGroupResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo_groups.CreateTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *todo_groups.CreateTodoGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodoGroup provides a mock function with given fields: _a0, _a1
func (_m *TodoGroupsServiceServer) DeleteTodoGroup(_a0 context.Context, _a1 *todo_groups.DeleteTodoGroupRequest) (*todo_groups.DeleteTodoGroupResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *todo_groups.DeleteTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *todo_groups.DeleteTodoGroupRequest) *todo_groups.DeleteTodoGroupResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo_groups.DeleteTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *todo_groups.DeleteTodoGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoGroup provides a mock function with given fields: _a0, _a1
func (_m *TodoGroupsServiceServer) GetTodoGroup(_a0 context.Context, _a1 *todo_groups.GetTodoGroupRequest) (*todo_groups.GetTodoGroupResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *todo_groups.GetTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *todo_groups.GetTodoGroupRequest) *todo_groups.GetTodoGroupResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo_groups.GetTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *todo_groups.GetTodoGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTodoGroups provides a mock function with given fields: _a0, _a1
func (_m *TodoGroupsServiceServer) ListTodoGroups(_a0 context.Context, _a1 *todo_groups.ListTodoGroupsRequest) (*todo_groups.ListTodoGroupsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *todo_groups.ListTodoGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *todo_groups.ListTodoGroupsRequest) *todo_groups.ListTodoGroupsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo_groups.ListTodoGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *todo_groups.ListTodoGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodoGroup provides a mock function with given fields: _a0, _a1
func (_m *TodoGroupsServiceServer) UpdateTodoGroup(_a0 context.Context, _a1 *todo_groups.UpdateTodoGroupRequest) (*todo_groups.UpdateTodoGroupResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *todo_groups.UpdateTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *todo_groups.UpdateTodoGroupRequest) *todo_groups.UpdateTodoGroupResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo_groups.UpdateTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *todo_groups.UpdateTodoGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}