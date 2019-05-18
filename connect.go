package main

import (
	"time"

	"github.com/paypal/gatt"
)

func onPeriphConnected(p gatt.Peripheral, err error) {
	logger.Info.Println("Connected")
	defer p.Device().CancelConnection(p)

	if err := p.SetMTU(500); err != nil {
		logger.Error.Printf("Failed to set MTU, err: %s\n", err)
	}

	// Discovery services
	ss, err := p.DiscoverServices(nil)
	if err != nil {
		logger.Error.Printf("Failed to discover services, err: %s\n", err)
		return
	}

	for _, s := range ss {
		msg := "Service: " + s.UUID().String()
		if len(s.Name()) > 0 {
			msg += " (" + s.Name() + ")"
		}
		logger.Trace.Println(msg)

		// Discovery characteristics
		cs, err := p.DiscoverCharacteristics(nil, s)
		if err != nil {
			logger.Error.Printf("Failed to discover characteristics, err: %s\n", err)
			continue
		}

		for _, c := range cs {
			msg := "  Characteristic  " + c.UUID().String()
			if len(c.Name()) > 0 {
				msg += " (" + c.Name() + ")"
			}
			msg += "\n    properties    " + c.Properties().String()
			logger.Trace.Println(msg)

			// Read the characteristic, if possible.
			if (c.Properties() & gatt.CharRead) != 0 {
				b, err := p.ReadCharacteristic(c)
				if err != nil {
					logger.Error.Printf("Failed to read characteristic, err: %s\n", err)
					continue
				}
				logger.Trace.Printf("    value         %x | %q\n", b, b)
			}

			// Discovery descriptors
			ds, err := p.DiscoverDescriptors(nil, c)
			if err != nil {
				logger.Error.Printf("Failed to discover descriptors, err: %s\n", err)
				continue
			}

			for _, d := range ds {
				msg := "  Descriptor      " + d.UUID().String()
				if len(d.Name()) > 0 {
					msg += " (" + d.Name() + ")"
				}
				logger.Trace.Println(msg)

				// Read descriptor (could fail, if it's not readable)
				b, err := p.ReadDescriptor(d)
				if err != nil {
					logger.Trace.Printf("Failed to read descriptor, err: %s\n", err)
					continue
				}
				logger.Trace.Printf("    value         %x | %q\n", b, b)
			}

			// Subscribe the characteristic, if possible.
			if (c.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
				if err := p.SetNotifyValue(c, notifier); err != nil {
					logger.Trace.Printf("Failed to subscribe characteristic, err: %s\n", err)
					continue
				}
			}

		}
		logger.Trace.Println()
	}

	logger.Info.Printf("Waiting for %s to get some notifiations, if any.\n", timeout)
	time.Sleep(timeout)
	logger.Info.Printf("Waited for %s to get some notifiations, if any.\n", timeout)
}

func onPeriphDisconnected(p gatt.Peripheral, err error) {
	if err != nil {
		logger.Error.Println("error:", err)
	}
	logger.Info.Println("Disconnected")
	p.Device().Connect(p)
	// close(done)
}
