package builtins

import (
	"fmt"
	"github.com/rowanajmarshall/stack-vm/utilities"
)

func SPrint(mem *utils.Stack) {
	val := SPeek(mem)
	fmt.Println(val)
}

func SPop(st *utils.Stack) int {
	l := len(*st)
	val := (*st)[l-1]
	*st = (*st)[:l-1]
	return val
}

func SPush(st *utils.Stack, v int) {
	*st = append(*st, v)
}

func SPeek(st *utils.Stack) int {
	return (*st)[len(*st)-1]
}

func SAdd(mem *utils.Stack) {
	num1 := SPop(mem)
	num2 := SPop(mem)
	SPush(mem, num1+num2)
}

func SDuplicate(mem *utils.Stack) {
	SPush(mem, SPeek(mem))
}

func SPrintStr(mem *utils.Stack, length int) {
	for i:=0;i<length;i++ {
		fmt.Print(string(SPop(mem)))
	}
	fmt.Print("\n")
}
