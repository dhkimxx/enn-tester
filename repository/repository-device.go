package repository

import (
	"enn_tester/entity"
	"enn_tester/repository/implements"

	"github.com/electricbubble/gadb"
)

type DeviceRepository interface {
	GetDeviceClient(Serial string) (*gadb.Device, error)
	GetDeviceList() ([]entity.DeviceInfo, error)
	GetDeviceInfo(Serial string) (*entity.DeviceInfo, error)
}

func GetDeviceRepository() DeviceRepository {
	var result DeviceRepository
	f := implements.DeviceRepository_implement{}
	result = &f
	return result
}
