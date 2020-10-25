package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/GoosvandenBekerom/tweaner/twitter"
)

func main() {
	id := flag.Int64("id", 0, "id of tweet to download")
	flag.Parse()

	if *id == 0 {
		flag.Usage()
		os.Exit(1)
	}

	secrets := twitter.InitSecrets()
	client := twitter.NewClient(secrets)

	tweet, err := client.GetTweet(*id)
	if err != nil {
		log.Fatalf("unable to get tweet: %v", err)
	}
	f, err := os.Create(fmt.Sprintf("%d.json", tweet.Id))
	if err != nil {
		log.Fatalf("unable to get tweet: %v", err)
	}
	defer f.Close()
	if err = json.NewEncoder(f).Encode(tweet); err != nil {
		log.Fatalf("failed to json encode tweet: %v", err)
	}
	log.Printf("saved tweet to file %s\n", f.Name())
}
