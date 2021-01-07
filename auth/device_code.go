package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	tokenUrlFmt      = "%s/oauth/token"
	deviceCodeUrlFmt = "%s/oauth/device/code"
)



func (c *Client) PoolForToken(deviceCode DeviceFlowInitResponse) (AccessCodeResponse, error) {
	count := 0
	for true {
		if count > 30 {
			break
		}
		payload := strings.NewReader(fmt.Sprintf("grant_type=urn:ietf:params:oauth:grant-type:device_code&device_code=%s&client_id=%s", deviceCode.DeviceCode, c.clientId))
		tokenUrl := fmt.Sprintf(tokenUrlFmt, c.auth0BaseUrl)
		req, _ := http.NewRequest("POST", tokenUrl, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		if res.StatusCode > 400 && res.StatusCode < 500 {
			time.Sleep(time.Second * time.Duration(deviceCode.Interval))
			count++
			continue
		}

		if res.StatusCode != 200 {
			return AccessCodeResponse{}, fmt.Errorf("error while waiting for access code. code: %d", res.StatusCode)
		}

		body, _ := ioutil.ReadAll(res.Body)

		//fmt.Println(res)
		//fmt.Println(string(body))

		var resp AccessCodeResponse
		json.Unmarshal([]byte(body), &resp)

		return resp, nil
	}

	return AccessCodeResponse{}, fmt.Errorf("request expired")
}

func (c *Client) InitAuth() (DeviceFlowInitResponse, error) {

	payload := strings.NewReader(fmt.Sprintf("client_id=%s&scope=openid offline_access&audience=%s", c.clientId, c.audience))
	deviceCodeUrl := fmt.Sprintf(deviceCodeUrlFmt, c.auth0BaseUrl)
	req, _ := http.NewRequest("POST", deviceCodeUrl, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var resp DeviceFlowInitResponse
	json.NewDecoder(res.Body).Decode(&resp)

	return resp, nil
}

type DeviceFlowInitResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURL string `json:"verification_uri"`
	Interval        int    `json:"interval"`
}

type AccessCodeResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}
