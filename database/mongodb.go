package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-gql-mongo/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var url string = "mongodb://localhost:27017"

type MongoDB struct {
	Client *mongo.Client
}

// ************ CONNECT to mongodb
func ConnectToMongoDb() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// CHECK CONNECTION
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- App has been Connected to MongoDB!(Graphql)")
	return &MongoDB{Client: client}
}
func (db *MongoDB) Createbook(input model.CreateBookInput) *model.Book {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	booksCollection := db.Client.Database("library-db").Collection("newbooks")
	insertResult, err := booksCollection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return &model.Book{
		ID:     insertedID,
		Title:  input.Title,
		Author: input.Author,
		Price:  input.Price,
	}
}

// UpdateBook is the resolver for the updateBook field.
func (db *MongoDB) Updatebook(id string, input model.UpdateBookInput) *model.Book {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	booksCollection := db.Client.Database("library-db").Collection("newbooks")
	update := bson.M{}

	if input.Title != nil {
		update["title"] = input.Title
	}
	if input.Author != nil {
		update["author"] = input.Author
	}
	if input.Price != nil {
		update["price"] = input.Price
	}
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: _id}}
	ubdatedBook := bson.M{"$set": update}

	updateResult := booksCollection.FindOneAndUpdate(ctx, filter, ubdatedBook, options.FindOneAndUpdate().SetReturnDocument(1))

	var book model.Book
	if err := updateResult.Decode(&book); err != nil {
		log.Fatal(err)
	}
	return &book
}

// DeleteBook is the resolver for the deleteBook field.
func (db *MongoDB) Deletebook(id string) *model.DeleteBookResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	booksCollection := db.Client.Database("library-db").Collection("newbooks")
	_id, _ := primitive.ObjectIDFromHex(id)

	// deleteResult, err := booksCollection.DeleteMany(ctx, bson.D{{Key: "_id", Value: _id}})
	deleteResult, err := booksCollection.DeleteOne(ctx, bson.D{{Key: "_id", Value: _id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("This book has been deleted successfully:: ", deleteResult)
	return &model.DeleteBookResponse{DeleteBookID: id}
}

// Books is the resolver for the books field.
func (db *MongoDB) GetBooks() []*model.Book {
	var results []*model.Book
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	booksCollection := db.Client.Database("library-db").Collection("newbooks")

	findOptions := options.Find()
	findOptions.SetLimit(100)

	// Finding multiple documents returns a cursor
	cur, err := booksCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(ctx) {
		var elem model.Book
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)
	return results
}

// Book is the resolver for the book field.
func (db *MongoDB) GetBookByID(id string) *model.Book {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	booksCollection := db.Client.Database("library-db").Collection("newbooks")
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: _id}}

	var result model.Book

	err := booksCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}
