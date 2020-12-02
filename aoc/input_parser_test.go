// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import (
	"reflect"
	"testing"
)

func TestParseRow(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name    string
		args    args
		want    PasswordEntry
		wantErr bool
	}{
		{"test a", args{"123456"}, PasswordEntry{}, true},
		{"test b", args{"abcdef"}, PasswordEntry{}, true},
		{"test c", args{"15- f: ffffffffffffffhf"}, PasswordEntry{}, true},
		{"test d", args{"-16 f: ffffffffffffffhf"}, PasswordEntry{}, true},
		{"test e", args{"15-16 : ffffffffffffffhf"}, PasswordEntry{}, true},
		{"test f", args{"15 16 f: ffffffffffffffhf"}, PasswordEntry{}, true},
		{"test g", args{"15-16 : "}, PasswordEntry{}, true},
		{"test 1", args{"15-16 f: ffffffffffffffhf"}, PasswordEntry{15, 16, 'f', "ffffffffffffffhf"}, false},
		{"test 2", args{"6-8 b: bbbnvbbb"}, PasswordEntry{6, 8, 'b', "bbbnvbbb"}, false},
		{"test 3", args{"6-10 z: zhzzzzfzzzzzzzzzpzz"}, PasswordEntry{6, 10, 'z', "zhzzzzfzzzzzzzzzpzz"}, false},
		{"test 4", args{"9-13 s: dmssskssqsssssf"}, PasswordEntry{9, 13, 's', "dmssskssqsssssf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRow(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
