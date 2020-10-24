package twitter

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const separator = `
===================================================================
`

// Secrets acts as a wrapper for all twitter related secrets
type Secrets struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// InitSecrets initializes all secrets from environment variables
// and return the Secrets wrapper
func InitSecrets() Secrets {
	fmt.Print(separator)
	fmt.Println("initializing secrets...")

	s := Secrets{
		ConsumerKey:       os.Getenv("TWEANER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWEANER_CONSUMER_SECRET"),
		AccessToken:       os.Getenv("TWEANER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWEANER_ACCESS_TOKEN_SECRET"),
	}

	// print length of secrets for debugging purposes
	fmt.Printf("consumer key: %s\n", mask(s.ConsumerKey))
	fmt.Printf("consumer secret: %s\n", mask(s.ConsumerSecret))
	fmt.Printf("access token: %s\n", mask(s.AccessToken))
	fmt.Printf("access token secret: %s", mask(s.AccessTokenSecret))
	fmt.Print(separator)

	return s
}

func mask(s string) string {
	return strings.Repeat("*", utf8.RuneCountInString(s))
}
