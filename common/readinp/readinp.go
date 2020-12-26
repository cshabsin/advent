package readinp

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Line is the value yielded by read - it either contains contents or an error.
type Line struct {
	Contents *string
	Error    error
}

// Value returns the trimmed contents of the line. Crashes on an error line.
func (l Line) Value() string {
	return strings.TrimSpace(*l.Contents)
}

// Read starts a goroutine that yields lines on the output channel.
func Read(filename string) (chan Line, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	rdr := bufio.NewReader(f)
	if err != nil {
		return nil, err
	}
	ch := make(chan Line)
	go func() {
		for {
			line, err := rdr.ReadString('\n')
			if err == io.EOF {
				close(ch)
				return
			}
			if err != nil {
				ch <- Line{Error: err}
				close(ch)
				return
			}
			ch <- Line{Contents: &line}
		}
	}()
	return ch, nil
}
