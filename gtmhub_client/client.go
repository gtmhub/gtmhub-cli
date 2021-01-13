package gtmhub_client

import (
	"bytes"
	"fmt"
	"gtmhub-cli/auth"
	"gtmhub-cli/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type GtmhubHttpClient struct{

}

func executeGlobalRequest(url, method string, body []byte) ([]byte, error) {
	return executeRequestInternal(url, method, "", body)
}

func executeRequestInternal(url, method, accountID string, body []byte) ([]byte, error) {
	breaker := 0
	for {
		if breaker > 10 {
			return nil, fmt.Errorf("could not make request")
		}
		token := config.GetToken()


		req, _ := http.NewRequest(method, url, bytes.NewReader(body))

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		if len(accountID) > 0 {
			req.Header.Add("gtmhub-accountid", accountID)
		}

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()

		if res.StatusCode == http.StatusUnauthorized {
			breaker++
			authClient, err := auth.GetClient()
			if err != nil {
				return nil, fmt.Errorf("could not initialize auth client")
			}
			err =  authClient.RefreshToken()
			if err != nil {
				log.Printf("error while refreshing")
				time.Sleep(time.Second * 5)
			}
			continue
		}
		bodyResp, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode > 300 {
			return nil, fmt.Errorf("something is not quite right. err: %s", string(bodyResp))
		}

		return bodyResp, nil
	}
}

func executeRequest(url, method string, body []byte) ([]byte, error){
	return executeRequestInternal(url, method, config.GetAccountId(), body)
}
