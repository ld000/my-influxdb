package run

import (
	"fmt"
)

// Command "influxd run" 结构体
type Command struct {
}

// NewCommand 返回个新的 Command
func NewCommand() *Command {
	return &Command{}
}

// Run 处理参数，运行 server
func (cmd *Command) Run(args ...string) error {
	fmt.Println("command.go run")
	return nil
}

// Close 结束 server
func (cmd *Command) Close() error {
	fmt.Println("command.go close")
	return nil
}
