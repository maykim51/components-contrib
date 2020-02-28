package twitter

import (
	"testing"
	"time"

	"github.com/dapr/components-contrib/bindings"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	m := bindings.Metadata{}
	currentTime := time.Now()
	m.Properties = map[string]string{
		"consumerKey":    "consumerKey",
		"consumerSecret": "consumerSecret",
		"accessToken":    "accessToken",
		"tokenSecret":    "toeknSecret",
		//Twitter API does not allow same status update -> attached current time to test status
		"status": "dapr status " + currentTime.Format("2006-01-02 15:04:05")}
	tw := NewTweet()
	err := tw.Init(m)
	assert.Nil(t, err)
}

func TestWrite(t *testing.T) {
	tw := NewTweet()
	r := bindings.WriteRequest{}
	currentTime := time.Now()
	r.Metadata = map[string]string{
		"consumerKey":    "consumerKey",
		"consumerSecret": "consumerSecret",
		"accessToken":    "accessToken",
		"tokenSecret":    "toeknSecret",
		//Twitter API does not allow same status update -> attached current time to test status
		"status": "dapr status " + currentTime.Format("2006-01-02 15:04:05")}
	err := tw.Write(&r)
	assert.Nil(t, err)
}
