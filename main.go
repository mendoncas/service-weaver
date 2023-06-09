package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	// Get a network listener on address "localhost:12345".
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := weaver.Get[DbClient](root)
	if err != nil {
		log.Fatal(err)
	}

	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("hello listener available on %v\n", lis)

	// serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "hello! your reversed name is %s!\n", reversed)
	})

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, reversed)
	})

	http.HandleFunc("/testConnection", func(w http.ResponseWriter, r *http.Request) {
		dbClient.StartConnection(r.Context())
	})

	// reverser.ReverseController(context.TODO())

	http.Serve(lis, nil)
}
