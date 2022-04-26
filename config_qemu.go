package vmmanager6

import (
	"fmt"
	"encoding/json"
//	"log"
)

type ConfigDisk struct {
	Size		int	    `json:"disk_mib"`
	DiskId		int	    `json:"id"`
}
type ClusterConfig struct {
	Id		int	    `json:"id"`
	DatacenterType  string	    `json:"datacenter_type"`
	Name		string	    `json:"name"`
	Type		string	    `json:"virtualization_type"`
}
type AccountConfig struct {
	Email		string	    `json:"email"`
	AccountId	int	    `json:"id"`
}
type OsConfig struct {
	Id		int	    `json:"id"`
}
type Ipv4Config struct {
	Interface	string	    `json:"interface"`
	Ip		string	    `json:"ip"`
}
// ConfigQemu - VMmanager6 API QEMU options
type ConfigQemu struct {
	Name            string      `json:"name"`
	Description     string      `json:"comment"`
	QemuCores       int         `json:"cpu_number"`
	Memory          int         `json:"ram_mib"`
	QemuDisks       ConfigDisk  `json:"disk"`
	Cluster		ClusterConfig	`json:"cluster"`
	Account		AccountConfig	`json:"account"`
	Domain		string	    `json:"domain"`
	Os		OsConfig    `json:"os"`
	IPv4		[]Ipv4Config	`json:"ip4"`
}

// CreateVm - Tell VMmanager 6 API to make the VM
func (config ConfigQemu) CreateVm(client *Client) (vmid int, err error) {
	params := map[string]interface{}{
		"name": config.Name,
		"comment": config.Description,
		"ram_mib": config.Memory,
		"hdd_mib": config.QemuDisks,
		"cpu_number": config.QemuCores,
		"cluster": 1,
		"account": 3,
		"domain": "fuck.me",
		"password": "fuckingfuck",
		"os": 1,
		"ipv4_number": 1,
	}
	vmid, err = client.CreateQemuVm(params)
	if err != nil {
                return 0, fmt.Errorf("error creating VM: %v (params: %v)", err, params)
        }

	return
}

func NewConfigQemuFromApi(vmr *VmRef, client *Client) (config *ConfigQemu, err error) {
        var vmConfig map[string]interface{}
	vmConfig, err = client.GetVmInfo(vmr)
	j, err := json.Marshal(vmConfig)
	err = json.Unmarshal(j, &config)
	return
}
