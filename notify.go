package main

import (
	"fmt"

	"github.com/paypal/gatt"
)

func notifier(c *gatt.Characteristic, b []byte, err error) {
	if err != nil {
		logger.Error.Println("error:", err)
		return
	}
	var temperature, humidity float64
	fmt.Sscanf(string(b), "T=%f H=%f", &temperature, &humidity)
	logger.Info.Printf("T=%.1f H=%.1f\n", temperature, humidity)
}
