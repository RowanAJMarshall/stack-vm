package builtins

import (
	"fmt"
	"github.com/rowanajmarshall/stack-vm/utilities"
)

func SPrint(mem *utils.Stack) {
	val := SPeek(mem)
	fmt.Println(val)
}

// Pops the top value on the stack
func SPop(st *utils.Stack) int {
	l := len(*st)
	val := (*st)[l-1]
	*st = (*st)[:l-1]
	return val
}

// Pushes the given value onto the stack
func SPush(st *utils.Stack, v int) {
	*st = append(*st, v)
}

func SPeek(st *utils.Stack) int {
	return (*st)[len(*st)-1]
}

// Adds the top two integers and pushes the result
func SAdd(mem *utils.Stack) {
	num1 := SPop(mem)
	num2 := SPop(mem)
	SPush(mem, num1+num2)
}

//
func SDuplicate(mem *utils.Stack) {
	SPush(mem, SPeek(mem))
}

func SMultiply(mem *utils.Stack) {
	num1 := SPop(mem)
	num2 := SPop(mem)
	SPush(mem, num1*num2)
}

// Pops each integer off the stack and prints the ASCII char equivalent until a null character is encountered
func SPrintStr(mem *utils.Stack) {
	for {
		char := SPop(mem)
		if char != 0 {
			fmt.Print(string(char))
		} else {
			break
		}
	}
	fmt.Print("\n")
}
