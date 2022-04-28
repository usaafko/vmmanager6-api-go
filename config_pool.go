package vmmanager6

import (
	"fmt"
	"encoding/json"
)
type ConfigPoolRanges struct {
	Range           string      `json:"name"`
	Id              int         `json:"id"`
}
type ConfigNewPool struct {
	Name		string	    	 `json:"name"`
	Note		string      	 `json:"note"`
	Ranges          []string `json: "ipnets"`
	Cluster		int		`json:"cluster"`
}
type ConfigPool struct {
	Id		string		`json:"id"`
	Name		string		`json:"name"`
	Note		string		`json:"note"`
	Ranges		[]ConfigPoolRanges `json:"ipnets"`
}

func (config ConfigNewPool) CreatePool(client *Client) (vmid string, err error) {
	vmid, err = client.CreatePool(config)
	if err != nil {
                return "", fmt.Errorf("error creating Pool: %v (params: %v)", err, config)
        }

	return
}

func NewConfigPoolFromApi(id string, client *Client) (config *ConfigPool, err error) {
	var poolConfig map[string]interface{}
	poolConfig, err = client.GetPoolInfo(id)
	j, err := json.Marshal(poolConfig)
	err = json.Unmarshal(j, &config)
	return
}

