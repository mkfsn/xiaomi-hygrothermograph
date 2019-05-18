package main

import (
	"flag"
	"strings"

	"github.com/paypal/gatt"
)

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	id := strings.ToUpper(flag.Args()[0])
	if strings.ToUpper(p.ID()) != id {
		return
	}

	// Stop scanning once we've got the peripheral we're looking for.
	p.Device().StopScanning()

	logger.Info.Printf("\nPeripheral ID:%s, NAME:(%s)\n", p.ID(), p.Name())
	logger.Info.Println("  Local Name        =", a.LocalName)
	logger.Info.Println("  TX Power Level    =", a.TxPowerLevel)
	logger.Info.Println("  Manufacturer Data =", a.ManufacturerData)
	logger.Info.Println("  Service Data      =", a.ServiceData)
	logger.Info.Println("")

	p.Device().Connect(p)
}

