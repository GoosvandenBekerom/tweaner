package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GoosvandenBekerom/tweaner/backup"
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

var (
	n             int
	dryrun        bool
	backupEnabled bool
	backupPath    string
)

func main() {
	fmt.Print(banner)

	flag.IntVar(&n, "n", 200, "max amount of tweets to delete")
	flag.BoolVar(&dryrun, "d", false, "dryrun, get tweets without deleting them")
	flag.BoolVar(&backupEnabled, "b", false, "enables backup support, when enabled, tweaner creates a backup of the deleted tweets at the path specified with -p")
	flag.StringVar(&backupPath, "p", "", "root path for the backup files, required when backups are enabled with -b")
	flag.Parse()

	if backupEnabled && backupPath == "" {
		fmt.Printf("\nincorrect usage: enabling backups requires a path to be set for the backup (set path with -p)\n\n")
		flag.Usage()
		os.Exit(1)
	}

	secrets := twitter.InitSecrets()
	client := twitter.NewClient(secrets)
	tweets, err := client.GetTweets(n)
	if err != nil {
		panic(fmt.Sprintf("unable to get tweets from authenticated user's timeline: %v", err))
	}

	for _, tweet := range tweets {
		if dryrun {
			fmt.Printf("[dryrun] got tweet with id: %d and content: %s\n", tweet.Id, tweet.Text)
			continue
		}
		if backupEnabled {
			if err := backup.Save(backupPath, tweet); err != nil {
				panic(fmt.Sprintf("unable to backup tweet with id %d: %v", tweet.Id, err))
			}
		}
		fmt.Printf("deleting tweet with id: %d and content: %s\n", tweet.Id, tweet.Text)
		old, err := client.DeleteTweet(tweet)
		if err != nil {
			panic(fmt.Sprintf("unable to delete tweet with id %d: %v", tweet.Id, err))
		}
		fmt.Printf("deleted tweet with id %d\n", old.Id)
	}
	return
}
