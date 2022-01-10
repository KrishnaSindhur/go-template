package contract

import (
	"net/http"

	"github.com/KrishnaSindhur/go-template/pkg/lib/contract"
)

var (
	ValidationError    = contract.Error{StatusCode: "50000", HTTPStatus: http.StatusBadRequest, StatusMessageFormat: "validationError"}
	FirstNameMissing   = contract.SubError{StatusCode: "50001", StatusMessageFormat: "firstname missing", Property: "firstName"}
	LastNameMissing    = contract.SubError{StatusCode: "50002", StatusMessageFormat: "lastname missing", Property: "lastName"}
	LocationMissing    = contract.SubError{StatusCode: "50003", StatusMessageFormat: "location missing", Property: "location"}
	EmailMissing       = contract.SubError{StatusCode: "50004", StatusMessageFormat: "email is missing", Property: "email"}
	PhoneMissing       = contract.SubError{StatusCode: "50005", StatusMessageFormat: "phone missing", Property: "phone"}
	UpdateValueMissing = contract.SubError{StatusCode: "50007", StatusMessageFormat: "update value missing", Property: "lastName,firstName,location,email,phone"}
)
