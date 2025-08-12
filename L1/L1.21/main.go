package main

import (
	"fmt"
	"io"
	"strings"
)

// Адаптируем string к io.Reader
type StringReaderAdapter struct {
	s string
	r *strings.Reader
}

func NewStringReaderAdapter(s string) *StringReaderAdapter {
	return &StringReaderAdapter{
		s: s,
		r: strings.NewReader(s),
	}
}

func (a *StringReaderAdapter) Read(p []byte) (n int, err error) {
	return a.r.Read(p)
}

func main() {
	str := "Hello, Adapter!"
	adapter := NewStringReaderAdapter(str)

	buf := make([]byte, len(str))
	n, err := adapter.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println("readed:", string(buf[:n]))
}
