# REST API

## Introduction

Communications between Librarian and its clients can happen in two ways:
+ Using a RESTful API over HTTP.
+ Local CLI client, which doesn't assumes no authentication.


## Return values

There are three standard return types:

* Standard return value
* Error

### Standard return value

For a standard synchronous operation, the following dict is returned:

```js
{
    "status": "Success",
    "status_code": 200,
}
```

### Error

There are various situations in which something may immediately go
wrong, in those cases, the following return value is used:

```js
{
    "type": "error",
    "error": "Failure",
    "error_code": 400,
    "metadata": {}                      // More details about the error
}
```

TODO HTTP code must be one of of 400, 401, 403, 404, 409, 412 or 500.

## Status codes

The Librarian REST API often has to return status information, be that the
reason for an error, the current state of an operation or the state of
the various resources it exports.

To make it simple to debug, all of those are always doubled. There is a
numeric representation of the state which is guaranteed never to change
and can be relied on by API clients. Then there is a text version meant
to make it easier for people manually using the API to figure out what's
happening.

In most cases, those will be called status and `status_code`, the former
being the user-friendly string representation and the latter the fixed
numeric value.

The codes are always 3 digits, with the following ranges:

* 100 to 199: resource state (started, stopped, ready, ...)
* 200 to 399: positive action result
* 400 to 599: negative action result
* 600 to 999: future use

### List of current status codes

Code  | Meaning
:---  | :------
200   | Success
400   | Failure
401   | Canceled

## Asynchronous operations

Any operation which may take more than a second to be done must be done
in the background, returning a background operation ID to the client.

The client will then be able to either poll for a status update or wait
for a notification using the long-poll API.

## Notifications

A WebSocket-based API is available for notifications, different notification
types exist to limit the traffic going to the client.

It's recommended that the client always subscribes to the operations
notification type before triggering remote operations so that it doesn't
have to then poll for their status.

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


## PUT vs PATCH

The Librarian API supports both PUT and PATCH to modify existing objects.

PUT replaces the entire object with a new definition, it's typically
called after the current object state was retrieved through GET.

PATCH can be used to modify a single field inside an object by only
specifying the property that you want to change. 


## API structure

Librarian has an auto-generated [Swagger](https://swagger.io/) specification describing its API endpoints.
The YAML version of this API specification can be found in [`rest-api.yaml`](rest-api.yaml).
