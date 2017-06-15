package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/rowanajmarshall/stack-vm/utilities"
	"github.com/rowanajmarshall/stack-vm/stack-functions"
)

// Use this to run "go run !(*_test).go"

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
		if command == "push" {
			builtins.SPush(&mem, arg[0])
		} else if command == "pop" {
			builtins.SPop(&mem)
		} else if command == "print" {
			builtins.SPrint(&mem)
		} else if command == "add" {
			builtins.SAdd(&mem)
		} else if command == "ifeq" {
			prev = i-1
			if builtins.SPeek(&mem) == arg[0] {
				i = labels[arg[1]]
			}
		} else if command == "dup" {
			builtins.SDuplicate(&mem)
		} else if command == "return" {
			i = prev
		} else if command == "printstr" {
			builtins.SPrintStr(&mem, arg[0])
		} else if command == "end" {
			break
		}
	}
}

func LoadFile(s string) utils.Bytefile {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
