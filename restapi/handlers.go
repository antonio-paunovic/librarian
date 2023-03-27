// This packages implements handler functions used in configure_librarian.go generated file.
package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	database "librarian/database"
	"librarian/models"
	"librarian/restapi/operations/books"
	"librarian/restapi/operations/collections"
	"strings"
)

// Return map of what value will be used for which parameters of the filter.
// Example:
// filter["author"] = "J.R.R. Tolkien" // should keep only the books written by this author.
//
// filterStr has a form of <filter-key> eq <filter-value> and <filter-key> eq <filter-value>
func parseFilter(filterStr *string) map[string]string {
	filter := map[string]string{}
	if filterStr == nil {
		return filter
	}

	for _, term := range strings.Split(*filterStr, "and") {
		termSlice := strings.Split(term, "eq")
		k, v := strings.Trim(termSlice[0], `" `), strings.Trim(termSlice[1], `" `)
		filter[k] = v
	}

	return filter
}

// Get books and filter the results. If no filter is given, all books are returned.
// Return `200` on success.
// Return `500` on database error.
func HandleGetBooks(params books.GetBooksParams) middleware.Responder {
	bs := []*models.Book{}
	if err := database.GetBooks(&bs, parseFilter(params.Filter)); err != nil {
		return books.NewGetBooksDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})

	}
	return books.NewGetBooksOK().WithPayload(bs)
}

// Post the book by inserting it into database.
// Return `200` on success.
// Return `500` on database error.
func HandlePostBooks(params books.PostBooksParams) middleware.Responder {
	if err := database.InsertBook(params.Body); err != nil {
		return books.NewPostBooksDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return books.NewPostBooksOK().WithPayload(params.Body)
}

// Get book by given isbn code.
// Return `404` if there is no such book in library.
// Return `500` on database error.
func HandleGetBooksIsbn(params books.GetBooksIsbnParams) middleware.Responder {
	filter := map[string]string{"isbn": params.Isbn}
	bs := []*models.Book{}
	if err := database.GetBooks(&bs, filter); err != nil {
		return books.NewGetBooksIsbnDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})

	}
	if len(bs) == 0 {
		return books.NewGetBooksIsbnNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String("Book not found")})
	}
	b := bs[0]
	return books.NewGetBooksIsbnOK().WithPayload(b)
}

// Delete book by given isbn code.
// Return `204` on success.
// Return `500` on database error.
func HandleDeleteBooksIsbn(params books.DeleteBooksIsbnParams) middleware.Responder {
	if err := database.DeleteBook(params.Isbn); err != nil {
		return books.NewDeleteBooksIsbnDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return books.NewDeleteBooksIsbnNoContent()
}

// Update book by given isbn code and body of information for updating.
// Return `200` on success, and given book.
// Return `500` on database error.
func HandlePutBooksIsbn(params books.PutBooksIsbnParams) middleware.Responder {
	if err := database.UpdateBook(params.Body, params.Isbn); err != nil {
		return books.NewPutBooksIsbnDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return books.NewPutBooksIsbnOK().WithPayload(params.Body)
}

// Get collections by given name, otherwise return all collections.
// Return `200` on success, and one or more collections.
// Return `404` if There is no collection for given name.
// Return `500` on database error.
func HandleGetCollections(params collections.GetCollectionsParams) middleware.Responder {
	cs := []*models.Collection{}
	if err := database.GetCollections(&cs); err != nil {
		return collections.NewGetCollectionsDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})

	}
	if params.Name != nil {
		for _, c := range cs {
			if *c.Name == *params.Name {
				return collections.NewGetCollectionsOK().WithPayload([]*models.Collection{c})
			}
		}
		return collections.NewGetCollectionsNameNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String("Collection not found")})
	}
	return collections.NewGetCollectionsOK().WithPayload(cs)
}

// Post the collection by inserting it into database.
// Return `500` on database error.
func HandlePostCollections(params collections.PostCollectionsParams) middleware.Responder {
	if err := database.InsertCollection(params.Body); err != nil {
		return collections.NewPostCollectionsDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return collections.NewPostCollectionsOK().WithPayload(params.Body)
}

// Delete the collection by given name.
// Return `500` on database error.
func HandleDeleteCollections(params collections.DeleteCollectionsParams) middleware.Responder {
	if err := database.DeleteCollection(params.Name); err != nil {
		return collections.NewDeleteCollectionsDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return collections.NewDeleteCollectionsNoContent()
}

// Update collection by given collection name and body of information for updating.
// Return `200` on success, and given collection.
// Return `500` on database error.
func HandlePutCollections(params collections.PutCollectionsParams) middleware.Responder {
	if err := database.UpdateCollection(params.Body, params.Name); err != nil {
		return collections.NewPutCollectionsDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return collections.NewPutCollectionsOK().WithPayload(params.Body)
}
