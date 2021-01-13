package gtmhub_client

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"gtmhub-cli/model"
	"net/http"
)

var (
	accountUrlFmt = "%s/api/v1/users/app/%s"
)

func (ghc GtmhubHttpClient) ResolveAccount(auth0UserId string) (model.AccountResolveResponse, error) {
	accountUrl := fmt.Sprintf(accountUrlFmt, config.GetGtmhubUrl(), auth0UserId)
	//url := fmt.Sprintf("%s/%s",accountUrl, config.GetAccountId())
	body, err := executeGlobalRequest(accountUrl, http.MethodGet, nil)
	if err != nil {
		return model.AccountResolveResponse{}, err
	}

	var accResponse model.AccountResolveResponse
	json.Unmarshal(body, &accResponse)

	return accResponse, nil
}