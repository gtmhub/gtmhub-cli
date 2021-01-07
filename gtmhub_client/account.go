package gtmhub_client

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"gtmhub-cli/model"
	"net/http"
)

var (
	accountUrlFmt = "%s/api/v1/accounts"
)

func (ghc GtmhubHttpClient) GetAccountDomain() (string, error) {
	accountUrl := fmt.Sprintf(accountUrlFmt, config.GetGtmhubUrl())
	url := fmt.Sprintf("%s/%s",accountUrl, config.GetAccountId())
	body, err := executeRequest(url, http.MethodGet, nil)
	if err != nil {
		return "", err
	}

	var accResponse model.AccountResponse
	json.Unmarshal(body, &accResponse)

	return accResponse.Domain, nil
}