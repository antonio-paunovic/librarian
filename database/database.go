package restapi

import (
	"database/sql"
	"fmt"
	"librarian/models"
	"log"
	"strings"

	"github.com/go-openapi/strfmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbname string) (err error) {
	DB, err = sql.Open("sqlite3", dbname)
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

func DeleteBook(isbn string) error {
	_, err := DB.Exec(fmt.Sprintf(`DELETE FROM book WHERE isbn="%s"`, isbn))
	if err != nil {
		log.Printf("Could not delete book: %v\n", err)
	}
	return err
}

func UpdateBook(book *models.Book, isbn string) error {
	var stmtStr string = `UPDATE book SET`
	if book.Title != "" {
		stmtStr += fmt.Sprintf(` title="%s",`, book.Title)
	}
	if book.Author != "" {
		stmtStr += fmt.Sprintf(` author="%s",`, book.Author)
	}
	if !book.Published.Equal(strfmt.Date{}) {
		stmtStr += fmt.Sprintf(` published="%s",`, book.Published)
	}
	if book.Genre != "" {
		stmtStr += fmt.Sprintf(` genre="%s",`, book.Genre)
	}
	stmtStr = strings.TrimRight(stmtStr, ", ") + fmt.Sprintf(` WHERE isbn="%s"`, isbn)

	_, err := DB.Exec(stmtStr)
	if err != nil {
		log.Printf("Could not update book: %v\n", err)
		return err
	}
	return err
}

func GetBooks(books *[]*models.Book, filter map[string]string) error {
	var stmtStr string
	if len(filter) == 0 {
		stmtStr = "SELECT * FROM book"
	} else {
		filterStmt := []string{}
		for k, v := range filter {
			if v != "" {
				filterStmt = append(filterStmt, fmt.Sprintf(`%s="%s"`, k, v))
			}
		}
		filterStmtStr := strings.Join(filterStmt, " AND ")
		stmtStr = "SELECT * FROM book WHERE " + filterStmtStr
	}

	rows, err := DB.Query(stmtStr)
	if err != nil {
		log.Fatalf("Failed to get %s", err)
		return nil
	}

	for rows.Next() {
		b := models.Book{}
		rows.Scan(&b.Isbn, &b.Title, &b.Author, &b.Published, &b.Genre)
		*books = append(*books, &b)
	}

	return err
}

func GetCollections(collections *[]*models.Collection) error {
	var stmtStr string
	stmtStr = "SELECT * FROM collection"
	rows, err := DB.Query(stmtStr)
	if err != nil {
		log.Fatalf("Failed to get %s", err)
		return nil
	}

	cs := map[string][]string{}

	for rows.Next() {
		var name string
		var isbn string
		rows.Scan(&name, &isbn)
		cs[name] = append(cs[name], isbn)
	}

	for name, isbns := range cs {
		c := models.Collection{isbns, &name}
		*collections = append(*collections, &c)
	}

	return err
}

func InsertCollection(collection *models.Collection) error {
	statement, err := DB.Prepare(`INSERT INTO collection VALUES (?,?)`)
	if err != nil {
		log.Printf("Could not prepare statement: %v\n", err)
		return err
	}

	for _, isbn := range collection.Isbn {
		_, err = statement.Exec(
			&collection.Name,
			isbn,
		)
		if err != nil {
			log.Printf("Could not insert collection: %v\n", err)
		}
	}
	return err
}

func DeleteCollection(name string) error {
	_, err := DB.Exec(fmt.Sprintf(`DELETE FROM collection WHERE name="%s"`, name))
	if err != nil {
		log.Printf("Could not delete collection: %v\n", err)
	}
	return err
}

func UpdateCollection(collection *models.Collection, name string) error {
	statement, err := DB.Prepare(`UPDATE collection SET name="?" WHERE name="?" `)
	if err != nil {
		log.Printf("Could not prepare statement: %v\n", err)
		return err
	}

	for _, isbn := range collection.Isbn {
		_, err = statement.Exec(
			isbn,
			&collection.Name,
		)
		if err != nil {
			log.Printf("Could not update collection: %v\n", err)
		}
	}

	return err
}
