package main

import (
	"testing"
	"os"
	"net/http"
	"net/http/httptest"
)

var a App


func TestMain(m *testing.M) {
	a = App{}
	a.initialize()

	code := m.Run()

	os.Exit(code)
}

func TestIllegalRequest(t *testing.T) {

	req, _ := http.NewRequest("POST", "/search", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusForbidden, response.Code)

	if body := response.Body.String(); body != "Allow: GET" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestAllowedRequest(t *testing.T){
	req, _ := http.NewRequest("GET", "/search?q=hello", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); len(body) < 0 {
		t.Errorf("Expected a populated array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	 a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}