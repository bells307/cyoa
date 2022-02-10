package main

import (
	"encoding/json"
	"flag"
	"gophercises/adventure/cli"
	"gophercises/adventure/server"
	"gophercises/adventure/story"
	"log"
	"os"
)

var (
	file        string
	runAsServer bool
)

func init() {
	flag.StringVar(&file, "f", "gopher.json", "Data file path")
	flag.BoolVar(&runAsServer, "s", true, "Run as server")
	flag.Parse()
}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error while reading %s: %v", file, err)
	}

	var sp story.StoryPartMap = make(story.StoryPartMap)
	err = json.Unmarshal(data, &sp)
	if err != nil {
		log.Fatalf("Error while parsing %s: %v", file, err)
	}

	if runAsServer {
		log.Println("Running application as server")
		if err = server.Run(sp); err != nil {
			log.Fatalf("Error running as server: %v", err)
		}
	} else {
		log.Println("Running application as CLI")
		cli.Run(sp)
	}
}
