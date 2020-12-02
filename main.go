package main

import (
	"log"
	"os"

	"github.com/nmueller-realdigital/aoc2020-day02/aoc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	log.Println("task completed successfully")
}

func run() error {
	file, err := os.Open("data/input.txt")
	if err != nil {
		return err
	}
	entries, err := aoc.ReadAndParseInput(file)
	if err != nil {
		return err
	}
	var countValid int
	for _, entry := range entries {
		if entry.IsValid() {
			countValid++
		}
	}
	log.Printf("count of valid passwords: %d\n", countValid)
	var countValidNew int
	for _, entry := range entries {
		if entry.IsValidNew() {
			countValidNew++
		}
	}
	log.Printf("count of valid passwords under new policies: %d\n", countValidNew)
	return nil
}
