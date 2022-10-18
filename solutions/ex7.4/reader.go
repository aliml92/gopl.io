package ex74

import "io"


type stringReader struct {
	val string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.val)
	r.val = r.val[n:]
	if len(r.val) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}