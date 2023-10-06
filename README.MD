# DESCRIPTION:

- A guideline of how to use graphql and golang with the help of this package [gqlgen](https://www.github.com/99designs/gqlgen)

# Instructions

- install and set up:
- run this:

0. `go get github.com/99designs/gqlgen`
1. `printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go`
2. `go mod tidy`
3. `go run github.com/99designs/gqlgen init`
4. `go mod tidy`
5. `go run server.go`

## sample of Queries:

### create a book

```
mutation CreateBook($input: createBookInput!){
  createBook(input:$input){
    _id
    TITLE
    AUTHOR
    PRICE
	}
}
```

```
{
  "input": {
    "TITLE":  "Gomorrah",
    "AUTHOR":  "Donna Ima",
    "PRICE": 3999
  }
}
```

### get all books

```
query GetBooks{
  books{
    _id
    TITLE
    AUTHOR
    PRICE
  }
}
```

### get a single book by ID

```
query GetBookByID($id: ID!){
  book(id:$id){
    _id
    TITLE
    AUTHOR
    PRICE
  }
}
```

```
{
  "id": "651f8f0b35b8d3006fd07ec7"
}
```

### delete a book

```
mutation DeleteQuery($id: ID!) {
  deleteBook(id:$id){
    deletedBookID
  }
}
```

```
{
  "id":"651f8f0b35b8d3006fd07ec7"
}
```

### update a book

```
mutation UpdateJob($id: ID!,$input: updateBookInput!){
  updateBook(id:$id,input:$input){
    _id
    TITLE
    AUTHOR
    PRICE
  }
}
```

```
{ "id": "651f8f0b35b8d3006fd07ec7",
  "input": {
    "PRICE": 88888
  }
}
```
