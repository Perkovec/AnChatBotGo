package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/perkovec/AnChatBotGo/commands"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

var CRegex map[string]*regexp.Regexp = map[string]*regexp.Regexp{}

var CRegexString = map[string]string{
	"some_command": "(?i)^/w*.*",
	"start":        "(?i)^(/start)$",
	"stop":         "(?i)^(/stop)$",
	"help":         "(?i)^(/help)$",
	"list":         "(?i)^(/list)$",
	"banlist":      "(?i)^(/banlist)$",
	"nick":         "(?i)^(/nick\\s)(.*)",            // 1 group = "/nick ", 2 group = nickname
	"kick":         "(?i)^(/kick\\s)(.*)",            // 1 group = "/kick ", 2 group = chat_id
	"ban":          "(?i)^(/ban\\s)(.*)",             // 1 group = "/ban ", 2 group = chat_id
	"unban":        "(?i)^(/unban\\s)(.*)",           // 1 group = "/unban ", 2 group = chat_id
	"rename":       "(?i)^(/rename)\\s(\\w*)\\s(.*)", // 1 group = "/rename ", 2 group = chat_id, 3 group = nick
	"id":           "(?i)^(/id)\\s(\\w*)\\s(\\w*)",   // 1 group = "/id ", 2 group = chat_id, 3 group = new chat_id
	"me":           "(?i)^(%)(.*)",                   // 1 group = "%", 2 group = text
	"me2":          "(?i)^(/me\\s)(.*)",
	"info":         "(?i)^(/info)(\\s(\\w*))?", // 1 group = "/info", 2 group = chat_id || undefined
	"clean":        "(?i)^(/clean)$",
}

func init() {
	for k, v := range CRegexString {
		compiled, err := regexp.Compile(v)
		if err != nil {
			fmt.Println("Error (compile regexp): ", err)
		}
		CRegex[k] = compiled
	}
}

func ProcessMsg(app commands.App, msg tgapi.Message) {
	text := strings.TrimSpace(msg.Text)

	if CRegex["start"].MatchString(text) {
		app.StartCommand(msg)
	} else if CRegex["stop"].MatchString(text) {
		app.StopCommand(msg)
	} else if CRegex["list"].MatchString(text) {
		app.ListCommand(msg)
	} else if CRegex["nick"].MatchString(text) {
		matches := CRegex["nick"].FindAllStringSubmatch(text, -1)
		app.NickCommand(msg, matches[0][2])
	} else if CRegex["help"].MatchString(text) {
		app.HelpCommand(msg)
	} else if CRegex["info"].MatchString(text) {
		matches := CRegex["info"].FindAllStringSubmatch(text, -1)
		var userid string
		if len(matches[0]) == 4 {
			userid = matches[0][3]
		}
		app.InfoCommand(msg, userid)
	} else if CRegex["me"].MatchString(text) {
		matches := CRegex["me"].FindAllStringSubmatch(text, -1)
		app.MeCommand(msg, matches[0][2])
	} else if CRegex["me2"].MatchString(text) {
		matches := CRegex["me2"].FindAllStringSubmatch(text, -1)
		app.MeCommand(msg, matches[0][2])
	} else {
		app.BroadcastMessage(msg)
	}
}
