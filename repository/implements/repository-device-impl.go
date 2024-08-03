package implements

import (
	"enn_tester/entity"
	"fmt"
	"strings"

	"github.com/electricbubble/gadb"
	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)

type DeviceRepository_implement struct{}

func (repo *DeviceRepository_implement) GetDeviceList() ([]entity.DeviceInfo, error) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if strings.Contains(usbid.Describe(desc), "Samsung") || strings.Contains(usbid.Describe(desc), "Google") {
			return true
		} else {
			return false
		}
	})

	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	infos := make([]entity.DeviceInfo, len(devs))

	for i, dev := range devs {

		desc := dev.Desc
		serial, err := dev.SerialNumber()
		if err != nil {
			infos[i] = entity.DeviceInfo{
				Bus:     desc.Bus,
				Address: desc.Address,
				Serial:  "",
				VID:     desc.Vendor,
				PID:     desc.Product,
			}
		} else {
			infos[i] = entity.DeviceInfo{
				Bus:     desc.Bus,
				Address: desc.Address,
				Serial:  serial,
				VID:     desc.Vendor,
				PID:     desc.Product,
			}
		}
		infos[i].InitName()
	}
	return infos, nil
}

func (repo *DeviceRepository_implement) GetDeviceInfo(Serial string) (*entity.DeviceInfo, error) {
	devices, err := repo.GetDeviceList()
	if err != nil {
		return nil, err
	}

	for _, dev := range devices {
		if strings.EqualFold(Serial, dev.Serial) {
			return &dev, nil
		}
	}

	return nil, fmt.Errorf("device not found")
}

func (repo *DeviceRepository_implement) GetDeviceClient(Serial string) (*gadb.Device, error) {
	adbClient, err := gadb.NewClient()
	if err != nil {
		return nil, err
	}

	adbDevs, err := adbClient.DeviceList()
	if err != nil {
		return nil, err
	}

	for _, adbDev := range adbDevs {

		state, err := adbDev.State()
		if err != nil {
			return nil, err
		}

		if strings.EqualFold(Serial, adbDev.Serial()) && state == gadb.StateOnline {
			return &adbDev, nil
		} else {
			return nil, fmt.Errorf("device is not online")
		}
	}
	return nil, fmt.Errorf("device is not detected")
}
