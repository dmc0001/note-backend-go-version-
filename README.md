# Note backend

This project is a simple notes management server built using Go's standard library. It supports basic CRUD operations for notes, including adding, updating, deleting, and fetching notes. The purpose of this project is to explore Go's capabilities for building RESTful APIs.

## Setup

Setting this up locally is trivial

* `$ git clone https://github.com/dmc0001/note-backend`
* `cd note-backend`
* `go run main.go`

## How to use it

This backend supports five actions: "create-note", "delete-node", "update-note", "fetch-note", and "list-notes"

### api/v1/notes/list (GET)

```
curl -X GET http://0.0.0.0:8080/api/v1/notes/list 
```


### Response body:
```
[
    {
        "id": 1,
        "title": "note 1",
        "description": "description note 1"
    },
    {
        "id": 2,
        "title": "note 2",
        "description": "description note 2"
    }
]

```
### Expected Status

* 404
* 405
* 200


### api/v1/notes/find (GET)

```
curl -X GET "http://0.0.0.0:8080/api/v1/notes/find?id=1"
```

### Response body:
```
{
  "id": 1,
  "title": "note 1",
  "description": "description note 1"
}

```
### Expected Status

* 404
* 405
* 200


### api/v1/notes/create (POST)

```
curl -X POST http://0.0.0.0:8080/api/v1/notes/create -d { "note":{ "id": 2, "title": "note 2","description": "description note 2"}}
```

### Request body:
```
{
  "note": { "id": 2,
    "title": "note 2",
    "description": "description note 2"
  }
}
```
### Expected Status
 
* 404 
* 405 
* 201 


### api/v1/notes/update (PUT)

```
curl -X PUT http://0.0.0.0:8080/api/v1/notes/update -d {  "note":{"id":2 title": "note 2","description": "description note 2"}}
```

### Request body:
```
{
  "note": { "id": 2,
    "title": "note 2",
    "description": "description note 2"
  }
}
```
### Expected Status

* 404 
* 405 
* 200


### api/v1/notes/delete (DELETE)

```
curl -X DELETE http://0.0.0.0:8080/api/v1/notes/delete -d {  "note":{"id":2 title": "note 2","description": "description note 2"}}
```

### Request body:
```
{
  "note": { "id": 2,
    "title": "note 2",
    "description": "description note 2"
  }
}
```
### Expected Status

* 404
* 405
* 200


## Features

Here's a list of features included in this project:

| Name                                                                   | Description                                                                        |
|------------------------------------------------------------------------|------------------------------------------------------------------------------------|
| [Routing](https://pkg.go.dev/net/http)                             | Uses Go's net/http package to define and handle routes.                                                |
| [JSON Encoding/Decoding](https://pkg.go.dev/encoding/json) | Handles request and response bodies using Go's encoding/json package.                  |
| [Error Handling](https://go.dev/blog/error-handling-and-go)               | Responds with appropriate HTTP status codes for various scenarios.                                      |



