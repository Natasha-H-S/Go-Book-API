package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Natasha-H-S/Go-Book-API/pkg/config"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

type ConnectionURL struct {
	User     string
	Password string
	Host     string
	Database string
}

type BookJson struct {
	ID          int    `db:"ID"`
	Title       string `db:"Title"`
	Description string `db:"Description"`
	Genre       string `db:"Genre"`
}

type Books struct {
	Books []BookJson
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("Endpoint Hit")
}

func getBooks(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	settings := mysql.ConnectionURL{
		User:     cfg.User,
		Password: cfg.Password,
		Host:     cfg.Host,
		Database: cfg.Database,
	}

	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()

	bookCollection := sess.Collection("Books")

	var res db.Result
	res = bookCollection.Find()

	var books []BookJson

	err = res.All(&books)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	for _, book := range books {
		fmt.Fprintf(w, "%d %s %s %s\n",
			book.ID,
			book.Title,
			book.Description,
			book.Genre)
	}
}

func handleRequests(cfg *config.Config) {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		getBooks(w, r, cfg)
	})
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	cfg := config.NewConfig()

	handleRequests(cfg)

}
