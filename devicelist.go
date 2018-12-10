// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// HomeMatic XML-API client library
//
// Implementation of deviceList API-method

package homematic

import (
	"net/url"
	"encoding/xml"
	"golang.org/x/net/html/charset"
)

type DeviceList struct {
	XMLName xml.Name `xml:"deviceList"`
	Device  []struct {
		Name        string `xml:"name,attr"`
		Address     string `xml:"address,attr"`
		IseID       string `xml:"ise_id,attr"`
		Interface   string `xml:"interface,attr"`
		DeviceType  string `xml:"device_type,attr"`
		ReadyConfig string `xml:"ready_config,attr"`
		Channel     struct {
			Name             string `xml:"name,attr"`
			Type             string `xml:"type,attr"`
			Address          string `xml:"address,attr"`
			IseID            string `xml:"ise_id,attr"`
			Direction        string `xml:"direction,attr"`
			ParentDevice     string `xml:"parent_device,attr"`
			Index            string `xml:"index,attr"`
			GroupPartner     string `xml:"group_partner,attr"`
			AesAvailable     string `xml:"aes_available,attr"`
			TransmissionMode string `xml:"transmission_mode,attr"`
			Visible          string `xml:"visible,attr"`
			ReadyConfig      string `xml:"ready_config,attr"`
			Operate          string `xml:"operate,attr"`
		} `xml:"channel"`
	} `xml:"device"`
}

// Returns a list of devices
func (api *XmlApi) DeviceList() (*DeviceList, error) {

	var url *url.URL
	url, err := url.Parse(api.apiURL)
	if err != nil {
		return nil, err
	}

	// Build URL
	url.Path += "devicelist.cgi"
	resp, err := api.client.Get(url.String())
	if err != nil {
		return nil, err
	}
	
	// Empty DeviceList object
	var dl DeviceList

	// Parse body
	decoder := xml.NewDecoder(resp.Body)
    decoder.CharsetReader = charset.NewReaderLabel
    err = decoder.Decode(&dl)

	return &dl, nil
}
