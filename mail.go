package main

import (
	"strings"
	"net/http"
	"alerter/log"
	"io/ioutil"
	"fmt"
	"time"
)

// 邮件内容
const (
	AlertMailSubject = "服务器报警邮件"
	AlertMailContent = "服务器于#param2#发生#param3#级报警; 此警告已经连续触发#param4#次 \n堆栈信息:\n#param5#        \n\n---- 来自系统自动发送"
)

const (
	DailyServerMailLimit   = 300 // 服务器每天发送邮件次数
	DialyEchoUserMailLimit = 15  // 每位用户每日发送邮件次数
)

type LevelInfo struct {
	Name  string
	Times int32
}

var LevelList = make(map[int32]LevelInfo)

func init() {
	LevelList[1] = LevelInfo{Name: "轻度", Times: 1}
	LevelList[2] = LevelInfo{Name: "中度", Times: 5}
	LevelList[3] = LevelInfo{Name: "重度", Times: 10}
	LevelList[4] = LevelInfo{Name: "严重", Times: 100}
}

func AlertToUser(userMail string, stackInfo *StackInfo) (int, error) {

	url := config.App.MailServerAddr

	levelTitle := calcLevel(stackInfo.Times)
	mailContent := setMailContent(stackInfo.PanicTime, levelTitle, stackInfo.Times, stackInfo.GetStackLines(), stackInfo.Env)
	mailSubject := levelTitle + AlertMailSubject

	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"from\"\r\n\r\nExcited User <" + config.App.ServerAccount + ">\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"to\"\r\n\r\n" + userMail + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"subject\"\r\n\r\n" + mailSubject + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"text\"\r\n\r\n" + mailContent + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	req, _ := http.NewRequest("POST", url, payload)

	req.SetBasicAuth("api", config.App.ApiKey)
	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("Cache-Control", "no-cache")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Infof(string(body))

	return 0, nil
}

func calcLevel(level int32) string {
	var levelName string
	for _, levelInfo := range LevelList {
		if level <= levelInfo.Times {
			levelName = levelInfo.Name
			break
		}
	}
	return "[" + levelName + "]"
}

func setMailContent(panicTime int64, levelName string, times int32, stackinfo, env string) string {

	tm := time.Unix(panicTime, 0)
	formatTime := tm.Format("2006-01-02 15:04:05")

	var paramsMailContent []string
	paramsMailContent = append(paramsMailContent, env)
	paramsMailContent = append(paramsMailContent, formatTime)
	paramsMailContent = append(paramsMailContent, levelName)
	paramsMailContent = append(paramsMailContent, fmt.Sprintf("%d", times))
	paramsMailContent = append(paramsMailContent, stackinfo)

	var rawContent string
	rawContent = AlertMailContent
	rawContent = strings.Replace(rawContent, "#param1#", paramsMailContent[0], -1)
	rawContent = strings.Replace(rawContent, "#param2#", paramsMailContent[1], -1)
	rawContent = strings.Replace(rawContent, "#param3#", paramsMailContent[2], -1)
	rawContent = strings.Replace(rawContent, "#param4#", paramsMailContent[3], -1)
	rawContent = strings.Replace(rawContent, "#param5#", paramsMailContent[4], -1)

	return rawContent
}
