package encode

import (
	"encoding/json"
	"io"

	"github.com/KrishnaSindhur/go-template/pkg/lib/contract"

	lgr "github.com/sirupsen/logrus"
)

func WriteJSON(w io.Writer, responseData interface{}) {
	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		lgr.Error("could not write response in json")
	}
}

func WriteResponseList(responseData interface{}, errors []contract.Error) contract.Response {
	response := contract.Response{Items: responseData, Errors: errors}
	return response
}
