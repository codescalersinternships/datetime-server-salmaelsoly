package routes

import (
	"encoding/json"
	"log"
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

func assertStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDataTimeGinHandler(t *testing.T) {
	tests := []struct {
		method     string
		endpoint   string
		statusCode int
		expected   string
	}{
		{http.MethodGet, "/datetime", http.StatusOK, time.Now().Format(time.UnixDate)},
		{http.MethodPost, "/datetime", http.StatusNotFound, ""},
	}
	for _, test := range tests {
		t.Run("Server Test", func(t *testing.T) {
			router := SetUpRouter()
			req := httptest.NewRequest(test.method, test.endpoint, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assertStatusCode(t, w.Code, test.statusCode)
			if test.method == http.MethodGet {
				println(w.Body)
				var jsonObject map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &jsonObject)
				if err != nil {
					log.Fatal(err)
				}
				print(jsonObject["message"])
				assertEqual(t, jsonObject["message"], test.expected)
			}
		})

	}
}
