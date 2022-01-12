package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KrishnaSindhur/go-template/pkg/contract"
	"github.com/KrishnaSindhur/go-template/pkg/handler"
	mocks "github.com/go-template/pkg/testlib/mocks/repository/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListAllEmployeeForLocation(t *testing.T) {
	loc := "ind"
	repomock := new(mocks.EmployeeRepositroy)
	r := httptest.NewRequest(http.MethodGet, "/employees?location=ind", nil).WithContext(context.Background())
	w := httptest.NewRecorder()
	var employees []contract.Employee
	employee := contract.Employee{
		ID:        "123-abc",
		FirstName: "abc",
		LastName:  "xyz",
		Location:  "ind",
		Email:     "abc@gmail.com",
		Phone:     "11233334",
	}
	employees = append(employees, employee)
	repomock.On("List", mock.AnythingOfType("*context.valueCtx"), loc).Return(employees, nil)

	muxRouter(handler.List(repomock), "employees", http.MethodGet).ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect HTTP status code")
	// assert.JSONEq()
	mock.AssertExpectationsForObjects(t, repomock)
}
