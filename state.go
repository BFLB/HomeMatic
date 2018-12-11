// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// HomeMatic XML-API client library
//
// Implementation of state API-method

package homematic

import (
	"net/url"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"errors"
)

type State struct {
	XMLName xml.Name `xml:"state"`
	Device  struct {
		Name          string `xml:"name,attr"`
		IseID         string `xml:"ise_id,attr"`
		Unreach       string `xml:"unreach,attr"`
		StickyUnreach string `xml:"sticky_unreach,attr"`
		ConfigPending string `xml:"config_pending,attr"`
		Channel       []struct {
			Name      string `xml:"name,attr"`
			IseID     string `xml:"ise_id,attr"`
			Datapoint []struct {
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
				IseID     string `xml:"ise_id,attr"`
				Value     string `xml:"value,attr"`
				Valuetype string `xml:"valuetype,attr"`
				Valueunit string `xml:"valueunit,attr"`
				Timestamp string `xml:"timestamp,attr"`
			} `xml:"datapoint"`
		} `xml:"channel"`
	} `xml:"device"`
}

// Returns a list of devices
func (api *XmlApi) State(IseID string) (*State, error) {

	if IseID == "" {
		return nil, errors.New("Empty name")
	}

	var u *url.URL
	u, err := url.Parse(api.apiURL)
	if err != nil {
		return nil, err
	}

	// Build URL
	u.Path += "state.cgi"

	// Set URL parameters
	params := url.Values{}
	params.Add("device_id", IseID)
	u.RawQuery = params.Encode()

	// HTTP GET
	resp, err := api.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	
	// Empty DeviceList object
	var state State
	
	decoder := xml.NewDecoder(resp.Body)
    decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&state)
	if err != nil {
		return nil, err
	}
	
	return &state, nil
}
