package server

import (
	"log"
	"net/http"
	"sync"
)

type Service struct {
	connectionString string
	items            map[string]Item
	sync.RWMutex
}

func NewService(connectionString string, items map[string]Item) *Service {
	return &Service{
		connectionString: connectionString,
		items:            items,
	}
}

func (s *Service) ListenAndServe() error {
	r := mux.NewRouter()

	r.HandleFunc("/item", logs(auth(s.PostItem)).Methods("POST"))
	r.HandleFunc("/item", logs(auth(s.GetItems)).Methods("GET"))
	r.HandleFunc("/item/{name}", logs(auth(s.GetItem)).Methods("GET"))
	r.HandleFunc("/item/{name}", logs(auth(s.PutItem)).Methods("PUT"))
	r.HandleFunc("/item/{name}", logs(auth(s.DeleteItem)).Methods("DELETE"))

	log.Printf("Starting server on %s", s.connectionString)
	err := http.ListenAndServe(s.connectionString, r)
	if err != nil {
		return err
	}
	return nil
}

func logs(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		log.Printf("%s %s", method, path)
		handlerFunc(w, r)
		return
	}
}
func auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Please supply Authorization token", http.StatusUnauthorized)
			return
		}

		handlerFunc(w, r)
		return
	}
}
