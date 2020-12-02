// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import "testing"

func TestPasswordEntry_IsValid(t *testing.T) {
	type fields struct {
		MinCount int
		MaxCount int
		Letter   rune
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test 1", fields{15, 16, 'f', "ffffffffffffffhf"}, true},
		{"test 2", fields{6, 8, 'b', "bbbnvbbb"}, true},
		{"test 3", fields{6, 10, 'z', "zhzzzzfzzzzzzzzzpzz"}, false},
		{"test 4", fields{9, 13, 's', "dmssskssqsssfff"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PasswordEntry{
				MinCount: tt.fields.MinCount,
				MaxCount: tt.fields.MaxCount,
				Letter:   tt.fields.Letter,
				Password: tt.fields.Password,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("PasswordEntry.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PasswordEntry{15, 16, 'f', "ffffffffffffffhf"}.IsValid()
	}
}

func TestPasswordEntry_IsValidNew(t *testing.T) {
	type fields struct {
		MinCount int
		MaxCount int
		Letter   rune
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test 1", fields{MinCount: 1, MaxCount: 3, Letter: 'a', Password: "abcde"}, true},
		{"test 2", fields{MinCount: 1, MaxCount: 3, Letter: 'b', Password: "cdefg"}, false},
		{"test 3", fields{MinCount: 2, MaxCount: 9, Letter: 'c', Password: "ccccccccc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PasswordEntry{
				MinCount: tt.fields.MinCount,
				MaxCount: tt.fields.MaxCount,
				Letter:   tt.fields.Letter,
				Password: tt.fields.Password,
			}
			if got := p.IsValidNew(); got != tt.want {
				t.Errorf("PasswordEntry.IsValidNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
