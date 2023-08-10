/*
go get github.com/gorilla/mux
go get github.com/urfave/negroni

curl -i 'http://localhost:5000/hello'
curl -i 'http://localhost:5000/hello?username=admin&password=password'
*/

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type badAuth struct {
	Username string
	Password string
}

func (b *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username != b.Username || password != b.Password {
		http.Error(w, "Unauthorized", 401)
		return
	}

	ctx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(ctx)
	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {

	// can also be retrieve directly from request, eg. username := r.URL.Query().Get("username")

	// %v can be used also even without the type assertion

	username := r.Context().Value("username").(string)
	/*
		The fmt.Fprintf function's format specifier %s can handle interface{} values
		as long as the underlying value is of a type that can be meaningfully represented as a string
		Better approach is to assert the type and check for errors:
		username, ok := r.Context().Value("username").(string)
	*/
	fmt.Fprintf(w, "Hello, %s", username)

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")
	n := negroni.Classic()
	n.Use(&badAuth{
		Username: "admin",
		Password: "password",
	})
	n.UseHandler(r)
	http.ListenAndServe(":5000", n)
}
