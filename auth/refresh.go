package auth

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"net/http"
	"strings"
)

var (
	refreshTokenUrlFmt = "%s/oauth/token"
)

func (c Client) RefreshToken() error {
	refreshToken := config.GetRefreshToken()
	payload := strings.NewReader(fmt.Sprintf("client_id=%s&grant_type=refresh_token&refresh_token=%s", c.clientId, refreshToken))
	refreshTokenUrl := fmt.Sprintf(refreshTokenUrlFmt, c.auth0BaseUrl)
	req, _ := http.NewRequest(http.MethodPost, refreshTokenUrl, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode > 300 {
		return fmt.Errorf("error when refreshing the token")
	}

	var resp RefreshTokenResponse
	json.NewDecoder(res.Body).Decode(&resp)

	config.SetToken(resp.AccessToken)

	return nil
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
