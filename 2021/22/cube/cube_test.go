package cube

import (
	"fmt"
	"testing"

	"github.com/cshabsin/advent/commongen/matrix"
)

func TestSubtract(t *testing.T) {
	c := Cube{Min: matrix.Point3{10, 10, 10}, Max: matrix.Point3{10, 12, 12}}
	d := Cube{Min: matrix.Point3{9, 9, 9}, Max: matrix.Point3{11, 11, 11}}
	fmt.Println(c, c.Volume())
	fmt.Println(d)
	sub := c.Subtract(d)
	vol := 0
	for _, s := range sub {
		vol += int(s.Volume())
	}
	fmt.Println(sub, vol)
}
