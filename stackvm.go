package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rowanajmarshall/stack-vm/stack-functions"
	"github.com/rowanajmarshall/stack-vm/utilities"
)

// Use this to run "go run !(*_test).go"

// Enums representing various instructions
const (
	PUSH     = "push"
	POP      = "pop"
	MULTI    = "multi"
	PRINT    = "print"
	ADD      = "add"
	IFEQ     = "ifeq"
	DUP      = "dup"
	RETURN   = "return"
	PRINTSTR = "printstr"
	END      = "end"
)

func main() {
	progArgs := os.Args[1:]
	mem := utils.Stack{}
	file := LoadFile(progArgs[0])
	labels := make(map[int]int)
	prev := 0

	for k := 0; k < len(file); k++ {
		label, num := utils.GetCommand(file[k])
		if label == "label" {
			// Minus 1 to compensate for zero-based arrays
			labels[num[0]] = k - 1
		} else {
			continue
		}
	}

	for i := 0; i < len(file); i++ {
		if i > len(file)-1 {
			panic(fmt.Sprintf("Line number out of bounds: %d\n", i))
		}
		command, arg := utils.GetCommand(file[i])
		if strings.HasPrefix(command, "//") || command == "label" {
			continue
		}

		switch command {
		case PUSH:
			builtins.SPush(&mem, arg[0])
		case POP:
			builtins.SPop(&mem)
		case MULTI:
			builtins.SMultiply(&mem)
		case PRINT:
			builtins.SPrint(&mem)
		case ADD:
			builtins.SAdd(&mem)
		case IFEQ:
			prev = i - 1
			if builtins.SPeek(&mem) == arg[0] {
				i = labels[arg[1]]
			}
		case DUP:
			builtins.SDuplicate(&mem)
		case RETURN:
			i = prev
		case PRINTSTR:
			builtins.SPrintStr(&mem)
		case END:
			break
		}
	}
}

// LoadFile Loads a given bytecode file with the URL s
func LoadFile(s string) utils.Bytefile {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
