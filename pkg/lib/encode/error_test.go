package encode_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KrishnaSindhur/go-template/pkg/lib/contract"
	"github.com/KrishnaSindhur/go-template/pkg/lib/encode"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func Test_WriteErrorResponseWithgivenError(t *testing.T) {
	w := httptest.NewRecorder()
	err := xerrors.Errorf("malformed request: %w", contract.UnknownError)

	encode.WriteErrorResponse(w, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error":[{"code":"500", "message":"UNKNOWN_ERROR"}]}`, w.Body.String(), "incorrect body")
}

func Test_WriteErrorResponseWithUnExpectedError(t *testing.T) {
	w := httptest.NewRecorder()
	err := xerrors.Errorf("malformed request")

	encode.WriteErrorResponse(w, err)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Empty(t, w.Body.String())

}
