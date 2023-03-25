// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	database "librarian/database"
	"librarian/models"
	"librarian/restapi/operations"
	"librarian/restapi/operations/books"
	"librarian/restapi/operations/collections"
)

//go:generate swagger generate server --target ../../librarian --name Librarian --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.LibrarianAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func createBook(params books.CreateBookParams) error {
	return database.InsertBook(params.Body)
}

func getBooks(books *[]*models.Book) (err error) {
	return database.GetBooks(books)
}

func configureAPI(api *operations.LibrarianAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	err := database.InitDB()
	if err != nil {
		log.Println("Failed to initialize db", err)
	}

	api.BooksGetBooksHandler = books.GetBooksHandlerFunc(func(params books.GetBooksParams) middleware.Responder {
		bs := []*models.Book{}
		if err := getBooks(&bs); err != nil {
			return middleware.NotImplemented("Failed to get books")
		}

		return books.NewGetBooksOK().WithPayload(bs)
	})

	if api.BooksGetBooksIsbnHandler == nil {
		api.BooksGetBooksIsbnHandler = books.GetBooksIsbnHandlerFunc(func(params books.GetBooksIsbnParams) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooksIsbn has not yet been implemented")
		})
	}
	if api.GetCollectionsHandler == nil {
		api.GetCollectionsHandler = operations.GetCollectionsHandlerFunc(func(params operations.GetCollectionsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCollections has not yet been implemented")
		})
	}

	api.BooksCreateBookHandler = books.CreateBookHandlerFunc(func(params books.CreateBookParams) middleware.Responder {
		if err := createBook(params); err != nil {
			return books.NewCreateBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return books.NewCreateBookOK().WithPayload(params.Body)
	})

	if api.CollectionsCreateCollectionHandler == nil {
		api.CollectionsCreateCollectionHandler = collections.CreateCollectionHandlerFunc(func(params collections.CreateCollectionParams) middleware.Responder {
			return middleware.NotImplemented("operation collections.CreateCollection has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
