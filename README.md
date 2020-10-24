# Tweaner - Cleanup your Twitter history

Tweaner will delete the full history of your twitter account,
to essentially give it a clean slate.

```
___________
\__    ___/_  _  __ ____ _____    ____   ___________
  |    |  \ \/ \/ // __ \\__  \  /    \_/ __ \_  __ \
  |    |   \     /\  ___/ / __ \|   |  \  ___/|  | \/
  |____|    \/\_/  \___  >____  /___|  /\___  >__|
                       \/     \/     \/     \/
The most overkill way to delete the history of a twitter account ™
```

## Usage

```
$ git clone https://github.com/GoosvandenBekerom/tweaner.git
$ cd tweaner
$ go install
$ export TWEANER_CONSUMER_KEY=<your-twitter-consumer-key>
$ export TWEANER_CONSUMER_SECRET=<your-twitter-consumer-secret>
$ export TWEANER_ACCESS_TOKEN=<your-twitter-app-access-token>
$ export TWEANER_ACCESS_TOKEN_SECRET=<your-twitter-app-access-token-secret>
$ tweaner
```

## Development

What I do to setup environment variables for local development is create a file called `.env` (which is .gitignore'd):

```
export TWEANER_CONSUMER_KEY=<your-twitter-consumer-key>
export TWEANER_CONSUMER_SECRET=<your-twitter-consumer-secret>
export TWEANER_ACCESS_TOKEN=<your-twitter-app-access-token>
export TWEANER_ACCESS_TOKEN_SECRET=<your-twitter-app-access-token-secret>
```

And then run tweaner like this:

```
$ source .env && go run main.go
```

## Why?

I wanted to start using my twitter more professionally, as in following people in the industry,
maybe occasionally tweeting some tech related stuff etc.

I extensively used my twitter during puberty, and I had like 25k tweets that are funny to read back,
but are in no way, shape or form what I want as the history of my "professional" twitter account, right.

Ofcourse there is no fun in manually cleaning up 25k tweets,
and since I'm in the process of learning Go at the time of writing this I figured, let's automate this.
