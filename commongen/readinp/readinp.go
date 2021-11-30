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

func (l Line[T]) Get() (T, error) {
	return l.Contents, l.Error
}

// Read starts a goroutine that yields lines on the output channel.
func Read[T any](filename string, parser func(c string) (T, error)) (chan Line[T], error) {
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
			line, readErr := rdr.ReadString('\n')
			line = strings.TrimSpace(line)
			if readErr == io.EOF && line == "" {
				close(ch)
				return
			} else if readErr != nil && readErr != io.EOF {
				ch <- Line[T]{Error: readErr}
				close(ch)
				return
			}
			t, err := parser(line)
			ch <- Line[T]{Contents: t, Error: err}
			if err != nil || readErr == io.EOF {
				close(ch)
				return
			}
		}
	}()
	return ch, nil
}
