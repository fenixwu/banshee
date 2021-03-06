# Banshee

[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://godoc.org/github.com/fenixwu/banshee)

Publish messages in slack channel, support fuzzy selecting channel.

## Regist a channel

```go
channelName := "log"
channelWebhook := "some slack webhook URL"
banshee.RegistChannel(channelName, channelWebhook)
```

## Publish a message

```go
log := banshee.New("SEARCH_PATTERN_FOR_CHANNEL_NAME")
err := log.SetMessage("test message").ExactPublish()

if err != nil {
  // error handle
}
```

## Publish Mode

### ExactPublish

Pattern "A" will make publishing messages in channel "A" exactly.

### FuzzyPublish

Pattern "ab" will make publishing messages in channel ".*[aA][bB].*", for example, "cab", "abc", "Ab", "aB", etc.
