package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(addr string, m *mux.Router) {
	s := &http.Server{
		Addr:    addr,
		Handler: m,
	}
	fmt.Printf("Listening %s\n", addr)
	log.Fatal(s.ListenAndServe())
}
