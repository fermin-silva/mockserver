package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

type EasyHttpResp struct {
	resp *http.Response
}

func (e EasyHttpResp) Body() (string, error) {
	defer e.resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(e.resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func Get(url string) (*EasyHttpResp, error) {
	url = filepath.Join("/", url)

	resp, err := netClient.Get("http://localhost:8080" + url)

	if err != nil {
		return nil, err
	}

	return &EasyHttpResp{resp}, err
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
