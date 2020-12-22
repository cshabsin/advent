package console

import (
	"fmt"
	"strconv"
	"strings"
)

type Console struct {
	Instructions []*Instruction
	Accumulator  int
	IP           int

	Executed map[int]bool
}

func New() *Console {
	return &Console{Executed: map[int]bool{}}
}

func (cons *Console) ReadInstruction(line string) error {
	i, err := ReadInstruction(line)
	if err != nil {
		return err
	}
	cons.Instructions = append(cons.Instructions, i)
	return nil
}

func (cons *Console) Run() int {
	for {
		if cons.Executed[cons.IP] {
			return cons.Accumulator
		}
		cons.Executed[cons.IP] = true
		instruction := cons.Instructions[cons.IP]
		switch instruction.Operation {
		case ACC:
			cons.Accumulator += instruction.Argument
			cons.IP++
		case JMP:
			cons.IP += instruction.Argument
		case NOP:
			cons.IP++
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
