// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// HomeMatic XML-API client library
//
// http-client

package homematic

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

type XmlApi struct {
	client    *http.Client
	baseURL   string
	apiURL    string
}

// Initializes a session.
func Init(host string, port string) (*XmlApi, error) {

	api := new(XmlApi)

	api.client = &http.Client{}
	api.baseURL = "http://" + host + "/"
	api.apiURL = api.baseURL + "addons/xmlapi/"

	return api, nil
}

// HTTP GET.
func (api *XmlApi) get(method string, params *url.Values) ([]byte, error) {

	var Url *url.URL
	Url, err := url.Parse(api.apiURL)
	if err != nil {
		return nil, err
	}

	// Build URL
	Url.Path += method + ".cgi"
	// Add optional parameters
	if params != nil {
		Url.RawQuery = params.Encode()
	}
	resp, err := api.client.Get(Url.String())
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
