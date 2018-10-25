package business

import (
	"apiTestLab/model/connection"
	"apiTestLab/model/entities"
)

// GetAllAuthor Get all list the author
func GetAllAuthor(author *entities.Author) []entities.Author {
	var authorList []entities.Author

	database := connection.SqliteConnect()

	rows, _ := database.Query("SELECT id, name, email, pass FROM Author")
	for rows.Next() {
		author := *author

		rows.Scan(&author.ID, &author.Name, &author.Email, &author.Pass)
		authorList = append(authorList, author)
	}

	return authorList
}

// SaveAuthor Save a author in DB
func SaveAuthor(author entities.Author) int64 {
	database := connection.SqliteConnect()

	statement, _ := database.Prepare("INSERT INTO Author (name, email, pass) VALUES (?, ?, ?)")
	result, _ := statement.Exec(author.Name, author.Email, author.Pass)
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected
}

// DeleteAuthor Delete data of the author in DB
func DeleteAuthor(author entities.Author) int64 {
	database := connection.SqliteConnect()

	statement, _ := database.Prepare("DELETE FROM Author WHERE id = ?")
	result, _ := statement.Exec(author.ID)
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected
}
