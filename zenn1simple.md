Go言語を初めて触るという人もいるかもしれないので


```go :simple_command/simple_command.go
package simple_command

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecSimpleCommand(words []string) error {
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	/* cmdを同期的に処理 */
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%s: %v\n", words[0], err)
	}

	fmt.Println("Command finished")
	return　err
}
```

goでmannualを見るには```go doc```を用いることができる。例えば```exec.Command```はコマンドのnameと引数arg(...は可変超配列である)を受け取り

::::details ```$go doc exec.Command```
```text

package exec // import "os/exec"

func Command(name string, arg ...string) *Cmd
    Command returns the Cmd struct to execute the named program with the given
    arguments.

    It sets only the Path and Args in the returned structure.

    If name contains no path separators, Command uses LookPath to resolve name
    to a complete path if possible. Otherwise it uses name directly as Path.

    The returned Cmd's Args field is constructed from the command name followed
    by the elements of arg, so arg should not include the command name itself.
    For example, Command("echo", "hello"). Args[0] is always name, not the
    possibly resolved Path.
```
::::

```go :main.go
package main

import (
	"shell/simple_command"
)

func main() {
	simple_command.ExecSimpleCommand([]string{"ls", "-l"})
}
```

```sh
shell
│   ├── go.mod
│   ├── main.go
│   ├── simple_command
│   │   └── simple_command.go
```

ディレクトリ階層はこんな感じにするといいかもしれない。```go.mod```はmainがある階層で
```sh
go mod init shell
```
をすると作成される。この引数は何でも良いが```build```した時のプログラム名になる。実行する方法は
```sh
go build
./shell
```
もしくは
```sh
go run .
```
で実行することができる。