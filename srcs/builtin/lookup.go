package builtin

import (
	"fmt"
	"os"
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

package builtin

import (
	"fmt"
	"os"
	"os/exec"
)

// ExecBuiltin: ビルトインコマンドを実行する関数
func ExecBuiltin(name string, words []string, inputFile *os.File, outputFile *os.File) (exitStatus int) {
	// 現在の標準入力と標準出力を保存
	originalStdin := os.Stdin
	originalStdout := os.Stdout

	os.Stdin = inputFile
	os.Stdout = outputFile

	// ビルトインコマンドの実行
	switch name {
	case "pwd":
		exitStatus = pwd(words)
	case "cd":
		exitStatus = cd(words)
	case "exit":
		exitStatus = exit(words)
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown builtin command: %s\n", name)
		exitStatus = 1
	}

	// 標準入力と標準出力を元に戻す
	os.Stdin = originalStdin
	os.Stdout = originalStdout

	return
}
