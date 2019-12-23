package api

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"

)

//Http handler for responding to http/s requests.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user := r.Context().Value("user");
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		fmt.Fprintf(w, "%s :\t%#v\n", k, v)
	  }
	w.Write([]byte("{\"Route\":\"Test\"}"))
}
