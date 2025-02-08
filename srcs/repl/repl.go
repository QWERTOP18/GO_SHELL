package repl

import (
	"fmt"
	"github.com/chzyer/readline"
	"shell/sh_exec"
	"strings"
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
			fmt.Println("Error reading input:", err)
			return
		}
		line = strings.TrimSpace(line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		sh_exec.ExecSimpleCommand(args)
	}
}
