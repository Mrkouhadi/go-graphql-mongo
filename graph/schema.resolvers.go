package graph

import (
	"context"
	"go-gql-mongo/database"
	"go-gql-mongo/graph/model"
)

var db = database.ConnectToMongoDb()

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.CreateBookInput) (*model.Book, error) {
	return db.Createbook(input), nil
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input model.UpdateBookInput) (*model.Book, error) {
	return db.Updatebook(id, input), nil
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.DeleteBookResponse, error) {
	return db.Deletebook(id), nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return db.GetBooks(), nil
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	return db.GetBookByID(id), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
