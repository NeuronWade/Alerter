package main

import (
	"strings"
	"github.com/labstack/echo"
	"alerter/obtime"
)

func StaticsAlert(c echo.Context) error {
	info := c.FormValue("stack_info")
	env := c.FormValue("env_info")
	var arrStr []string

	var stackHead string
	stackInfo := &StackInfo{}

	stackInfo.Env = env

	arrStr = strings.Split(info, "\n")
	for _, str := range arrStr {
		if strings.HasPrefix(str, "runtime/debug.Stack") {
			stackHead = str
		}
		stackInfo.StackLines = append(stackInfo.StackLines, str)
	}

	serverStack.Lock()

	if serverStack.StackList[stackHead] == nil {
		serverStack.StackList[stackHead] = &StackInfo{}
		serverStack.StackList[stackHead].PanicTime = obtime.Now().Unix()
	}

	serverStack.StackList[stackHead].StackLines = stackInfo.StackLines
	serverStack.StackList[stackHead].UpdateTime = obtime.Now().Unix()
	if serverStack.StackList[stackHead].Times < 100 {
		serverStack.StackList[stackHead].Times++
	}

	stackTimes := serverStack.StackList[stackHead].Times

	if config.App.MailNotify == "true" {
		if stackTimes == LevelList[1].Times || stackTimes == LevelList[2].Times ||
			stackTimes == LevelList[3].Times || stackTimes == LevelList[4].Times {
			for _, mailAddr := range config.Users.MailList {
				go AlertToUser(mailAddr, serverStack.StackList[stackHead])
			}
		}
	}

	serverStack.Unlock()
	return nil
}
