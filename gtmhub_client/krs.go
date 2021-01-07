package gtmhub_client

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"gtmhub-cli/model"
	"net/http"
	"net/url"
)

var (
	krUrlV2Fmt = "%s/api/v2/metrics"
	KrUrlV1Fmt = "%s/api/v1/metrics"
)

func (ghc GtmhubHttpClient) UpdateMetric(request model.CheckInMetricRequest, id string) error {

	krUrlV1 := fmt.Sprintf(KrUrlV1Fmt, config.GetGtmhubUrl())

	url := fmt.Sprintf("%s/%s/checkin", krUrlV1, id)
	body, _ := json.Marshal(request)
	_, err := executeRequest(url, http.MethodPost, body)

	return err
}

func (ghc GtmhubHttpClient) GetMetricsInCurrentSession() (model.Metrics, error) {

	krUrlV2 := fmt.Sprintf(krUrlV2Fmt, config.GetGtmhubUrl())

	sessionIds, err := ghc.GetActiveSessionsIDs()

	if err != nil {
		return model.Metrics{}, err
	}
	if len(sessionIds) == 0 {
		return model.Metrics{}, fmt.Errorf("no active sessions found")
	}

	queryparam := sessionIds.ToQueryIDs()

	query := fmt.Sprintf("{sessionId:{$in:%s}, ownerId:\"%s\"}", queryparam, config.GetUserID())
	query = url.QueryEscape(query)
	query = "?filter=" + query

	body, err := executeRequest(krUrlV2+query, http.MethodGet, nil)
	if err != nil {
		return model.Metrics{}, err
	}

	var fullMetricResponse model.FulLMetricResponse
	json.Unmarshal(body, &fullMetricResponse)

	return fullMetricResponse.Items, nil
}






