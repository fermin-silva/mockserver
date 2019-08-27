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

func (e EasyHttpResp) Headers() http.Header {
	return e.resp.Header
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
		return false, fmt.Errorf("Error mashalling string 1 %s. Error: %s", s1, err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 %s. Error: %s", s2, err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func EqualJsonGet(url, expected string) error {
	resp, err := Get(url)

	if err != nil {
		return fmt.Errorf("Error while getting url %s: %s", url, err)
	}

	body, _ := resp.Body()

	if ok, err := AreEqualJSON(expected, body); err != nil {
		return fmt.Errorf("Error while comparing %s and %s. Error: %s", expected, body, err)
	} else if !ok {
		return fmt.Errorf("Expecting\n%s\nbut got\n%s", expected, body)
	}

	return nil
}

func MapContainsExpected(expected, actual map[string][]string) error {
	for k, v := range expected {
		actualV, ok := actual[k]

		if !ok {
			return fmt.Errorf("Expected key %s not found in actual map %v", k, actual)
		}

		if !reflect.DeepEqual(v, actualV) {
			return fmt.Errorf("For key %s, expecting %s but got %s", k, v, actualV)
		}
	}

	return nil
}

func (e EasyHttpResp) HasMatchedFile(expected string) error {
	val := e.Headers().Get("Matched-File")

	if expected != val {
		return fmt.Errorf("Expecting Matched-File Header %s but got %s", expected, val)
	}

	return nil
}
