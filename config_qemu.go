package vmmanager6

import (
	"fmt"
)

// ConfigQemu - VMmanager6 API QEMU options
type ConfigQemu struct {
        VmID            int         `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"comment"`
	QemuCores       int         `json:"cpu_number"`
	Memory          int         `json:"ram_mib"`
	QemuDisks       int         `json:"disk_mib"`
}

// CreateVm - Tell VMmanager 6 API to make the VM
func (config ConfigQemu) CreateVm(vmr *VmRef, client *Client) (err error) {
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
	exitStatus, err := client.CreateQemuVm(params)
        if err != nil {
                return fmt.Errorf("error creating VM: %v, error status: %s (params: %v)", err, exitStatus, params)
        }

	return
}
