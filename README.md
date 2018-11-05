# Banshee

[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://godoc.org/github.com/fenixwu/banshee)

Publish messages in slack channel, support fuzzy selecting channels.

## Register a channel

```go
banshee.RegistChannel("channelName", "some slack webhook URL")
```

## Publish a message

```go
logger := banshee.New("SEARCH_PATTERN_FOR_CHANNEL_NAME")

if err := logger.SetMessage("test message").ExactPublish(); err != nil {
  // error handle
}
```

## Publish Mode

### ExactPublish

Pattern "A" will make publishing messages in channel "A" exactly.

### FuzzyPublish

Pattern "ab" will make publishing messages in channel ".*[aA][bB].*", for example, "cab", "abc", "Ab", "aB", etc.
