package handler_test

import (
	"net/http"

	"github.com/gorilla/mux"
)

//create mocks: mockery --dir=pkg --output=pkg/testlib/mocks --all --keeptree
func muxRouter(handler http.Handler, path string, method string) *mux.Router {
	if method == "" {
		method = http.MethodGet
	}

	r := mux.NewRouter()
	r.Handle(path, handler).Methods(method)
	return r
}
