package vmmanager6

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Client struct {
	session		*Session
	ApiUrl		string
	Username	string
	Password	string
	TaskTimeout	int
}
type VmRef struct {
	vmId	int
}

func NewClient(apiUrl string, hclient *http.Client, tls *tls.Config, taskTimeout int) (client *Client, err error) {
	var sess *Session
	sess, err = NewSession(apiUrl, hclient, tls)
	if err == nil {
		client = &Client{session: sess, ApiUrl: apiUrl, TaskTimeout: taskTimeout}
	}
	return client, err
}
func (c *Client) SetAPIToken(token string) {
	c.session.SetAPIToken(token)
}
func (c *Client) Login(username string, password string) (err error) {
	c.Username = username
	c.Password = password
	return c.session.Login(username, password)
}
func (c *Client) GetItemConfigMapStringInterface(url, text string) (map[string]interface{}, error) {
	data, err := c.GetItemConfig(url, text)
	if err != nil {return nil, err}
	return data["data"].(map[string]interface{}), err
}

func (c *Client) GetItemConfigString(url, text string) (string, error) {
	data, err := c.GetItemConfig(url, text)
	if err != nil {return "", err}
	return data["data"].(string), err
}

func (c *Client) GetItemConfigInterfaceArray(url, text string) ([]interface{}, error) {
	data, err := c.GetItemConfig(url, text)
	if err != nil {return nil, err}
	return data["data"].([]interface{}), err
}

func (c *Client) GetItemConfig(url, text string) (config map[string]interface{}, err error) {
	err = c.GetJsonRetryable(url, &config, 3)
	if err != nil {return nil, err}
	if config["data"] == nil {return nil, fmt.Errorf(text + " CONFIG not readable")}
	return
}

func (c *Client) GetJsonRetryable(url string, data *map[string]interface{}, tries int) error {
	var statErr error
	for ii := 0; ii < tries; ii++ {
		_, statErr = c.session.GetJSON(url, nil, nil, data)
		if statErr == nil {
			return nil
		}
		log.Printf("[DEBUG][GetJsonRetryable] Sleeping for %d seconds before asking url %s", ii+1, url)
		time.Sleep(time.Duration(ii+1) * time.Second)
	}
	return statErr
}
func (c *Client) GetNodeList() (list map[string]interface{}, err error) {
	err = c.GetJsonRetryable("/node", &list, 3)
	return
}
func (c *Client) GetVmList() (list map[string]interface{}, err error) {
	err = c.GetJsonRetryable("/host", &list, 3)
	return
}
func (c *Client) DeleteUrl(url string) (err error) {
	_, err = c.session.Delete(url, nil, nil)
	return
}
func NewVmRef(vmId int) (vmr *VmRef) {
        vmr = &VmRef{vmId: vmId}
        return
}

func (c *Client) GetVmInfo(vmr *VmRef) (vmInfo map[string]interface{}, err error) {
        resp, err := c.GetVmList()
        vms := resp["list"].([]interface{})
        for vmii := range vms {
                vm := vms[vmii].(map[string]interface{})
		if vm["id"].(int) == vmr.vmId {
                        vmInfo = vm
                        return
                }
        }
        return nil, fmt.Errorf("vm '%d' not found", vmr.vmId)
}


