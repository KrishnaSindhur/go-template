package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KrishnaSindhur/go-template/pkg/contract"
	"github.com/KrishnaSindhur/go-template/pkg/lib/encode"
	"github.com/KrishnaSindhur/go-template/pkg/repository/postgres"
	"github.com/gorilla/mux"
	lgr "github.com/sirupsen/logrus"
)

const (
	empIDParamKey    = "id"
	locationParamKey = "location"
)

func CreateEmployee(er postgres.EmployeeRepositroy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateEmployee, err := readRequestBodyToCreateEmployee(r)
		if err != nil {
			encode.WriteErrorResponse(w, err)
			return
		}

		employee, err := er.Create(r.Context(), CreateEmployee)

		if err != nil {
			lgr.Errorf("could not create employee: %v", err.Error())
			encode.WriteErrorResponse(w, err)
			return
		}
		lgr.Info("emplyee successfully created")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		encode.WriteJSON(w, employee)
	}
}

func List(er postgres.EmployeeRepositroy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loc string
		if params, paramExit := r.URL.Query()[locationParamKey]; paramExit {
			loc = params[0]
		}
		employeeList, err := er.List(r.Context(), loc)

		if err != nil {
			lgr.Errorf("failed to get all employee details:%v", err.Error())
			encode.WriteErrorResponse(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encode.WriteJSON(w, encode.WriteResponseList(employeeList, nil))
	}
}

func Get(er postgres.EmployeeRepositroy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		empID := mux.Vars(r)[empIDParamKey]
		employee, err := er.Get(r.Context(), empID)

		if err != nil {
			lgr.Errorf("failed to get emplyee data: %v", err.Error())
			encode.WriteErrorResponse(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		encode.WriteJSON(w, employee)
	}
}

func readRequestBodyToCreateEmployee(r *http.Request) (contract.CreateEmployee, error) {
	var emp contract.CreateEmployee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		lgr.Errorf("malformed request body: %v", err.Error())
		return contract.CreateEmployee{}, contract.MalformedRequestBody
	}

	if err := emp.Validate(); err != nil {
		lgr.Errorf("validation failed: %s", err.Error())
		return contract.CreateEmployee{}, err
	}
	return emp, nil
}
