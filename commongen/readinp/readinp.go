package readinp

import (
	"bufio"
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
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	ch := make(chan Line[T])
	go func() {
		for {
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					ch <- Line[T]{Error: err}
				}
				close(ch)
				return
			}
			line := scanner.Text()
			t, err := parser(strings.TrimSpace(line))
			ch <- Line[T]{Contents: t, Error: err}
			if err != nil {
				// TODO: do we really want to stop on the first parse error?
				close(ch)
				return
			}
		}
	}()
	return ch, nil
}
