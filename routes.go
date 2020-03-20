package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Route defines a single endpoint
type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}

// Subrouter defines a set of routes
// for a single path prefix
type Subrouter struct {
	Routes []Route
	Prefix string
}

func (a *App) getGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := findAllGenres(a.Db)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, genres)
}

func (a *App) getGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid genre id")
		return
	}

	g := Genre{ID: id}
	err = g.findOneGenre(a.Db)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			writeError(w, http.StatusNotFound, "Genre not found")
		default:
			writeError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	writeResponse(w, http.StatusOK, g)
}

func (a *App) postGenre(w http.ResponseWriter, r *http.Request) {
	g := Genre{}

	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	result, err := g.createGenre(a.Db)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, result)
}

func (a *App) deleteGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid genre id")
		return
	}

	g := Genre{ID: id}
	err = g.deleteOneGenre(a.Db)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			writeError(w, http.StatusNotFound, "Genre not found")
		default:
			writeError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	writeResponse(w, http.StatusOK, g)
}
