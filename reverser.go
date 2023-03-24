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

func (r *reverser) Reverse(c context.Context, s string) (string, error) {

	log := r.Logger()
	log.Info("reverti uma string!")
	runes := []rune(s)
	n := len(runes)
	root := weaver.Init(context.Background())
	dbClient, err := weaver.Get[DbClient](root)
	if err != nil {
		log.Error("erro ao obter cliente!", err)
	}

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	dbClient.Insert(c, "Runas", "Invertidas", string(runes))
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
