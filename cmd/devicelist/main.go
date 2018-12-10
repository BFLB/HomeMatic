// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// HomeMatic XML-API client library
//
// Example command devicelist
// Loads devicelist information and prints some fields
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	hm "github.com/BFLB/HomeMatic"
)

// Comman-line Arguments
var (
	host          = flag.String("H", "", "Synology Address")
	port          = flag.String("p", "80", "Port")
)

func main() {

	// Parse command-line args
	flag.Parse()

	// Init HomeMatic connection
	api, err := hm.Init(*host, *port)
	if err != nil {
		log.Fatal("Init returned error: ", err)
	}

	dl, err := api.DeviceList()
	if err != nil {
		log.Fatal("DeviceList: ", err)
		fmt.Printf("Pups")
	}

	// initialize tabwriter
	w := new(tabwriter.Writer)
	
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	
	defer w.Flush()

	
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", "Name", "Address", "IseID", "Interface", "DeviceType", "ReadyConfig")
	for i := 0; i < len(dl.Device); i++ {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", dl.Device[i].Name, dl.Device[i].Address, dl.Device[i].IseID, dl.Device[i].Interface, dl.Device[i].DeviceType,dl.Device[i].ReadyConfig )
	}
}
