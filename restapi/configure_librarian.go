// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	database "librarian/database"
	"librarian/restapi/operations"
	"librarian/restapi/operations/books"
	"librarian/restapi/operations/collections"
)

//go:generate swagger generate server --target ../../librarian --name Librarian --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.LibrarianAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
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

	// Initialize database.
	dbName := os.Getenv("DB_PATH")
	if len(dbName) == 0 {
		dbName = "database/librarian.db"
		log.Printf("Using the default db: %s\b", dbName)
	}
	err := database.InitDB(dbName)
	if err != nil {
		log.Printf("Failed to initialize db %s\n", err)
	}

	// Set handlers.
	api.BooksGetBooksHandler = books.GetBooksHandlerFunc(HandleGetBooks)
	api.BooksPostBooksHandler = books.PostBooksHandlerFunc(HandlePostBooks)
	api.BooksGetBooksIsbnHandler = books.GetBooksIsbnHandlerFunc(HandleGetBooksIsbn)
	api.BooksDeleteBooksIsbnHandler = books.DeleteBooksIsbnHandlerFunc(HandleDeleteBooksIsbn)
	api.BooksPutBooksIsbnHandler = books.PutBooksIsbnHandlerFunc(HandlePutBooksIsbn)

	api.CollectionsGetCollectionsHandler = collections.GetCollectionsHandlerFunc(HandleGetCollections)
	api.CollectionsPostCollectionsHandler = collections.PostCollectionsHandlerFunc(HandlePostCollections)
	api.CollectionsDeleteCollectionsHandler = collections.DeleteCollectionsHandlerFunc(HandleDeleteCollections)
	api.CollectionsPutCollectionsHandler = collections.PutCollectionsHandlerFunc(HandlePutCollections)
	//

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
