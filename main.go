package main

import (
	"fmt"
	"log"
	"net/http"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

type ConnectionURL struct {
	User     string
	Password string
	Host     string
	Database string
	Options  map[string]string
}

var settings = mysql.ConnectionURL{
	Database: `Go_Book_API`, // Database name
	Host:     `localhost`,   // Server IP or name
	User:     `root`,        // Username // Will be environment variable
	Password: `Password`,    // Password  //Will be environment variable
}

type BookJson struct {
	ID          int    `db:"BookID"`
	Title       string `db:"Title"`
	Description string `db:"Desciption"`
	Genre       string `db:"Genre"`
}

type Books struct {
	Books []BookJson
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("Endpoint Hit")
}

// func booksHandler(w http.ResponseWriter, r *http.Request) {
// 	jsonFile, err := os.Open("BookData.json")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Successfully opened json")
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)

// 	var books Books

// 	json.Unmarshal(byteValue, &books)

// 	json.NewEncoder(w).Encode(books)
// }

func DatabaseSetUp(w http.ResponseWriter, r *http.Request) {
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

func handleRequests() {
	http.HandleFunc("/books", DatabaseSetUp)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()

}
