package gtmhub_client

import (
	"encoding/json"
	"fmt"
	"gtmhub-cli/config"
	"gtmhub-cli/model"
	"net/http"
	"net/url"
)

//var (
//	getListsUrl = fmt.Sprintf("%s/api/v1/lists", gtmhubBaseUrl)
//	loadListUrl = fmt.Sprintf("%s/api/v1/lists",gtmhubBaseUrl)
//	loadListFmt = loadListUrl + "/%s/load"
//)

var (
	listsBaseUrlFmt = "%s/api/v1/lists"
	loadListsFmt    = listsBaseUrlFmt + "/%s/load"
)

func (ghc GtmhubHttpClient) LoadList(list model.ListResponse) ([]map[string]interface{}, error) {
	url := fmt.Sprintf(loadListsFmt, config.GetGtmhubUrl(), list.ID)
	requestBody := model.LoadRequest{
		Columns: list.Columns,
		Filter: model.Filter{
			BooleanOperator: "and",
			RuleBounds:      list.Filter.RuleBounds,
		}}

	bodyR, _ := json.Marshal(requestBody)

	//fmt.Printf(string(bodyR))

	body, err := executeRequest(url, http.MethodPost, bodyR)
	var response map[string]interface{}

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &response)
	itemsArray, ok := response["items"]
	if ok == false {
		return nil, fmt.Errorf("there was an error getting the items from the list")
	}

	items := itemsArray.([]interface{})
	result := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		result = append(result, item.(map[string]interface{}))
	}

	return result, nil
}

func (ghc GtmhubHttpClient) GetListsByName(listName string) (model.FullListResponse, error) {
	filter := fmt.Sprintf("{title:{$regex:\".*%s.*\"}, listType:\"Key Result\"}", listName)
	return getListsByFilter(filter, ghc)
}

func (ghc GtmhubHttpClient) GetListsByID(id string) (model.FullListResponse, error) {
	filter := fmt.Sprintf("{id:{$regex:\".*%s.*\"}, listType:\"Key Result\"}", id)

	return getListsByFilter(filter, ghc)
}

func getListsByFilter(filter string, ghc GtmhubHttpClient) (model.FullListResponse, error) {
	filter = url.QueryEscape(filter)
	baseFilter := "?fields=title,columns,filter&limit=100&skip=0&filter=" + filter

	listUrl := fmt.Sprintf(listsBaseUrlFmt, config.GetGtmhubUrl())
	body, err := executeRequest(listUrl+baseFilter, http.MethodGet, nil)
	if err != nil {
		return model.FullListResponse{}, err
	}

	var response model.FullListResponse
	json.Unmarshal(body, &response)

	return response, nil

}

//func (ghc GtmhubHttpClient) LoadList(lName string) (model.KRListFullResponse, error) {
//	listId, err := GetListID(lName)
//	if err != nil {
//		return model.KRListFullResponse{}, err
//	}
//
//	loadListFmt := fmt.Sprintf(loadListsFmt, config.GetGtmhubUrl(), listId)
//
//
//	body, err := executeRequest(fmt.Sprintf(loadListFmt, listId), http.MethodGet, nil)
//
//	if err != nil {
//		return model.KRListFullResponse{}, err
//	}
//
//	var resp model.KRListFullResponse
//	json.Unmarshal(body, &resp)
//
//	return resp, nil
//
//
//}
//
//func GetListID(lName string) (string, error) {
//	query := ""
//	if len(lName) > 0 {
//		query = fmt.Sprintf("{\"title\":\"%s\"}", lName)
//		query = url.QueryEscape(query)
//		query = "?filter=" + query
//	}
//
//	getListsUrl := fmt.Sprintf(listsBaseUrlFmt, config.GetGtmhubUrl())
//
//	body, err := executeRequest(fmt.Sprintf("%s%s", getListsUrl, query), http.MethodGet, nil)
//
//	if err != nil {
//		return "", err
//	}
//
//	var resp model.ListFullResponse
//	json.Unmarshal(body, &resp)
//	if len(resp.Items) > 0 {
//		return resp.Items[0].ID, nil
//	}
//	return "", fmt.Errorf("no lists found")
//}
