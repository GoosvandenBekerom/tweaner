package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/GoosvandenBekerom/tweaner/twitter"
)

const banner = `
___________                                          
\__    ___/_  _  __ ____ _____    ____   ___________ 
  |    |  \ \/ \/ // __ \\__  \  /    \_/ __ \_  __ \
  |    |   \     /\  ___/ / __ \|   |  \  ___/|  | \/
  |____|    \/\_/  \___  >____  /___|  /\___  >__|   
                       \/     \/     \/     \/       
The most overkill way to delete the history of a twitter account ™
`
const separator = `
===================================================================
`

func main() {
	fmt.Print(banner)
	fmt.Print(separator)

	consumerKey := os.Getenv("TWEANER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWEANER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWEANER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWEANER_ACCESS_TOKEN_SECRET")

	// print length of secrets for debugging purposes
	fmt.Printf("consumer key: %s\n", mask(consumerKey))
	fmt.Printf("consumer secret: %s\n", mask(consumerSecret))
	fmt.Printf("access token: %s\n", mask(accessToken))
	fmt.Printf("access token secret: %s\n", mask(accessTokenSecret))
	fmt.Print(separator)

	client := twitter.NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	tweets, err := client.GetUserTimeline(url.Values{
		"count": []string{"10"},
	})
	if err != nil {
		log.Fatalf("unable to get tweets from user timeline: %v", err)
	}

	for _, tweet := range tweets {
		fmt.Println(tweet.Text)
	}
}

func mask(s string) string {
	return strings.Repeat("*", utf8.RuneCountInString(s))
}
