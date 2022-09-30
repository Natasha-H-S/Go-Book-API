package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type BookJson struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
}

type Books struct {
	Books []BookJson `json:"books"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("Endpoint Hit")
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("BookData.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var books Books

	json.Unmarshal(byteValue, &books)

	json.NewEncoder(w).Encode(books)
}

func handleRequests() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()

}
