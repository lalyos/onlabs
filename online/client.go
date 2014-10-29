package online

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type Client struct {
	Token  string
	client *http.Client
}

func NewClient(token string) Client {
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

	if resp.StatusCode != 200 {
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
	//log.Debug(string(body))

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
	//log.Debug(string(body))

	var serRes GetServersResp
	err = json.Unmarshal(body, &serRes)
	if err != nil {
		return nil, err
	}

	return serRes.Servers, nil
}
