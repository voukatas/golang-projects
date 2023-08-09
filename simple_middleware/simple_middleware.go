package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("start")

	l.Inner.ServeHTTP(w, r)

	log.Println("end")

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")

}

func main() {

	f := http.HandlerFunc(hello)
	l := logger{Inner: f}

	http.ListenAndServe(":5000", &l)

}
