package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	models := flag.Args()

	// TOOD: better validation of model names
	if len(models) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	pkg := os.Getenv("GOPACKAGE")
	if pkg == "" {
		fmt.Println("The environmental variable $GOPACKAGE MUST be set (and is set automatically by go generate)")
		os.Exit(1)
	}

	results := make(chan error)
	for i := range models {
		model := models[i]
		go func() {
			filename := strings.ToLower(model) + ".gen.go"
			results <- writeRepoToFile(filename, model, pkg)
		}()
	}
	errCount := 0
	for i := 0; i < len(models); i++ {
		if err := <-results; err != nil {
			fmt.Fprintln(os.Stderr, err)
			errCount += 1
		}
	}
	if errCount > 0 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
