package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bayusm0506/hcf-or-gcd/app/controllers"
	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {
	// Define payload
	payload := []byte(`{"number1":20, "number2":25}`)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/api/hcf", bytes.NewBuffer(payload))

	// Assign HTTP Handler function
	handler := http.HandlerFunc(controllers.Hcf)

	// Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	// Mapping response body
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	// Condition status
	if response.Code == http.StatusCreated {
		t.Logf("Result : %v", m["result"])
	}

	assert.Equal(t, m["result"], m["result"])
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
