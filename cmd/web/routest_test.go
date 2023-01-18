package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/hakansenn/goWeb/internals/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing test passed
	default:
		t.Error(fmt.Sprintf("Type is not chi.mux but is %T", v))
	}
}
