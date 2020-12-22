package console

import (
	"fmt"
	"strconv"
	"strings"
)

type Console struct {
	Instructions []*Instruction
	Accumulator  int
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

func (cons *Console) Run(ip int, executed map[int]bool) int {
	if executed == nil {
		executed = map[int]bool{}
	}
	for {
		if executed[ip] {
			return cons.Accumulator
		}
		executed[ip] = true
		instruction := cons.Instructions[ip]
		switch instruction.Operation {
		case ACC:
			cons.Accumulator += instruction.Argument
			ip++
		case JMP:
			if ip+1 == len(cons.Instructions) {
				return cons.Accumulator
			}
			ip += instruction.Argument
		case NOP:
			if ip+instruction.Argument == len(cons.Instructions) {
				return cons.Accumulator
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
