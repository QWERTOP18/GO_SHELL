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

func ExecBuiltin(name string, words []string, inputFile *os.File, outputFile *os.File) (exitStatus int) {
	// 現在の標準入力と標準出力を保存
	originalStdin := os.Stdin
	originalStdout := os.Stdout

	os.Stdin = inputFile
	os.Stdout = outputFile

	// ビルトインコマンドの実行
	exitStatus = builtins[name](words)

	// 標準入力と標準出力を元に戻す
	os.Stdin = originalStdin
	os.Stdout = originalStdout

	return
}
