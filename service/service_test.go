package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type ClientStub struct {
}

func (c *ClientStub) Do(req *http.Request) (*http.Response, error) {
	fmt.Println("request method : ", req.Method)
	if req.Method != "GET" {
		return &http.Response{StatusCode: http.StatusMethodNotAllowed,
			Status: fmt.Sprintf("%d", http.StatusMethodNotAllowed)}, nil
	}
	// if u,p, ok := req.BasicAuth();
	return &http.Response{StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewReader([]byte("unit test")))}, nil
}

func TestGetCurrency(t *testing.T) {
	c := &ClientStub{}
	url := "http://testurl.com"
	res, err := GetCurrency(c, url, "GET")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	if res == nil {
		t.Fail()
	}
	res, err = GetCurrency(c, url, "POST")
	if err == nil {
		t.Errorf("Error :%v", err)
	}

}
