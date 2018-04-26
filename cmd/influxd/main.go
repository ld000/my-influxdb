// influxd
package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/ld000/my-influxdb/cmd/influxd/run"
)

func main() {
	fmt.Println("hello world")

	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Main 主结构体
type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// NewMain 返回 Main 实例
func NewMain() *Main {
	return &Main{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// Run 运行 CLI 命令参数
func (m *Main) Run(args ...string) error {
	// TODO args 参数处理

	// switch args[0] {
	// case "", "run":
	cmd := run.NewCommand()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run: %s", err)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	<-signalCh
	go cmd.Close()
	// default:
	// 	return fmt.Errorf(`unknown command "%s"`+"\n"+`Run 'influxd help' for usage`+"\n\n", args[0])
	// }

	return nil
}
