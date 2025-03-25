package main

import (
	"flag"
	"go-print/src"
	"log"
)

func main() {
	configPath := flag.String("config", "print.yaml", "Path to the config file")
	flag.Parse()

	cfg, err := src.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	files, err := src.GetFiles(".", cfg)
	if err != nil {
		log.Fatalf("Failed to get files: %v", err)
	}

	if err := src.PrintMarkdown(cfg.OutputPath, files); err != nil {
		log.Fatalf("Failed to print markdown: %v", err)
	}
}
