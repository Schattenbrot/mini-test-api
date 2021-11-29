package main

import (
	"encoding/json"
	"net/http"

	"github.com/mini-test-api/models"
)

type PostPayload struct {
	ID   string `json:"_id"`
	Text string `json:"text"`
}

// createPost is the handler for the InsertPost method.
func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	var payload PostPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var post models.Post
	post.Text = payload.Text

	type jsonResp struct {
		OK bool   `json:"ok"`
		ID string `json:"_id"`
	}

	id, err := app.models.DB.InsertPost(post)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
		ID: id.Hex(),
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
