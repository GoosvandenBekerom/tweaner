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

## Installing

Make sure you have the following binaries installed:

- `git` - let's be honest, what are you doing here if you dont have git installed..
- `go` - https://golang.org/

```
$ git clone https://github.com/GoosvandenBekerom/tweaner.git
$ cd tweaner
$ go install
```

## Usage

Make sure you have a [twitter developer account](https://developer.twitter.com/) linked to the account you want to clean up.  
Add an app to that account, get the following 4 secrets and set them as environment variables.

```
$ export TWEANER_CONSUMER_KEY=<your-twitter-consumer-key>
$ export TWEANER_CONSUMER_SECRET=<your-twitter-consumer-secret>
$ export TWEANER_ACCESS_TOKEN=<your-twitter-app-access-token>
$ export TWEANER_ACCESS_TOKEN_SECRET=<your-twitter-app-access-token-secret>

# run default (no backup support)
$ tweaner

# run tweaner for just 5 tweets
$ tweaner -n 5

# run tweaner for specific id
$ tweaner -id 123456789

# run tweaner in dryrun mode (no backups/deletions)
$ tweaner -d

# run with backup support
$ tweaner -b -p "/put/backup/path/here"

Usage of tweaner:
  -b    enables backup support, when enabled, tweaner creates a backup of the deleted tweets at the path specified with -p
  -d    dryrun, get tweets without deleting them
  -id int
        tweet id, when provided the amount given in -n will be ignored
  -n int
        max amount of tweets to delete (default 200)
  -p string
        root path for the backup files, required when backups are enabled with -b
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

I extensively used my twitter during high school/college (i.e. puberty), and I had like 25k tweets that are funny to read back,
but are in no way, shape or form what I want as the history of my "professional" twitter account.

Ofcourse there is no fun in manually cleaning up 25k tweets,
and since I'm in the process of learning Go at the time of writing this I figured, let's automate this.

And while I'm at it, why not open source it 🤷‍♂️.
