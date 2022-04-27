package vmmanager6

import (
	"crypto/tls"
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
//	"time"
//	"fmt"
)
//func TestCreateVM(t *testing.T) {
//	*Debug = true
//	client, err := NewClient("https://astra.ispsystem.net", nil, &tls.Config{InsecureSkipVerify: true}, 300)
//	assert.Nil(t, err)
//	err = client.Login("admin@example.com", "oA4hX9rL")
//	assert.Nil(t, err)
//	config := ConfigNewQemu{
//                Name:         "some",
//                Description:  "",
//                Memory:       512,
//                QemuCores:    1,
//                QemuDisks:    6000,
//		Cluster: 1,
//		Account: 3,
//		Domain: "some.fuck",
//		Os: 25,
//		IPv4: 1,
//		Password: "fuck@Me123",
//        }
//	vmid, err := config.CreateVm(client)
//	assert.Nil(t, err)
//	log.Printf(">>> VMID:  %v", fmt.Sprint(vmid))
//
//
//}
//
//func TestGetVmInfo(t *testing.T) {
//	*Debug = true
//	client, err := NewClient("https://astra.ispsystem.net", nil, &tls.Config{InsecureSkipVerify: true}, 300)
//	assert.Nil(t, err)
//	err = client.Login("admin@example.com", "oA4hX9rL")
//	assert.Nil(t, err)
//	vmr := &VmRef{vmId: 41}
//	config, err := NewConfigQemuFromApi(vmr, client)
//	log.Printf(">>> VM config:  %#v", config)
//}
//
//func TestDeleteVm(t *testing.T) {
//	*Debug = true
//	client, err := NewClient("https://astra.ispsystem.net", nil, &tls.Config{InsecureSkipVerify: true}, 300)
//	assert.Nil(t, err)
//	err = client.Login("admin@example.com", "oA4hX9rL")
//	assert.Nil(t, err)
//	vmr := &VmRef{vmId: 41}
//	err = client.DeleteQemuVm(vmr)
//	assert.Nil(t, err)
//}
//func TestChangeParams(t *testing.T) {
//	*Debug = true
//	client, err := NewClient("https://astra.ispsystem.net", nil, &tls.Config{InsecureSkipVerify: true}, 300)
//	assert.Nil(t, err)
//	err = client.Login("admin@example.com", "oA4hX9rL")
//	assert.Nil(t, err)
//	vmr := &VmRef{vmId: 42}
//	err = client.ChangePassword(vmr, "SomeFucking")
//	assert.Nil(t, err)
//	err = client.ChangeOwner(vmr, 3)
//	assert.Nil(t, err)
//	config := ReinstallOS{
//                        Id:             12,
//                        Password:       "FuckingFuck",
//                        EmailMode:      "saas_only",
//                }
//        err = config.ReinstallOS(vmr, client)
//	assert.Nil(t, err)
//	config2 := UpdateConfigQemu{
//		Name: "testing",
//		Description: "russian desc",
//	}
//	err = config2.UpdateConfig(vmr, client)
//	assert.Nil(t, err)
//	config3 := ResourcesQemu{
//		Cores: 1,
//		Memory: 1024,
//	}
//	err = config3.UpdateResources(vmr, client)
//	assert.Nil(t, err)
//	config4 := ConfigDisk{
//		Id: 43,
//		Size: 14000,
//	}
//	err = config4.UpdateDisk(client)
//	assert.Nil(t, err)
//
//}
func TestCreateNetwork(t *testing.T) {
	*Debug = true
	client, err := NewClient("https://astra.ispsystem.net", nil, &tls.Config{InsecureSkipVerify: true}, 300)
	assert.Nil(t, err)
	err = client.Login("admin@example.com", "oA4hX9rL")
	assert.Nil(t, err)
	config := ConfigNewNetwork{
		Name: "1.1.0.0/24",
		Gateway: "1.1.0.1",
		Note: "",
	}
	vmid, err := config.CreateNetwork(client)
	log.Printf(">>> Network id %v", vmid)
	assert.Nil(t, err)
	time.Sleep(5 * time.Second)
	var config2 *ConfigNetwork
	config2, err = NewConfigNetworkFromApi(vmid, client)
	log.Printf(">>> Network properties\n%#v", config2)
	assert.Nil(t, err)
}
