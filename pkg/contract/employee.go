package contract

import "strings"

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Location  string `json:"location"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type CreateEmployee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Location  string `json:"location"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (req CreateEmployee) Validate() error {
	valErr := ValidationError

	if len(strings.TrimSpace(req.FirstName)) <= 0 {
		valErr.Errors = append(valErr.Errors, FirstNameMissing)
	}
	if len(strings.TrimSpace(req.LastName)) <= 0 {
		valErr.Errors = append(valErr.Errors, LastNameMissing)
	}
	if len(strings.TrimSpace(req.Location)) <= 0 {
		valErr.Errors = append(valErr.Errors, LocationMissing)
	}

	if len(valErr.Errors) > 0 {
		return valErr
	}
	return nil
}

type UpdateEmployee struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Location  string `json:"location,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (req UpdateEmployee) Validate() error {
	valErr := ValidationError

	if len(strings.TrimSpace(req.FirstName)) <= 0 && len(strings.TrimSpace(req.LastName)) <= 0 && len(strings.TrimSpace(req.Location)) <= 0 &&
		len(strings.TrimSpace(req.Email)) <= 0 && len(strings.TrimSpace(req.Phone)) <= 0 {
		valErr.Errors = append(valErr.Errors, UpdateValueMissing)
	}

	if len(valErr.Errors) > 0 {
		return valErr
	}
	return nil
}
