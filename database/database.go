package restapi

import (
	"database/sql"
	"log"
	// "strings"

	_ "github.com/mattn/go-sqlite3"
	"librarian/models"
)

var DB *sql.DB

func InitDB() (err error) {
	DB, err = sql.Open("sqlite3", "database/librarian.db")
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func InsertBook(book *models.Book) error {
	statement, err := DB.Prepare(`INSERT INTO book VALUES (?,?,?,?,?)`)
	if err != nil {
		log.Printf("Could not prepare statement: %v\n", err)
		return err
	}
	_, err = statement.Exec(
		&book.Isbn,
		book.Title,
		book.Author,
		book.Published,
		book.Genre,
	)
	if err != nil {
		log.Printf("Could not insert book: %v\n", err)
	}
	return err
}

func GetBooks(books *[]*models.Book) error {
	log.Println(DB)
	statement, _ := DB.Prepare(`SELECT * FROM book`)
	row, err := statement.Query()
	if err != nil {
		log.Fatalf("Failed to get %s", err)
		return err
	}
	for row.Next() {
		b := models.Book{}
		row.Scan(&b.Isbn, &b.Title, &b.Author, &b.Published, &b.Genre)
		log.Println("GOT ", b)
		*books = append(*books, &b)
	}

	return err
}
