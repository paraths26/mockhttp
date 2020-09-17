package service

import (
	"errors"
	"io/ioutil"
	"net/http"
)

//Client:
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

//GetCurrency :
func GetCurrency(c Client, url, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("un expected http status :" + res.Status)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil || body == nil {
		return nil, err
	}
	return body, nil
}
