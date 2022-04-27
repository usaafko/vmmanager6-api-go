package vmmanager6

import (
	"fmt"
	"encoding/json"
//	"log"
)

type ConfigNewNetwork struct {
	Name		string	    `json:"name"`
	Gateway		string	    `json:"gateway"`
	Note		string      `json:"note"`
}
type ConfigNetwork struct {
	Family		int		`json:"family"`
	Gateway		string		`json:"gateway"`
	Id		int		`json:"id"`
	Name		string		`json:"name"`
	Note		string		`json:"note"`
	Size		string		`json:"size"`
	Used		string		`json:"using_ip"`
}

func (config ConfigNewNetwork) CreateNetwork(client *Client) (vmid string, err error) {
	vmid, err = client.CreateNetwork(config)
	if err != nil {
                return "", fmt.Errorf("error creating Network: %v (params: %v)", err, config)
        }

	return
}

func NewConfigNetworkFromApi(id string, client *Client) (config *ConfigNetwork, err error) {
        var netConfig map[string]interface{}
	netConfig, err = client.GetNetworkInfo(id)
	j, err := json.Marshal(netConfig)
	err = json.Unmarshal(j, &config)
	return
}

