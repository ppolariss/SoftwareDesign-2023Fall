package command

import (
	"fmt"
	e "design/myError"
	"strconv"
	"strings"
)

type delete struct {
	line_num int
	content  string
}

func (c *delete) Execute() error {
	fmt.Println("delete")
	return nil
	// 删除指定标题或⽂本。如果指定⾏号，则删除指定⾏。当删除的是标题时，其⼦标题
	// 和内容不会被删除。
}

func (c *delete) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("delete: args error")
	} else if len(args) == 2 {
		c.line_num = -1
		c.content = args[1]
	} else {
		num, err := strconv.Atoi(args[1])
		// The line number is not specified
		if err != nil {
			c.line_num = -1
			slice_args := args[1:]
			c.content = strings.Join(slice_args, " ")
		}
		// The line number is specified
		// 要得到一个最大整数
		if num < 0 || num > 10000 {
			return e.NewMyError("delete: line number error")
		}
		c.line_num = num
		slice_args := args[2:]
		c.content = strings.Join(slice_args, " ")
	}
	return nil
}
