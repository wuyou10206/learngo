package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}
func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}
func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{noError, 200, "no error"},
}

func checkHttp(response *http.Response, code int, message string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != code || body != message {
		t.Errorf("expect (%d,%s) ;got (%d,%s)", code, message, response.StatusCode, body)
	}
}
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/list/fib2.txt", nil)
		f(response, request)
		//b, _ := ioutil.ReadAll(response.Body)
		//body := strings.Trim(string(b), "\n")
		//if response.Code != tt.code || body != tt.message {
		//	t.Errorf("expect (%d,%s) ;got (%d,%s)", tt.code, tt.message, response.Code, body)
		//}
		checkHttp(response.Result(), tt.code, tt.message, t)
	}
}
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		sever := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(sever.URL)
		checkHttp(response, tt.code, tt.message, t)
	}

}
