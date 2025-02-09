package executor

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// func ExecSimpleCommand(words []string) int {
// 	cmd := exec.Command(words[0], words[1:]...)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

//		err := cmd.Run()
//		if err != nil {
//			if exitErr, isExitError := err.(*exec.ExitError); isExitError {
//				// プロセスの終了コードを取得
//				if waitStatus, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
//					return waitStatus.ExitStatus()
//				}
//			}
//			fmt.Fprintf(os.Stderr, "Error: %v: %v\n", words[0], "command not found")
//			return 127
//		}
//		return 0
//	}
func ExecSimpleCommandSync(words []string, inputFile *os.File, outputFile *os.File) int {
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdin = inputFile
	cmd.Stdout = outputFile
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exitErr, isExitError := err.(*exec.ExitError); isExitError {
			// プロセスの終了コードを取得
			if waitStatus, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
				return waitStatus.ExitStatus()
			}
		}
		fmt.Fprintf(os.Stderr, "Error: %v: %v\n", words[0], "command not found")
		return 127
	}
	return 0
}

/*
signalを受け取るためにはSimple Commandも非同期的に処理しないといけない
builtinの処理も
*/

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
