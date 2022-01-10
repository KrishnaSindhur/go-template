package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KrishnaSindhur/go-template/pkg/handler"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.Health(w, r)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect http code")
	assert.JSONEq(t, `{"status":"ok"}`, w.Body.String(), "Incorrect response body")
}
