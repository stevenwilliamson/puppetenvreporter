package slack

import (
	"github.com/nlopes/slack"
)

// Option is a type alias for a function that takes a pointer
// to a SlackClient. We use this to expose options in a clean extendable way
// See https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md
type Option func(*SlackClient)

// SlackClient Represents a Slack client
// This package wraps a more complex slack client and provides a super simple interface
// for when we just want to send messages to slack
type SlackClient struct {
	Token    string
	api      *slack.Client
	username string
	channel  string
}

// PostMessage sends a message to a slack channel,
// Params:
// channel: Name of slack channel
// message: Message to send
//
// Returns:
// err: non nil if there was an error
func (s *SlackClient) PostMessage(message string, setters ...Option) error {

	// s.client() gets a client ready to use
	s.client()

	// Set any option state
	// Options set in one PostMessage persist for future calls
	// until you set them again
	for _, setter := range setters {
		setter(s)
	}

	msgParams := slack.NewPostMessageParameters()
	msgParams.Username = s.username
	_, _, err := s.client().PostMessage(s.channel, message, msgParams)
	return err
}

// Creates a client if we do not yet have one handy
// if we do we return the existing one
func (s *SlackClient) client() *slack.Client {
	if s.api == nil {
		api := slack.New(s.Token)
		s.api = api
		s.channel = "#test"
	}
	return s.api
}

// Username Sets the username that the message appears from
func Username(username string) Option {
	return func(args *SlackClient) {
		args.username = username
	}
}

// Channel sets the channel to post to
// we default to #test
func Channel(channel string) Option {
	return func(args *SlackClient) {
		args.channel = channel
	}
}
