// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// HomeMatic XML-API client library
//
// Example command devicelist
// Loads state information and prints some fields
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	hm "github.com/BFLB/HomeMatic"
)

// Comman-line Arguments
var (
	host          = flag.String("H", "", "Synology Address")
	port          = flag.String("p", "80", "Port")
	name          = flag.String("n", "", "Devicename")
)

func main() {

	// Parse command-line args
	flag.Parse()

	// Init HomeMatic connection
	api, err := hm.Init(*host, *port)
	if err != nil {
		log.Fatal("Init returned error: ", err)
	}

	// Get device-id
	deviceName := *name
	deviceID   := ""

	dl, err := api.DeviceList()
	if err != nil {
		log.Fatal("DeviceList: ", err)
	}

	for i := 0; i < len(dl.Device); i++ {
		if dl.Device[i].Name == deviceName {
			deviceID = dl.Device[i].IseID
			break
		}
	}

	fmt.Printf("DeviceID: %s\n", deviceID)

	state, err := api.State(deviceID)
	if err != nil {
		log.Fatal("State: \n", err)
	}

	var name        string
	var unreach 	bool
	var lowBat  	bool
	var temperature float64
	var humidity	int64

	name = state.Device.Name

	unreach = true
	unreach, _ = strconv.ParseBool(state.Device.Unreach) 

	lowBat = true
	for i := 0; i < len(state.Device.Channel); i++ {
		for x := 0; x < len(state.Device.Channel[i].Datapoint); x++ {
			if state.Device.Channel[i].Datapoint[x].Type == "LOWBAT" {
				lowBat, _ = strconv.ParseBool(state.Device.Channel[i].Datapoint[x].Value)
				break
			}
		}
	}

	for i := 0; i < len(state.Device.Channel); i++ {
		for x := 0; x < len(state.Device.Channel[i].Datapoint); x++ {
			if state.Device.Channel[i].Datapoint[x].Type == "TEMPERATURE" {
				temperature, _ = strconv.ParseFloat(state.Device.Channel[i].Datapoint[x].Value, 64)
				break
			}
		}
	}

	for i := 0; i < len(state.Device.Channel); i++ {
		for x := 0; x < len(state.Device.Channel[i].Datapoint); x++ {
			if state.Device.Channel[i].Datapoint[x].Type == "HUMIDITY" {
				humidity, _ = strconv.ParseInt(state.Device.Channel[i].Datapoint[x].Value, 10, 64)
				break
			}
		}
	}

	fmt.Printf("Name        : %s\n", name)
	fmt.Printf("Unreach     : %t\n", unreach)
	fmt.Printf("LowBat      : %t\n", lowBat)
	fmt.Printf("Temperature : %.1f\n", temperature)
	fmt.Printf("Humidity    : %.d\n", humidity)
	
}
