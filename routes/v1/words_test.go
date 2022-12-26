package routes_test

import (
	"backend/routes/v1"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestIsValidRequestSuccess(t *testing.T) {
	vars := map[string]string{
		"word": "which",
	}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(vars)
	if err != nil {
		log.Panicf("couldn't convert vars to json %s", err.Error())
	}

	r, _ := http.NewRequest(http.MethodPost, "/v1/isValid", body)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, vars)
	routes.IsValidWord(w, r)

	expected := routes.IsValidResponse{
		Word:    "which",
		IsValid: true,
	}

	expectedJSON, _ := json.Marshal(expected)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedJSON, w.Body.Bytes())
}

func TestIsValidRequest400(t *testing.T) {
	vars := map[string]string{
		"words": "which",
	}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(vars)
	if err != nil {
		log.Panicf("couldn't convert vars to json %s", err.Error())
	}

	r, _ := http.NewRequest(http.MethodPost, "/v1/isValid", body)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, vars)
	routes.IsValidWord(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
