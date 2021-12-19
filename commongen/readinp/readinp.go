package readinp

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func (l Line[T]) MustGet() T {
	if l.Error != nil {
		log.Fatal(l.Error)
	}
	return l.Contents
}

type Consumer[T any] interface {
	// Parse returns the object to return and true, for lines that yield an
	// object, or an empty/nil object and false, for lines that should not
	// yield a result.
	Parse(string) (T, bool, error)
}

type consumerFunc[T any] func(string) (T, error)

func (c consumerFunc[T]) Parse(fn string) (T, bool, error) {
	v, err := c(fn)
	return v, true, err
}

func Read[T any](filename string, parser func(c string) (T, error)) (chan Line[T], error) {
	return ReadConsumer[T](filename, consumerFunc[T](parser))
}

func MustRead[T any](filename string, parser func(c string) (T, error)) chan Line[T] {
	ch, err := Read(filename, parser)
	if err != nil {
		log.Fatal(err)
	}
	return ch
}

// Read starts a goroutine that yields lines on the output channel.
func ReadConsumer[T any](filename string, consumer Consumer[T]) (chan Line[T], error) {
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
			t, ret, err := consumer.Parse(strings.TrimSpace(line))
			if ret {
				ch <- Line[T]{Contents: t, Error: err}
			}
		}
	}()
	return ch, nil
}

func S(s string) (string, error) {
	return s, nil
}

// Atoi is a wrapper around strconv.Atoi that logs fatal on error.
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
