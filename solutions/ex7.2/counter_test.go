package ex72

import (
	"bytes"
	"testing"
)



func TestCountingWriter(t *testing.T) {
	tests := []struct{
		name string
		input []byte
		want int64
	}{
		{name: "empty", input: []byte{}, want: 0},
		{name: "simple", input: []byte("hi there!"), want: 9},
	}
	for _, tc := range tests {
		b := &bytes.Buffer{}
		c, n := CountingWriter(b)
		c.Write(tc.input)
		if *n != tc.want{
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, *n)
		}
	}
	 
}