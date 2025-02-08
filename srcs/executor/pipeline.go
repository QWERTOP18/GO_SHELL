package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec1pipe(words1, words2 []string) error {
	cmd1 := exec.Command(words1[0], words1[1:]...)
	cmd2 := exec.Command(words2[0], words2[1:]...)

	stdoutPipe, err := cmd1.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// 標準入出力とエラー出力を設定
	cmd1.Stdin = os.Stdin
	cmd1.Stderr = os.Stderr
	cmd2.Stdin = stdoutPipe
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr

	if err := cmd1.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v: %v\n", words1[0], err)
	}
	if err := cmd2.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v: %v\n", words2[0], err)
	}

	cmd1.Wait()
	err = cmd2.Wait()
	fmt.Println("Pipe finished")
	return err
}

func ExecPipeline(lists [][]string) {
	// listsの長さに基づいてexecListsを初期化
	execLists := make([]*exec.Cmd, len(lists))

	// コマンドを実行可能なexec.Cmdに変換してexecListsにセット
	for i, list := range lists {
		execLists[i] = exec.Command(list[0], list[1:]...)
	}

	// パイプラインをつなげる
	for i := 0; i < len(execLists)-1; i++ {
		// execLists[i] の標準出力を execLists[i+1] の標準入力に接続
		stdoutPipe, err := execLists[i].StdoutPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error setting up pipe: %v\n", err)
			return
		}
		execLists[i+1].Stdin = stdoutPipe
	}
	//最初のコマンドと最後のコマンドの接続先
	execLists[0].Stdin = os.Stdin
	execLists[len(execLists)-1].Stdout = os.Stdout

	// すべてのコマンドを実行
	for _, execCommand := range execLists {
		if err := execCommand.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "command %s: %v\n", execCommand.Args[0], err)
			return
		}
	}

	// すべてのコマンドが終了するまで待機
	for _, execCommand := range execLists {
		if err := execCommand.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "command %s: %v\n", execCommand.Args[0], err)
		}
	}
}

// func main() {
// 	// 3つのコマンドをパイプラインで実行
// 	ExecPipeline([][]string{
// 		{"sleep", "5"},
// 		{"ls"},
// 		{"grep","a"},
// 		//{"cat"},
// 	})
// }

// func main() {
// 	// 1つ目のコマンド：sleep 5
// 	sleepCmd := exec.Command("sleep", "5")

// 	// 2つ目のコマンド：ls
// 	lsCmd := exec.Command("ls")

// 	// 3つ目のコマンド：cat
// 	catCmd := exec.Command("cat")

// 	// パイプラインのつなげ方：sleep -> ls -> cat
// 	// sleepCmdの出力をlsCmdの入力に接続
// 	lsCmd.Stdin, _ = sleepCmd.StdoutPipe()
// 	// lsCmdの出力をcatCmdの入力に接続
// 	catCmd.Stdin, _ = lsCmd.StdoutPipe()

// 	// catCmdの標準出力を直接os.Stdoutに設定して、その場で出力
// 	catCmd.Stdout = os.Stdout

// 	// コマンドを順番に非同期実行
// 	if err := sleepCmd.Start(); err != nil {
// 		fmt.Fprintf(os.Stderr, "sleepCmd: %v\n", err)
// 	}
// 	if err := lsCmd.Start(); err != nil {
// 		fmt.Fprintf(os.Stderr, "lsCmd: %v\n", err)
// 	}
// 	if err := catCmd.Start(); err != nil {
// 		fmt.Fprintf(os.Stderr, "catCmd: %v\n", err)
// 	}

// 	// 全てのコマンドが終了するのを待つ
// 	sleepCmd.Wait();
// 	lsCmd.Wait();
// 	if err := catCmd.Wait(); err != nil {

// 	}
// }
