// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
)

var (
	// ErrPatternDoesNotMatch is returned when parsing a row with invalid content.
	ErrPatternDoesNotMatch error = errors.New("the regexp pattern does not match")
)

// ReadAndParseInput reads challenge data from an io.Reader and returns its parsed contents.
func ReadAndParseInput(r io.Reader) ([]PasswordEntry, error) {
	var entries []PasswordEntry
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		entry, err := ParseRow(scanner.Text())
		if err != nil {
			return []PasswordEntry{}, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

// ParseRow parses a row in the challenge data and compiles a PasswordEntry from it.
func ParseRow(row string) (PasswordEntry, error) {
	re, err := regexp.Compile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	if err != nil {
		return PasswordEntry{}, err
	}
	matches := re.FindStringSubmatch(row)
	if matches == nil {
		return PasswordEntry{}, ErrPatternDoesNotMatch
	}
	minCount, _ := strconv.Atoi(matches[1])
	maxCount, _ := strconv.Atoi(matches[2])
	letter := []rune(matches[3])[0]
	password := matches[4]
	return PasswordEntry{minCount, maxCount, letter, password}, nil
}
