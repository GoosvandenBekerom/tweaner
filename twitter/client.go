package twitter

import "github.com/ChimeraCoder/anaconda"

// Client can be used to interact with the Twitter API
type Client struct {
	*anaconda.TwitterApi
}

// NewClient constructs a new Client
func NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Client {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return &Client{
		anaconda.NewTwitterApi(accessToken, accessTokenSecret),
	}
}
