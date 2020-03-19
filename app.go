package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	database := client.Database("")
	return database, nil
}

func (a *App) initialiseRoutes() {}

func dbContext(i time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), i*time.Second)
	return ctx, cancel
}
