package vmmanager6

import (
	"fmt"
	"encoding/json"
//	"log"
)

type ConfigDisk struct {
	Size		int	    `json:"disk_mib"`
	Id		int	    `json:"id"`
}
type ClusterConfig struct {
	Id		int	    `json:"id"`
	DatacenterType  string	    `json:"datacenter_type"`
	Name		string	    `json:"name"`
	Type		string	    `json:"virtualization_type"`
}
type AccountConfig struct {
	Email		string	    `json:"email"`
	Id		int	    `json:"id"`
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
type UpdateConfigQemu struct {
	Name            string      `json:"name"`
	Description     string      `json:"comment"`
}
type ConfigNewQemu struct {
	Name            string      `json:"name"`
	Description     string      `json:"comment"`
	QemuCores       int         `json:"cpu_number"`
	Memory          int         `json:"ram_mib"`
	QemuDisks       int	    `json:"hdd_mib"`
	Cluster		int	    `json:"cluster"`
	Account		int	    `json:"account"`
	Domain		string	    `json:"domain"`
	Os		int         `json:"os"`
	IPv4		int	    `json:"ipv4_number"`
	Password	string	    `json:"password"`
}
type ReinstallOS struct {
	Id		int         `json:"os"`
	Password	string      `json:"password"`
	EmailMode	string      `json:"send_email_mode"`
}
type ResourcesQemu struct {
    Cores		int	`json:"cpu_number"`
    Memory		int	`json:"ram_mib"`
}
// CreateVm - Tell VMmanager 6 API to make the VM
func (config ConfigNewQemu) CreateVm(client *Client) (vmid int, err error) {
	vmid, err = client.CreateQemuVm(config)
	if err != nil {
                return 0, fmt.Errorf("error creating VM: %v (params: %v)", err, config)
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

func (config ResourcesQemu) UpdateResources(vmr *VmRef, client *Client) (err error) {
	err = client.UpdateQemuResources(vmr, config)
	if err != nil {
                return fmt.Errorf("error updating resources of VM id %v: %v (params: %v)", vmr.vmId, err, config)
	}
	return
}

func (config ConfigDisk) UpdateDisk(client *Client) (err error) {
	err = client.UpdateQemuDisk(config)
	if err != nil {
                return fmt.Errorf("error updating disk of VM: %v (params: %v)", err, config)
	}
	return
}

func (config UpdateConfigQemu) UpdateConfig(vmr *VmRef, client *Client) (err error) {
	err = client.UpdateQemuConfig(vmr, config)
	if err != nil {
                return fmt.Errorf("error updating config of VM id %v: %v (params: %v)", vmr.vmId, err, config)
	}
	return
}

func (config ReinstallOS) ReinstallOS(vmr *VmRef, client *Client) (err error) {
	err = client.ReinstallQemu(vmr, config)
	if err != nil {
                return fmt.Errorf("error reinstalling of VM id %v: %v (params: %v)", vmr.vmId, err, config)
	}
	return
}
