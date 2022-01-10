package encode

import (
	"net/http"

	"golang.org/x/xerrors"

	"github.com/KrishnaSindhur/go-template/pkg/lib/contract"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	var errResponse contract.Error
	w.Header().Set("Content-Type", "application/json")
	if !xerrors.As(err, &errResponse) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(errResponse.HTTPStatus)
	WriteJSON(w, WriteResponseList(nil, []contract.Error{errResponse}))
}
