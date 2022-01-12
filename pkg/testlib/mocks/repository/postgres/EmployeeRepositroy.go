// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	contract "github.com/KrishnaSindhur/go-template/pkg/contract"
	mock "github.com/stretchr/testify/mock"
)

// EmployeeRepositroy is an autogenerated mock type for the EmployeeRepositroy type
type EmployeeRepositroy struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, emp
func (_m *EmployeeRepositroy) Create(ctx context.Context, emp contract.CreateEmployee) (contract.Employee, error) {
	ret := _m.Called(ctx, emp)

	var r0 contract.Employee
	if rf, ok := ret.Get(0).(func(context.Context, contract.CreateEmployee) contract.Employee); ok {
		r0 = rf(ctx, emp)
	} else {
		r0 = ret.Get(0).(contract.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, contract.CreateEmployee) error); ok {
		r1 = rf(ctx, emp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, empID
func (_m *EmployeeRepositroy) Delete(ctx context.Context, empID string) error {
	ret := _m.Called(ctx, empID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, empID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, empID
func (_m *EmployeeRepositroy) Get(ctx context.Context, empID string) (contract.Employee, error) {
	ret := _m.Called(ctx, empID)

	var r0 contract.Employee
	if rf, ok := ret.Get(0).(func(context.Context, string) contract.Employee); ok {
		r0 = rf(ctx, empID)
	} else {
		r0 = ret.Get(0).(contract.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, empID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, location
func (_m *EmployeeRepositroy) List(ctx context.Context, location string) ([]contract.Employee, error) {
	ret := _m.Called(ctx, location)

	var r0 []contract.Employee
	if rf, ok := ret.Get(0).(func(context.Context, string) []contract.Employee); ok {
		r0 = rf(ctx, location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contract.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, empID, updateEmp
func (_m *EmployeeRepositroy) Patch(ctx context.Context, empID string, updateEmp contract.UpdateEmployee) (contract.Employee, error) {
	ret := _m.Called(ctx, empID, updateEmp)

	var r0 contract.Employee
	if rf, ok := ret.Get(0).(func(context.Context, string, contract.UpdateEmployee) contract.Employee); ok {
		r0 = rf(ctx, empID, updateEmp)
	} else {
		r0 = ret.Get(0).(contract.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, contract.UpdateEmployee) error); ok {
		r1 = rf(ctx, empID, updateEmp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
