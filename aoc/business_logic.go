// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

// PasswordEntry represents a row in the password input file.
type PasswordEntry struct {
	MinCount int
	MaxCount int
	Letter   rune
	Password string
}

// IsValid returns true if Letter appears at least MinCount times and at maximum MaxCount times in Password.
func (p PasswordEntry) IsValid() bool {
	var count int
	for _, letter := range p.Password {
		if letter == p.Letter {
			count++
		}
	}
	return count >= p.MinCount && count <= p.MaxCount
}

// IsValidNew returns true if Letter appears either at position MinCount or at position MaxCount in Password.
func (p PasswordEntry) IsValidNew() bool {
	// Toboggan Corporate Policies have no concept of "index zero", hence shift index by 1
	positionOne := p.MinCount - 1
	positionTwo := p.MaxCount - 1
	runes := []rune(p.Password)
	return (runes[positionOne] == p.Letter || runes[positionTwo] == p.Letter) && !(runes[positionOne] == p.Letter && runes[positionTwo] == p.Letter)
}
