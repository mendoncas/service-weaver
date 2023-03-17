package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

// Reverser component.
type Reverser interface {
	Reverse(context.Context, string) (string, error)
	ReverseController(context.Context) error
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser]
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}

func (rev *reverser) ReverseController(context.Context) error {
	logger := rev.Logger()
	logger.Info("componente reverser inicializado")
	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		defer logger.Info("acabei de reverter uma string")
		reversed, err := rev.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			logger.Error("erro", nil)
		}
		fmt.Fprintf(w, reversed)
	})
	return nil
}
