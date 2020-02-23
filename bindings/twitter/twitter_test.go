package twitter

import (
	"testing"

	"github.com/dapr/components-contrib/bindings"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	m := bindings.Metadata{}
	m.Properties = map[string]string{"status": "status"}
	tw := NewTweet()
	err := tw.Init(m)
	assert.Nil(t, err)
}

func TestWrite(t *testing.T) {
	tw := NewTweet()
	r := bindings.WriteRequest{}
	err := tw.Write(&r)
	assert.Nil(t, err)
}
