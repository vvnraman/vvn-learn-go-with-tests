package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server_v1(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

func Server_v2(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}

func Server_v3(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

type Store_v2 interface {
	Fetch(ctx context.Context) (string, error)
}

func Server_v4(store Store_v2) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

func main() {
	fmt.Println("vim-go")
}
