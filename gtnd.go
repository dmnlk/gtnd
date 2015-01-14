package gtnd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	URL = "http://api.atnd.org/events/"
)

func Search(parameter *SearchParam) (*SearchResult, error) {

	param := url.Values{}
	param.Add("keyword", parameter.Keyword)
	param.Add("format", "json")

	resp, err := http.Get(URL + "?" + param.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *SearchResult
	if err = json.Unmarshal(val, &result); err != nil {
		return nil, err
	}
	return result, nil
}
