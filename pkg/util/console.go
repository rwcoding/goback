package util

import (
	"os"
	"strings"
)

type Console struct {
	Flags    map[string]string
	Commands map[string]bool
}

func NewConsole(args []string) *Console {
	console := &Console{
		Flags:    map[string]string{},
		Commands: map[string]bool{},
	}
	if args == nil {
		console.Parse(os.Args)
	} else {
		console.Parse(args)
	}
	return console
}

// Parse 初始化控制台参数
func (c *Console) Parse(args []string) {
	jump := 0
	length := len(args)
	for i := 0; i < length; i++ {
		if jump == i {
			continue
		}
		if strings.Contains(args[i], "=") {
			arr := strings.Split(args[i], "=")
			if strings.HasPrefix(arr[0], "--") {
				c.Flags[arr[0][2:]] = arr[1]
			} else if strings.HasPrefix(arr[0], "-") {
				c.Flags[arr[0][1:]] = arr[1]
			} else {
				c.Flags[arr[0]] = arr[1]
			}
			continue
		}
		if strings.HasPrefix(args[i], "-") {
			if strings.HasPrefix(args[i], "--") {
				c.Flags[args[i][2:]] = ""
			} else {
				c.Flags[args[i][1:]] = ""
			}
		}
	}
}

// HasCommand 是否存在指令
func (c *Console) HasCommand(name string) bool {
	_, ok := c.Commands[name]
	return ok
}

// HasFlag 是否存在参数
func (c *Console) HasFlag(name string) bool {
	_, ok := c.Flags[name]
	return ok
}

// GetFlag 参数值
func (c *Console) GetFlag(name string) string {
	return c.Flags[name]
}

// GetFlagAuto 参数值，如果没有取默认值
func (c *Console) GetFlagAuto(name string, defaultValue string) string {
	val, ok := c.Flags[name]
	if ok {
		return val
	}
	return defaultValue
}
