package api

import (
	"net/http"
)

//Http handler for responding to http/s requests.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("{\"Route\":\"Test\"}"))
}
