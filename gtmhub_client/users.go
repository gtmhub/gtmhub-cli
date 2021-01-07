package gtmhub_client

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"gtmhub-cli/model"
	"net/http"
	url2 "net/url"
)

var(
	userUrlFmt = "%s/api/v1/users"
)

func (ghc GtmhubHttpClient) GetUserID(domain, auth0id string) (string, error) {

	userUrl := fmt.Sprintf(userUrlFmt, config.GetGtmhubUrl())

	url := fmt.Sprintf("%s/%s/%s", userUrl, domain, url2.QueryEscape(auth0id))

	body, err := executeRequest(url, http.MethodGet, nil)
	if err != nil {
		return "", err
	}

	var resp model.ResolveUserResponse
	json.Unmarshal(body, &resp)
	return resp.User.ID, nil
}


