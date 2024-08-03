package entity

import (
	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)

type DeviceInfo struct {
	Bus     int
	Address int
	Serial  string
	VID     gousb.ID
	Vendor  string
	PID     gousb.ID
	Product string
}

func (di *DeviceInfo) InitName() {
	if v, ok := usbid.Vendors[di.VID]; ok {
		di.Vendor = v.Name
		if p, ok := v.Product[di.PID]; ok {
			di.Product = p.Name
		}
	}
}
