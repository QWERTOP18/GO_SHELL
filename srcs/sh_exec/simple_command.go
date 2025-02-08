package sh_exec

import (
	"fmt"
	"os"
	"os/exec"
)

// func ExecSimpleCommand(words []string, rdIn []string, rdOut []string) error {
func ExecSimpleCommand(words []string) error {
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	/* pipelineは最後の一つ以外非同期的に処理 */
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v: %v\n", words[0], err)
	}

	fmt.Println("Command finished")
	return err
}

func ExecSimpleCommandRedirection(words []string, inputFile string, outputFile string) error {
	cmd := exec.Command(words[0], words[1:]...)

	if inputFile != "" {
		inFile, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:%v: %v\n", inputFile, err)
			return err
		}
		defer inFile.Close()
		cmd.Stdin = inFile
	} else {
		cmd.Stdin = os.Stdin
	}
	if outputFile != "" {
		outFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:%v: %v\n", outputFile, err)
			return err
		}
		defer outFile.Close()
		cmd.Stdout = outFile
	} else {
		cmd.Stdout = os.Stdout
	}

	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
	}

	fmt.Println("Command finished")
	return err
}
