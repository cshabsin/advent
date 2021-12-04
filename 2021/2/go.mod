module github.com/cshabsin/advent/2021/2

go 1.18

replace (
	github.com/cshabsin/advent/2021/1/day1 => ./day1
	github.com/cshabsin/advent/commongen/readinp => ../../commongen/readinp
)

require github.com/cshabsin/advent/commongen/readinp v0.0.0-20211130024607-75e1f16a053d
