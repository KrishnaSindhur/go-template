package contract

import "net/http"

type Response struct {
	Items  interface{} `json:"items,omitempty"`
	Errors []Error     `json:"error,omitempty"`
}

type Error struct {
	HTTPStatus          int        `json:"-"`
	StatusCode          string     `json:"code"`
	StatusMessageFormat string     `json:"message"`
	Errors              []SubError `json:"errors,omitempty"`
}

func (err Error) Error() string {
	return err.StatusMessageFormat
}

var UnknownError = Error{HTTPStatus: http.StatusInternalServerError, StatusCode: "500", StatusMessageFormat: "UNKNOWN_ERROR"}

type SubError struct {
	StatusCode          string `json:"code"`
	StatusMessageFormat string `json:"message"`
	Property            string `json:"property,omitempty"`
}
