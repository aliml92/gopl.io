package ex71

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{name: "empty", input: []byte(""), want: 0},
		{name: "simple", input: []byte("first line\nsecond line\nthird line\n"), want: 3},
	}
	for _, tc := range tests {
		var c LineCounter
		n, err := c.Read(tc.input)
		if n != len(tc.input) {
			t.Errorf("input length: %d != %d ", len(tc.input), n)
		}
		if err != nil {
			t.Errorf("err: %v", err)
		}
		if tc.want != int(c) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, c)
		}
	}
}

func TestWordCounter(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{name: "empty", input: []byte(""), want: 0},
		{name: "simple", input: []byte("first line\nsecond line\nthird line\n"), want: 6},
	}
	for _, tc := range tests {
		var c WordCounter
		n, err := c.Read(tc.input)
		if n != len(tc.input) {
			t.Errorf("input length: %d != %d ", len(tc.input), n)
		}
		if err != nil {
			t.Errorf("err: %v", err)
		}
		if tc.want != int(c) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, c)
		}
	}
}
