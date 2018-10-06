package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"apiTestLab/model/business"
	"apiTestLab/model/entities"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const PATH_FILE_DB = "./db/database.db"

func main() {
	_, err := os.Stat(PATH_FILE_DB)
	if err != nil && os.IsNotExist(err) {
		install()
	}

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", getBooks).Methods("GET")

	routes.HandleFunc("/authors", getAuthors).Methods("GET")
	routes.HandleFunc("/authors", postAuthors).Methods("POST")

	var port = ":4000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}

/**
 * Install the API
 * create the DB
 */
func install() {
	fmt.Println("Create database...")
	database, _ := sql.Open("sqlite3", PATH_FILE_DB)

	createTableAuthors := "CREATE TABLE IF NOT EXISTS Authors (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT, pass, TEXT)"
	statement, _ := database.Prepare(createTableAuthors)
	statement.Exec()
}

func getAuthors(responseWriter http.ResponseWriter, responseRead *http.Request) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	authors := business.GetAllAuthor(&entities.Author{})
	json.NewEncoder(responseWriter).Encode(authors)
}

func postAuthors(responseWriter http.ResponseWriter, responseRead *http.Request) {
	fmt.Println("Call the POST")
}

func getBooks(responseWriter http.ResponseWriter, responseRead *http.Request) {
	var books = []entities.Book{
		entities.Book{Title: "Senhor dos Aneis", Value: 50.00},
		entities.Book{Title: "Uma Bela mulher", Value: 25.50},
	}
	json.NewEncoder(responseWriter).Encode(books)
}
