package handler

import (
	"encoding/json"
	"gophercises/adventure/story"
	"log"
	"net/http"
	"strings"
)

type routeHandler struct {
	storyParts map[string]story.StoryPart
}

func NewRouteHandler(storyParts map[string]story.StoryPart) *routeHandler {
	h := routeHandler{storyParts}
	return &h
}

func (h routeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if storyPart, ok := h.storyParts[path]; ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		resp, err := json.Marshal(storyPart)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(resp)
	} else {
		http.NotFound(w, r)
	}
}
