package repl

import (
	"fmt"
	"os"
	"shell/executor"
	"strings"

	"github.com/chzyer/readline"
)

const PS1 = "🐠$ "

func Start() {
	rl, err := readline.New(PS1)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				// Ctrl-C が押されたら次のプロンプトを表示する
				continue
			}

			fmt.Println("exit: ", err)
			return
		}

		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		// コマンドの実行
		executor.ExecSimpleCommandSync(args, os.Stdin, os.Stdout)
	}
}
