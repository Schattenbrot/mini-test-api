package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes handles all the URL-routing.
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1//posts", app.createPost)

	return router
}
