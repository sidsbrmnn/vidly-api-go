package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Genre defines a genre document
type Genre struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

// Genres define an array of 'Genre'
type Genres []Genre

func findAllGenres(db *mongo.Database) (Genres, error) {
	collection := db.Collection("genre")
	ctx, cancel := dbContext(30)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	genres := Genres{}
	for cur.Next(ctx) {
		var genre Genre
		cur.Decode(&genre)
		genres = append(genres, genre)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

func (g *Genre) findOneGenre(db *mongo.Database) error {
	collection := db.Collection("genre")
	ctx, cancel := dbContext(5)
	defer cancel()

	filter := bson.M{"_id": g.ID}
	err := collection.FindOne(ctx, filter).Decode(&g)
	return err
}

func (g *Genre) createGenre(db *mongo.Database) (*mongo.InsertOneResult, error) {
	collection := db.Collection("genre")
	ctx, cancel := dbContext(30)
	defer cancel()

	result, err := collection.InsertOne(ctx, g)
	return result, err
}

func (g *Genre) deleteOneGenre(db *mongo.Database) error {
	collection := db.Collection("genre")
	ctx, cancel := dbContext(5)
	defer cancel()

	filter := bson.M{"_id": g.ID}
	err := collection.FindOneAndDelete(ctx, filter).Decode(&g)
	return err
}
