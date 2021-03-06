package online

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

type Client struct {
	Token  string
	client *http.Client
}

func NewClient() Client {
	return NewClientWithToken(os.Getenv("ONLINE_TOKEN"))
}

func NewClientWithToken(token string) Client {
	c := Client{
		Token:  token,
		client: &http.Client{},
	}
	return c
}

func (c Client) getApiResource(resName string) ([]byte, error) {
	url := fmt.Sprintf("https://api.cloud.online.net/%s", resName)
	log.Debugf("url:%q", url)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (c Client) postApiResource(resName string, json []byte) ([]byte, error) {
	url := fmt.Sprintf("https://api.cloud.online.net/%s", resName)
	log.Debugf("url:%q", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Debug(resp.Body)

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

// Images return the list of available images
func (c Client) Images() ([]Image, error) {

	body, err := c.getApiResource("images")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var imgRes GetImagesResp
	err = json.Unmarshal(body, &imgRes)
	if err != nil {
		return nil, err
	}

	return imgRes.Images, nil
}

func (c Client) Servers() ([]Server, error) {

	body, err := c.getApiResource("servers")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var serRes GetServersResp
	err = json.Unmarshal(body, &serRes)
	if err != nil {
		return nil, err
	}

	return serRes.Servers, nil
}

func (c Client) Volumes() ([]Volume, error) {

	body, err := c.getApiResource("volumes")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var resp GetVolumesResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Volumes, nil
}

func (c Client) Snapshots() ([]Snapshot, error) {

	body, err := c.getApiResource("snapshots")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var resp GetSnapshotsResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Snapshots, nil
}

func (c Client) IPs() ([]IP, error) {

	body, err := c.getApiResource("ips")
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var resp GetIPsResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.IPs, nil
}

type GetActionsResp struct {
	Actions []string
}

func (c Client) ListActions(serverId string) ([]string, error) {
	body, err := c.getApiResource(fmt.Sprintf("servers/%s/action", serverId))
	if err != nil {
		return nil, err
	}
	log.Debugf("API resp: %s", string(body))

	var resp GetActionsResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Actions, nil
}

func (c Client) DoActions(serverId string, action string) (string, error) {
	json := fmt.Sprintf(`{"action": "%s"}`, action)
	body, err := c.postApiResource(fmt.Sprintf("servers/%s/action", serverId), []byte(json))
	if err != nil {
		return "", err
	}
	log.Debugf("API resp: %s", string(body))

	return string(body), nil
}
