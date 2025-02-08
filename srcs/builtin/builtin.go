package builtin

import (
	"fmt"
	"os"
	"strconv"
)

func pwd(words []string) int {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: pwd:%v\n", err)
		return 1
	}
	fmt.Println(cwd)
	return 0
}

func cd(words []string) int {
	if len(words) != 2 {
		fmt.Fprintf(os.Stderr, "Error: cd: too many arguments\n")
		return 1
	}
	err := os.Chdir(words[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cd: %v\n", err)
		return 1
	}
	return 0
}

func exit(words []string) int {
	if len(words) != 2 {
		fmt.Fprintf(os.Stderr, "Error: exit: too many arguments\n")
		return 1
	}
	status, err := strconv.Atoi(words[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: exit: %s: inumeric argument required\n", words[1])
		return 1
	}
	os.Exit(status)
}
