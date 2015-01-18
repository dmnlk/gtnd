package gtnd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	URL = "http://api.atnd.org/events/"
)

func Search(parameter *SearchParam) (*SearchResult, error) {
	a, _ := addOptions(URL, parameter)
	//	pp.Print(a)
	//	param := url.Values{}
	//	param.Add("keyword", parameter.Keyword)
	//	param.Add("format", "json")

	resp, err := http.Get(a)
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

func addOptions(r string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return r, nil
	}
	u, err := url.Parse(r)
	if err != nil {
		return r, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return r, err
	}
	u.RawQuery = qs.Encode()
	return u.String(), nil
}
