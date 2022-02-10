package server

import (
	"fmt"
	"gophercises/adventure/story"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func Run(sp story.StoryPartMap) error {
	s := &http.Server{
		Addr:    ":8080",
		Handler: routeHandler{sp},
	}

	log.Printf("Start listening on %v\n", s.Addr)

	return s.ListenAndServe()
}

type routeHandler struct {
	storyParts story.StoryPartMap
}

func (h routeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if sp, ok := h.storyParts[path]; ok {
		tmpl, err := template.ParseFiles("views/story.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		tmpl.Execute(w, sp)
	} else {
		http.NotFound(w, r)
	}
}
