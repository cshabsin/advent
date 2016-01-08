package main

import (
	"flag"
	"fmt"
	"strings"
)

func has_illegal_letter(password string) bool {
	if strings.IndexByte(password, 'i') != -1 {
		return true
	}
	if strings.IndexByte(password, 'o') != -1 {
		return true
	}
	if strings.IndexByte(password, 'l') != -1 {
		return true
	}
	return false
}

func has_straight(password string) bool {
	runlen := 0
	var prev byte
	for i:=0; i<len(password); i++ {
		if password[i] == prev+1 {
			runlen++
		} else {
			runlen = 0
		}
		if runlen == 2 {
			return true
		}
		prev = password[i]
	}
	return false
}

func has_two_doubles(password string) bool {
	i := 0
	found := false
	var prev byte
	for ; i<len(password); i++ {
		if password[i] == prev {
			found = true
			break
		}
		prev = password[i]
	}
	if !found {
		return false
	}
	var newprev byte
	for ; i<len(password); i++ {
		if password[i] != prev && password[i] == newprev {
			return true
		}
		newprev = password[i]
	}
	return false
}

func incr_letter(password string, index int) string {
	var new_password string
	if password[index] == 'z' {
		new_password = password[:index] + "a" + password[index+1:]
		return incr_letter(new_password, index-1)
	}
	c := password[index] + 1
	if c == 'i' || c == 'o' || c == 'l' {
		c++
	}
	new_password = fmt.Sprintf("%s%c%s", password[:index], c, password[index+1:])
	return new_password
}

func NextPassword(password string) string {
	for {
		password = incr_letter(password, len(password)-1)
		if has_illegal_letter(password) {
			continue
		}
		if !has_straight(password) {
			continue
		}
		if !has_two_doubles(password) {
			continue
		}
		return password
	}
}

func main() {
	var password = flag.String("password", "hxbxwxba", "Original password")
	flag.Parse()
	fmt.Print(*password, " -> ", NextPassword(*password), "\n")
}
