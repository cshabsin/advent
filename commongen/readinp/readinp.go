package readinp

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Line is the value yielded by read - it either contains contents or an error.
type Line[T any] struct {
	Contents T
	Error    error
}

type Parser[T any] func(c string) (T, error)

// Read starts a goroutine that yields lines on the output channel.
func Read[T any](filename string, parser Parser[T]) (chan Line[T], error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	rdr := bufio.NewReader(f)
	if err != nil {
		return nil, err
	}
	ch := make(chan Line[T])
	go func() {
		for {
			line, err := rdr.ReadString('\n')
			if err == io.EOF {
				close(ch)
				return
			}
			t, err := parser(strings.TrimSpace(line))
			ch <- Line[T]{Contents: t, Error: err}
			if err != nil {
				close(ch)
				return
			}
		}
	}()
	return ch, nil
}
