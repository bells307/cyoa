package main

import (
	"encoding/json"
	"flag"
	"gophercises/adventure/handler"
	"gophercises/adventure/story"
	"log"
	"net/http"
	"os"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "gopher.json", "Data file path")
	flag.Parse()
}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error while reading %s: %v", file, err)
	}

	var storyParts map[string]story.StoryPart = make(map[string]story.StoryPart)
	err = json.Unmarshal(data, &storyParts)
	if err != nil {
		log.Fatalf("Error while parsing %s: %v", file, err)
	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler.NewRouteHandler(storyParts),
	}

	log.Fatal(s.ListenAndServe())
}
