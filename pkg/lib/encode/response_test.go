package encode_test

import (
	"bytes"
	"testing"

	"github.com/KrishnaSindhur/go-template/pkg/lib/contract"
	"github.com/KrishnaSindhur/go-template/pkg/lib/encode"
	"github.com/stretchr/testify/assert"
)

func TestWriteErrorResponseJSONByCreatingResponseList(t *testing.T) {
	var w bytes.Buffer
	response := map[string]string{"A": "one", "B": "two"}

	encode.WriteJSON(&w, encode.WriteResponseList(response, nil))

	assert.JSONEq(t, `{"items":{"A": "one", "B":"two"}}`, w.String(), "incorrect response written")
}

func TestWriteResponseJSON(t *testing.T) {
	var w bytes.Buffer
	response := map[string]string{"A": "one", "B": "two"}

	encode.WriteJSON(&w, response)

	assert.JSONEq(t, `{"A": "one", "B":"two"}`, w.String(), "incorrect response written")
}

func TestWriteResponseJSONWuthErrors(t *testing.T) {
	var w bytes.Buffer

	encode.WriteJSON(&w, encode.WriteResponseList(nil, []contract.Error{contract.UnknownError}))

	assert.JSONEq(t, `{"error":[{"code": "500", "message":"UNKNOWN_ERROR"}}`, w.String(), "incorrect response written")
}

func TestWriteResponseJSONWhenJSONEncodeErrors(t *testing.T) {
	var w bytes.Buffer
	response := func() string { return "response" }

	encode.WriteJSON(&w, response)

	assert.Empty(t, w.String())
}
