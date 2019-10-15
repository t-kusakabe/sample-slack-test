package main

import (
	"context"
	"log"

	slackbot "github.com/lusis/go-slackbot"
	// slacktest "github.com/lusis/slack-test"
	slack "github.com/nlopes/slack"
	slacktest "github.com/t-kusakabe/slacktest"
)

func globalMessageHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "I see your message", slackbot.WithoutTyping)
}

func main() {
	s := slacktest.NewTestServer()
	s.SetBotName("MyBotName")

	slack.SLACK_API = "http://" + s.ServerAddr + "/"
	go s.Start()

	bot := slackbot.New("ABCEDFG")
	bot.Hear("this is a channel message").MessageHandler(globalMessageHandler)
	go bot.Run()

	s.SendMessageToChannel("#random", "this is a channel message")
	for m := range s.SeenFeed {
		log.Printf("saw message in slack: %s", m)
	}
}
