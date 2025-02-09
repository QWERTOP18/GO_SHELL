package executor

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// func ExecSimpleCommandSync(words []string, inputFile *os.File, outputFile *os.File) int {
// 	cmd := exec.Command(words[0], words[1:]...)
// 	cmd.Stdin = inputFile
// 	cmd.Stdout = outputFile
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		if exitErr, isExitError := err.(*exec.ExitError); isExitError {
// 			// プロセスの終了コードを取得
// 			if waitStatus, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
// 				return waitStatus.ExitStatus()
// 			}
// 		}
// 		fmt.Fprintf(os.Stderr, "Error: %v: %v\n", words[0], "command not found")
// 		return 127
// 	}
// 	return 0
// }

func setupSignal(cmd *exec.Cmd) {
	// シグナルを受け取るチャネルを作成
	sigchan := make(chan os.Signal, 1)

	// 親プロセスがSIGINTを受け取ったときに処理する
	signal.Notify(sigchan, syscall.SIGINT)

	go func() {
		for range sigchan {
			// 親プロセスがSIGINTを受け取った場合、子プロセスにそのシグナルを送る
			//fmt.Printf("Parent received signal: %v, forwarding to child...\n", sig)
			// 子プロセスにSIGINTを送る
			if cmd.Process != nil {
				cmd.Process.Signal(syscall.SIGINT)
			}
		}
	}()
}

func ExecSimpleCommandSync(words []string, inputFile *os.File, outputFile *os.File) int {
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdin = inputFile
	cmd.Stdout = outputFile
	cmd.Stderr = os.Stderr

	//backgroundでは必要!!!
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Setsid: true,
	// }
	setupSignal(cmd)

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
