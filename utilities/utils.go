package utils

import (
	"strconv"
	"strings"
)

type Stack []int
type Bytefile []string
type FrameStack struct {
	mem Stack
}

func GetCommand(s string) (string, []int) {
	command := strings.Split(s, ":")
	arg := []int{}
	if len(command) == 1 {
		return command[0], []int{}
	} else {
		for i := 1; i < len(command); i++ {
			newArg, err := strconv.Atoi(command[i])
			if err != nil {
				panic(err)
			}
			arg = append(arg, newArg)
		}
		return command[0], arg
	}
}

func ExtractAssignment(s string) {

}
