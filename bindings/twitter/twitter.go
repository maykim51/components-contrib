package twitter

import (
	"errors"
	"fmt"

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

type OutputBinding interface {
	Init(metadata bindings.Metadata) error
	Write(req *bindings.WriteRequest) error
}

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

func NewTweet() *Tweet {
	return &Tweet{}
}

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

	t.metadata = twitterTweet

	return nil
}

func (t *Tweet) Write(req *bindings.WriteRequest) error {
	fmt.Println("Writing Test Started.")
	consumer := oauth.NewConsumer(req.Metadata["consumerKey"], req.Metadata["consumerSecret"], oauth.ServiceProvider{})
	accessToken := &oauth.AccessToken{Token: req.Metadata["accessToken"], Secret: req.Metadata["tokenSecret"]}
	response, err := consumer.Post(twitterURLBase, req.Metadata, accessToken)

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
