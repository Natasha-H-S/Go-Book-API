package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Natasha-H-S/Go-Book-API/internal/repository"
	"github.com/Natasha-H-S/Go-Book-API/pkg/config"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("Endpoint Hit")
}

func handleRequests(cfg *config.Config) {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		repository.GetBooks(w, r, cfg)
	})
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	cfg := config.NewConfig()

	handleRequests(cfg)

}
