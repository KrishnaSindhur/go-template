package contract

import (
	"testing"

	"github.com/KrishnaSindhur/go-template/pkg/lib/contract"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeCreateRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateEmployee
		wantErr bool
		subErr  []contract.SubError
	}{
		{name: "Empty request", req: CreateEmployee{}, wantErr: true,
			subErr: []contract.SubError{FirstNameMissing}},
		{name: "Invalid request", req: CreateEmployee{FirstName: "abc"}, wantErr: true,
			subErr: []contract.SubError{LastNameMissing}},
		{name: "invalid request", req: CreateEmployee{FirstName: "abc", LastName: "xyz"}, wantErr: true,
			subErr: []contract.SubError{LocationMissing}},
		{name: "valid request", req: CreateEmployee{FirstName: "abc", LastName: "xyz", Location: "eng"}, wantErr: false,
			subErr: []contract.SubError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.wantErr {
				assert.Error(t, err, "Expected error")
				assert.IsType(t, err, contract.Error{})
				assert.Subset(t, err.(contract.Error).Errors, tt.subErr)
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
		})
	}
}

func TestEmployeeUpdateRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     UpdateEmployee
		wantErr bool
		subErr  []contract.SubError
	}{
		{name: "Empty request", req: UpdateEmployee{}, wantErr: true,
			subErr: []contract.SubError{UpdateValueMissing}},
		{name: "Valid request", req: UpdateEmployee{FirstName: "abc"}, wantErr: false,
			subErr: []contract.SubError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.wantErr {
				assert.Error(t, err, "Expected error")
				assert.IsType(t, err, contract.Error{})
				assert.Subset(t, err.(contract.Error).Errors, tt.subErr)
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
		})
	}
}
