package twitter

import (
	"errors"
	"fmt"
	"time"

	"github.com/dapr/components-contrib/bindings"
	"github.com/mrjones/oauth"
)

const (
	consumerKey    = "consumerKey"
	consumerSecret = "consumerSecret"
	accessToken    = "accessToken"
	tokenSecret    = "tokenSecret"
	status         = "status"
	twitterURLBase = "https://api.twitter.com/1.1/statuses/update.json"
)

//OutputBinding Default comment
type OutputBinding interface {
	Init(metadata bindings.Metadata) error
	Write(req *bindings.WriteRequest) error
}

//Tweet Default comment
type Tweet struct {
	metadata twitterMetadata
}

type twitterMetadata struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	tokenSecret    string
	status         string
}

//NewTweet Default comment
func NewTweet() *Tweet {
	return &Tweet{}
}

//Init Default comment
func (t *Tweet) Init(metadata bindings.Metadata) error {
	twitterTweet := twitterMetadata{
		status: "default post",
	}

	if metadata.Properties["consumerKey"] == "" {
		return errors.New("\"consumerKey\" is a required field")
	}
	if metadata.Properties["consumerSecret"] == "" {
		return errors.New("\"consumerSecret\" is a required field")
	}
	if metadata.Properties["accessToken"] == "" {
		return errors.New("\"accessToken\" is a required field")
	}
	if metadata.Properties["tokenSecret"] == "" {
		return errors.New("\"tokenSecret\" is a required field")
	}
	if metadata.Properties["status"] == "" {
		return errors.New("\"status\" is a required field")
	}

	twitterTweet.consumerKey = metadata.Properties["consumerKey"]
	twitterTweet.consumerSecret = metadata.Properties["consumerSecret"]
	twitterTweet.accessToken = metadata.Properties["accessToken"]
	twitterTweet.tokenSecret = metadata.Properties["tokenSecret"]
	twitterTweet.status = metadata.Properties["status"]

	return nil
}

func (t *Tweet) Write(req *bindings.WriteRequest) error {
	fmt.Println("Writing Test Started.")
	consumer := oauth.NewConsumer(t.metadata.consumerKey, t.metadata.consumerSecret, oauth.ServiceProvider{})
	accessToken := &oauth.AccessToken{Token: t.metadata.accessToken,
		Secret: t.metadata.tokenSecret}
	currentTime := time.Now()
	response, err := consumer.Post(twitterURLBase, map[string]string{"consumerKey": t.metadata.consumerKey, "consumerSecret": t.metadata.consumerSecret,
		"accessToken": t.metadata.consumerSecret, "tokenSecret": t.metadata.tokenSecret,
		"status": "dapr status " + currentTime.Format("2006-01-02 15:04:05")}, accessToken)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	fmt.Println("Response:", response.StatusCode, response.Status)
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		return fmt.Errorf("error from Twitter: %s", response.Status)
	}

	return nil
}
