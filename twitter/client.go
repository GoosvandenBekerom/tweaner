package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// Tweet represents a tweet
type Tweet struct {
	anaconda.Tweet
}

// Client can be used to interact with the Twitter API
type Client interface {
	// GetTweet returns the tweet with the given id
	GetTweet(id int64) (Tweet, error)
	// GetTweets returns tweets of the authenticated user's timeline
	// 	- parameter 'max' specifies the maximum amount of tweets you want to receive
	GetTweets(max int) ([]Tweet, error)
	// DeleteTweet deletes the given tweet IF it was originally posted by the authenticated user
	// and returns the deleted tweet
	DeleteTweet(Tweet) (Tweet, error)
}

type client struct {
	api *anaconda.TwitterApi
}

func (c *client) GetTweet(id int64) (Tweet, error) {
	fmt.Printf("attempting to get tweet with id %d\n", id)
	tweet, err := c.api.GetTweet(id, url.Values{})
	return Tweet{tweet}, err
}

func (c *client) GetTweets(max int) (result []Tweet, e error) {
	fmt.Printf("attempting to get up to %d tweets\n", max)
	v := url.Values{}
	v.Set("count", fmt.Sprintf("%d", max))
	tweets, err := c.api.GetUserTimeline(v)
	if err != nil {
		e = err
		return
	}
	fmt.Printf("got %d tweets\n", len(tweets))
	for _, tweet := range tweets {
		result = append(result, Tweet{tweet})
	}
	return
}

func (c *client) DeleteTweet(tweet Tweet) (Tweet, error) {
	old, err := c.api.DeleteTweet(tweet.Id, false)
	return Tweet{old}, err
}

// NewClient constructs a new Client
func NewClient(secrets Secrets) Client {
	anaconda.SetConsumerKey(secrets.ConsumerKey)
	anaconda.SetConsumerSecret(secrets.ConsumerSecret)
	return &client{
		api: anaconda.NewTwitterApi(secrets.AccessToken, secrets.AccessTokenSecret),
	}
}
