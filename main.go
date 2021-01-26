package main

import (
	"fmt"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type BorrowBot struct {
	bot reddit.Bot
}

func (r *BorrowBot) Post(p *reddit.Post) error {
	fmt.Println(fmt.Sprintf("Examining post title: %s", p.Title))
	if parsers.shouldConsider(p.Title) {
		fmt.Println(fmt.Sprintf("processed post with title %s at link %s", p.Title, p.Permalink))
	}
	return nil
}

func main() {
	bot, err := reddit.NewBotFromAgentFile("creds", 1*time.Minute)
	if err != nil {
		fmt.Println("Failed to create new bot from file", err)
		return
	}

	subredditsToWatch := []string{
		"borrow",
	}

	cfg := graw.Config{Subreddits: subredditsToWatch}
	handler := &BorrowBot{bot: bot}

	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}
