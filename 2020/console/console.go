package console

import (
	"fmt"
	"strconv"
	"strings"
)

type Console struct {
	Instructions []*Instruction
}

func New() *Console {
	return &Console{}
}

func (cons *Console) ReadInstruction(line string) error {
	i, err := ReadInstruction(line)
	if err != nil {
		return err
	}
	cons.Instructions = append(cons.Instructions, i)
	return nil
}

func copyMap(m map[int]bool) map[int]bool {
	m2 := map[int]bool{}
	for i, b := range m {
		m2[i] = b
	}
	return m2
}

// Returns accumulator before end, and boolean "terminates".
func (cons *Console) Run(ip int, executed map[int]bool, changed bool) (int, bool) {
	var acc int
	if executed == nil {
		executed = map[int]bool{}
	} else {
		executed = copyMap(executed)
	}
	for {
		if ip == len(cons.Instructions) {
			return acc, true
		}
		if executed[ip] {
			return acc, false
		}
		executed[ip] = true
		instruction := cons.Instructions[ip]
		switch instruction.Operation {
		case ACC:
			acc += instruction.Argument
			ip++
		case JMP:
			if !changed {
				if subAcc, terminates := cons.Run(ip+1, executed, true); terminates {
					return acc + subAcc, true
				}
			}
			ip += instruction.Argument
		case NOP:
			if !changed {
				if subAcc, terminates := cons.Run(ip+instruction.Argument, executed, true); terminates {
					return acc + subAcc, true
				}
			}
			ip++
		}
	}
}

type OpType int

const (
	ACC OpType = iota
	JMP
	NOP
)

var OpMap = map[string]OpType{
	"nop": NOP,
	"acc": ACC,
	"jmp": JMP,
}

type Instruction struct {
	Operation OpType
	Argument  int
}

func ReadInstruction(line string) (*Instruction, error) {
	fields := strings.Split(line, " ")
	arg, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}
	if op, ok := OpMap[fields[0]]; !ok {
		return nil, fmt.Errorf("Unknown operation: %q (line %q)", fields[0], line)
	} else {
		return &Instruction{op, arg}, nil
	}
}
