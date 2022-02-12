package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/bells307/cyoa/story"
)

func Run(port int, arcs story.ArcMap, entrypoint string) error {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: server{arcs: arcs, entrypoint: entrypoint},
	}

	log.Printf("Start listening on %v\n", s.Addr)

	return s.ListenAndServe()
}

type server struct {
	arcs       story.ArcMap
	entrypoint string
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		http.Redirect(w, r, s.entrypoint, http.StatusFound)
		return
	}

	path = strings.TrimPrefix(path, "/")
	if arc, ok := s.arcs[path]; ok {
		tmpl, err := template.ParseFiles("views/story_page.html")
		if err != nil {
			log.Printf("Error while parsing story template file: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		if err := tmpl.Execute(w, arc); err != nil {
			log.Printf("Error while executing story template: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%s", err)))
		}
	} else {
		http.NotFound(w, r)
	}
}
