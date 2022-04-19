package vmmanager6

import (
	"crypto/tls"
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestClient_Login(t *testing.T) {
	*Debug = true
	client, err := NewClient("https://astra.ispsystem.net/vm/v3", nil, &tls.Config{InsecureSkipVerify: true}, 300)
	assert.Nil(t, err)
	err = client.Login("admin@example.com", "oA4hX9rL")
	assert.Nil(t, err)
	resp, err := client.GetVmList()
	log.Printf("%v", resp)
	log.Printf("AAAAAAAAAAAAAAAAAAAAAAAAAAA %v", resp)
	assert.Nil(t, err)
	//vms := resp["data"].([]interface{})
	//log.Printf("%v", vms)
	//for vmii := range vms {
	//	log.Printf("%v", vmii)
	//}
}

