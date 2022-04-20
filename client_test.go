package vmmanager6

import (
	"crypto/tls"
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestClient_Login(t *testing.T) {
	*Debug = true
	client, err := NewClient("https://astra.ispsystem.net/vm/v3", nil, &tls.Config{InsecureSkipVerify: true}, 300)
	assert.Nil(t, err)
	err = client.Login("admin@example.com", "oA4hX9rL")
	assert.Nil(t, err)
	vm := &VmRef{vmId: 2}

	resp, err := client.GetVmState(vm)
	assert.Nil(t, err)
	log.Printf("%#v", resp)
}

