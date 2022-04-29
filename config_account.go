package vmmanager6

import (
	"fmt"
	"encoding/json"
)

type ConfigNewAccount struct {
	Email		string		`json:"email"`
	Role		string		`json:"role"`
	Password	string		`json:"password"`
}
type ConfigAccount struct {
	State		string		`json:"state"`
	Role		string		`json:"role"`
	Id		string		`json:"int"`
	Email		string		`json:"email"`
}

func (config ConfigNewAccount) CreateAccount(client *Client) (vmid string, err error) {
	vmid, err = client.CreateAccount(config)
	if err != nil {
        return "", fmt.Errorf("error creating account: %v (params: %v)", err, config)
    }

	return
}

func NewConfigAccountFromApi(id string, client *Client) (config *ConfigAccount, err error) {
	var accountConfig map[string]interface{}
	accountConfig, err = client.GetAccountInfo(id)
	j, err := json.Marshal(accountConfig)
	err = json.Unmarshal(j, &config)
	return
}

