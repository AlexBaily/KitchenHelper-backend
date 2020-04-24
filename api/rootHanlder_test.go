package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Route":"Test"}`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned incorrect body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
