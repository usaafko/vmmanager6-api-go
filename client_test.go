package vmmanager6

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Login(t *testing.T) {
	*Debug = true
	client, err := NewClient("https://astra.ispsystem.net/vm/v3", nil, &tls.Config{InsecureSkipVerify: true}, 300)
	assert.Nil(t, err)
	err = client.Login("admin@example.com", "fuck")
	assert.Nil(t, err)
}

