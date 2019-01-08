package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {
	// TODO: better argument parsing
	flag.Parse()
	models := flag.Args()

	// TOOD: better validation of model names
	if len(models) == 0 {
		fmt.Println("Argukment names must be passed")
		os.Exit(1)
	}
	pkg := os.Getenv("GOPACKAGE")
	if pkg == "" {
		fmt.Println("The environmental variable $GOPACKAGE MUST be set (and is set automatically by go generate)")
		os.Exit(1)
	}

	var g errgroup.Group
	for i := range models {
		model := models[i]
		g.Go(func() error {
			filename := strings.ToLower(model) + ".gen.go"
			return writeRepoToFile(filename, model, pkg)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "go-gen-sql: %v", err)
		os.Exit(1)
	}
}
