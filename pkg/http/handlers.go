package routes

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func DateTime(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "%v", time.Now().Format(time.UnixDate))
	slog.Info(req.RemoteAddr + " Requested date time")

}
