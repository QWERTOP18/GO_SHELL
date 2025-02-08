package sh_exec

import (
	"fmt"
	"testing"
)

func TestExecSimpleCommand(t *testing.T) {
	tests := []struct {
		words []string
	}{
		{
			words: []string{"ls", "-l", "--color=auto"},
		},
		{
			words: []string{"echo", "Hello🐭"},
		},
		{
			words: []string{"non_existent_command"},
		},
	}

	for _, tt := range tests {
		fmt.Println("\n=== Executing command ----------------")
		ExecSimpleCommand(tt.words)
	}
}

func TestExecSimpleCommandRedirection(t *testing.T) {
	words := []string{"grep", "func"}
	inputFile := "simple_command_test.go"
	outputFile := "output.txt"
	ExecSimpleCommandRedirection(words, inputFile, outputFile)
}
