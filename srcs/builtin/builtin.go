package builtin

import (
	"fmt"
	"os"
	"strconv"
)

var builtins = map[string]func([]string) int{
	"pwd":  pwd,
	"cd":   cd,
	"exit": exit,
}

func LookupBuiltin(name string) bool {
	_, exists := builtins[name]
	return exists
}

func ExecBuiltin(name string, words []string, inputFile *os.File, outputFile *os.File) (exitStatus int) {
	// 現在の標準入力と標準出力を保存
	originalStdin := os.Stdin
	originalStdout := os.Stdout

	os.Stdin = inputFile
	os.Stdout = outputFile

	exitStatus = builtins[name](words)

	// 標準入力と標準出力を元に戻す
	os.Stdin = originalStdin
	os.Stdout = originalStdout

	return
}

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
		os.Exit(255)
	}
	os.Exit(status)
	return 0
}
