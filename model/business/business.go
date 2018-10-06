package business

import (
	"apiTestLab/model/entities"
)

// GetAllAuthor Get all list the author
func GetAllAuthor(author *entities.Author) []entities.Author {
	author01 := *author
	author01.ID = 1
	author01.Name = "Author 01"
	author01.Email = "author01@teste.com"
	author01.Pass = "1231"

	author02 := *author
	author02.ID = 2
	author02.Name = "Author 02"
	author02.Email = "author02@teste.com"
	author02.Pass = "1232"

	author03 := *author
	author03.ID = 3
	author03.Name = "Author 03"
	author03.Email = "author03@teste.com"
	author03.Pass = "1233"

	var authorList = []entities.Author{
		author01,
		author02,
		author03,
	}
	return authorList
}
