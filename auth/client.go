package auth

import (
	"fmt"
	"gtmhub-cli/config"
)

var (
	client *Client
	dcs = map[string]DC{
		"eu": {clientID: "g3uiRNIOdckZ8uE68bRNuMCJRMQwYsdU", dcUrl: "https://auth.gtmhub.com", apiUrl: "https://app.gtmhub.com/api"},
		"us": {clientID: "0Wg1NXn0slGzgtX4aRvdwJpiECR83HYN", dcUrl: "https://auth.us.gtmhub.com", apiUrl: "https://app.us.gtmhub.com/api"},
	}
)

type DC struct {
	clientID string
	dcUrl string
	apiUrl string
}


type Client struct{
	auth0BaseUrl string
	clientId string
	audience string
}

func GetClient() (*Client, error) {
	dc := config.GetGtmhubDC()
	if client == nil {
		if err := initClient(dc); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func initClient(dc string) error{
	conf, known := dcs[dc]
	if known == false {
		return fmt.Errorf("uknown data center specified: %s.", dc)
	}

	client = &Client{
		clientId: conf.clientID,
		auth0BaseUrl: conf.dcUrl,
		audience: conf.apiUrl,
	}

	return nil
}
