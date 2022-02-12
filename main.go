package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/bells307/cyoa/cli"
	"github.com/bells307/cyoa/server"
	"github.com/bells307/cyoa/story"
)

var (
	runAsServer bool
	port        int
	file        string
	entrypoint  string
)

func init() {
	flag.BoolVar(&runAsServer, "s", false, "Run as server")
	flag.IntVar(&port, "p", 8888, "Server listening port")
	flag.StringVar(&file, "f", "gopher.json", "Data file path")
	flag.StringVar(&entrypoint, "e", "intro", "Story entry point")
	flag.Parse()
}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error while reading %s: %v", file, err)
	}

	var arcs story.ArcMap = make(story.ArcMap)
	err = json.Unmarshal(data, &arcs)
	if err != nil {
		log.Fatalf("Error while parsing %s: %v", file, err)
	}

	if _, ok := arcs[entrypoint]; !ok {
		log.Fatalf("Story entrypoint \"%s\" not found", entrypoint)
	}

	if runAsServer {
		log.Println("Running application as server")

		if err = server.Run(port, arcs, entrypoint); err != nil {
			log.Fatalf("Error running as server: %v", err)
		}
	} else {
		log.Println("Running application as CLI")

		if err = cli.Run(arcs, entrypoint); err != nil {
			log.Fatalf("Error running as CLI: %v", err)
		}
	}
}
