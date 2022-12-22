package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func Deneme(next http.Handler) http.Handler {
	return next
}
