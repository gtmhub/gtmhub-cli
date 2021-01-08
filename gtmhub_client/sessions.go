package gtmhub_client

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"gtmhub-cli/model"
	"net/http"
)

var (
	sessionUrlFmt = "%s/api/v1/sessions"
)

func (ghc GtmhubHttpClient) GetActiveSessionsIDs() (model.IDResponses, error) {

	sessionStatus := "open"
	timeFrameFilter := "current"

	query := fmt.Sprintf("filter=%s&status=%s", timeFrameFilter, sessionStatus)
	query = "?" + query

	sessionUrl := fmt.Sprintf(sessionUrlFmt, config.GetGtmhubUrl())

	resp, err := executeRequest(sessionUrl + query, http.MethodGet, nil)
	if err != nil {
		return model.IDResponses{}, err
	}

	var response model.FullIDResponse
	json.Unmarshal(resp, &response)

	return response.Items, nil

}


