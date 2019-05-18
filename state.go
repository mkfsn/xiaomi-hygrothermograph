package main

import (
	"github.com/paypal/gatt"
)

func onStateChanged(d gatt.Device, s gatt.State) {
	logger.Info.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		logger.Trace.Println("Scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}


