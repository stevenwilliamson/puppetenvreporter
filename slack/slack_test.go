package slack

import (
	"os"
	"testing"
)

var SlackAPIToken string

func init() {
	SlackAPIToken = os.Getenv("SLACK_API_TOKEN")
}

/*
  TestPostMessage
  Is an integration test, export a SLACK_API_TOKEN environment variable
  with a valid token for the tests
  TODO add flag for integration tests
  TODO convert to table test test invalid token
*/
func TestPostMessage(t *testing.T) {
	c := &SlackClient{Token: SlackAPIToken}
	err := c.PostMessage("This is a test", Channel("#general"), Username("testbot"))
	if err != nil {
		t.Errorf("Error posting to slack %v", err)
	}
}
