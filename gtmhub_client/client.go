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

//const(
//	gtmhubBaseUrl = "https://app.gtmhub.com"
//)

type GtmhubHttpClient struct{

}

func executeRequest(url, method string, body []byte) ([]byte, error){
	breaker := 0
	for {
		if breaker > 10 {
			return nil, fmt.Errorf("could not make request")
		}
		token := config.GetToken()
		accountID := config.GetAccountId()

		req, _ := http.NewRequest(method, url, bytes.NewReader(body))

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("gtmhub-accountid", accountID)

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()

		if res.StatusCode == http.StatusUnauthorized {
			//// token expired, use the refresh token
			//// todo: implement it
			//return nil, fmt.Errorf("token expired")
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
			fmt.Println(string(bodyResp))
			return nil, fmt.Errorf("something is not quite right")
		}

		//bodyResp, _ := ioutil.ReadAll(res.Body)

		//fmt.Println(string(bodyResp))

		return bodyResp, nil
	}
}
