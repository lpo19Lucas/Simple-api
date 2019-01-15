package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"apiTestLab/model/business"
	"apiTestLab/model/connection"
	"apiTestLab/model/entities"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	install()

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", getBooks).Methods("GET")

	routes.HandleFunc("/authors", getAuthors).Methods("GET")
	routes.HandleFunc("/authors", postAuthor).Methods("POST")
	routes.HandleFunc("/authors", deleteAuthor).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}

/**
 * Install the API
 * create the DB
 */
func install() {
	fmt.Println("## Create database")
	database := connection.MySqlConnect()

	selectTableAuthors := "SELECT * FROM Author"
	rows, _ := database.Query(selectTableAuthors)

	if rows == nil {
		fmt.Println("## Create table")
		createTableAuthors := "CREATE TABLE IF NOT EXISTS Author (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), email VARCHAR(255), pass VARCHAR(255))"
		statement, error := database.Prepare(createTableAuthors)

		if error != nil {
			fmt.Println("Erro", error)
		}

		statement.Exec()
		statement.Close()

		fmt.Println("## Insert data in table")
		statement, error = database.Prepare("INSERT INTO Author (name, email, pass) VALUES (?, ?, ?)")

		if error != nil {
			fmt.Println("ERRO: ", error)
		}

		statement.Exec("Lucas", "lucas01@teste.com", "1231")
		statement.Exec("Aurino", "author02@teste.com", "1232")
		statement.Exec("Ed", "author03@teste.com", "1233")
		statement.Close()
	}

}

func getAuthors(responseWriter http.ResponseWriter, responseRead *http.Request) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	authors := business.GetAllAuthor(&entities.Author{})
	json.NewEncoder(responseWriter).Encode(authors)
}

func postAuthor(responseWriter http.ResponseWriter, responseRead *http.Request) {
	var author entities.Author

	decoder := json.NewDecoder(responseRead.Body)
	err := decoder.Decode(&author)
	if err != nil {
		fmt.Println("Erro", err)
	}

	business.SaveAuthor(author)
}

func deleteAuthor(responseWriter http.ResponseWriter, responseRead *http.Request) {
	var author entities.Author

	decoder := json.NewDecoder(responseRead.Body)
	decoder.Decode(&author)

	business.DeleteAuthor(author)
}

func getBooks(responseWriter http.ResponseWriter, responseRead *http.Request) {
	var books = []entities.Book{
		entities.Book{Title: "Senhor dos Aneis", Value: 50.00},
		entities.Book{Title: "Uma Bela mulher", Value: 25.50},
	}
	json.NewEncoder(responseWriter).Encode(books)
}
