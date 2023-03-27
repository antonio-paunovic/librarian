# REST API

## Introduction

This API implements handling of books and collections of books.

## API Structure
#### Book
    + GET get books depending on filter or a specific book by isbn code.
    + POST create a book
    + PUT update a book by isbn
    + DELETE delete a book by isbn

#### Collection
    + GET get collections depending on filter or a specific collection by the collection name.
    + POST create a collection
    + PUT update a collection by name
    + DELETE delete a collection by name

## Return values

Various status codes are returned depending on the operation and its success.
Back-end returns code '500' when database can't execute necessary operation.
It also return the corresponding error message which is useful for debugging.

### List of current status codes

Code  | Meaning
:---  | :------
200   | Success
204   | Successful delete
400   | Failure
404   | Resource not found
500   | Internal server error


## Filtering

To filter your results on certain values, filter is implemented for collections.
A `filter` argument can be passed to a GET query against a collection.

Filtering is available for the books and collections of books.

There is no default value for filter which means that all results found will
be returned. The following is the language used for the filter argument:

    ?filter=field_name eq desired_field_assignment

The language follows the OData conventions for structuring REST API filtering
logic. Logical operators are also supported for filtering: equals (`eq`) and (`and`).
For instance, to filter books on genre and author you would pass:

    ?filter=genre eq book_genre and author eq book_author

Here is an example of GET query for the filtering methods:

    ?filter=genre eq book_genre and author eq book_author


## API structure

Librarian has specification describing its API endpoints from which the code is generated at `swagger.yml`.
Additional documentation could be generated from it.
