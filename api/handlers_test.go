package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func assertEqual(t *testing.T, got, want any) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertStatusCode(t *testing.T, got,want int){
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestDataTimeHandler(t *testing.T) {
	tests := []struct {
		method     string
		endpoint   string
		data       io.Reader
		statusCode int
		expected   string
	}{
		{http.MethodGet, "/datetime", nil, http.StatusOK, time.Now().Format(time.UnixDate)},
		{http.MethodPost, "/datetime", nil, http.StatusBadRequest, ""},
	}
	for _, test := range tests {
		t.Run("Server Test", func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.endpoint, test.data)
			w := httptest.NewRecorder( )
			DateTime(w, req)
			assertStatusCode(t,w.Result().StatusCode,test.statusCode)
			if test.method == http.MethodGet{
				actual,_:=io.ReadAll(w.Result().Body)
				assertEqual(t,string(actual),test.expected)
			}
		})

	}
}
