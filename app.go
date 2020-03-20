package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"
const databaseName = "vidly-go"

// App defines the application
type App struct {
	Router *mux.Router
	Db     *mongo.Database
}

// Run reciever starts to listen
// and serve the application
func (a *App) Run(addr string) {
	fmt.Println("Listening on port", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Initialise reciever initialises
// the application
func (a *App) Initialise() {
	var err error
	ctx, _ := dbContext(10)
	a.Db, err = a.connectDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	a.Router = mux.NewRouter()
	a.initialiseRoutes()
}

func (a *App) connectDatabase(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	database := client.Database(databaseName)
	return database, nil
}

func (a *App) initialiseRoutes() {
	a.Router.HandleFunc("/api/genres", a.getGenres).Methods("GET")
	a.Router.HandleFunc("/api/genres/{id}", a.getGenre).Methods("GET")
	a.Router.HandleFunc("/api/genres", a.postGenre).Methods("POST")
	a.Router.HandleFunc("/api/genres/{id}", a.deleteGenre).Methods("DELETE")
}
