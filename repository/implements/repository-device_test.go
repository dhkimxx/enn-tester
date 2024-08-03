package implements

import (
	"fmt"
	"testing"
)

func Test_GetDeviceList(t *testing.T) {
	fmt.Println("================== TEST Function ==================")
	repo := DeviceRepository_implement{}
	devices, err := repo.GetDeviceList()
	if err != nil {
		t.Error(err)
	}

	for _, device := range devices {
		t.Log(device.Serial)
	}
}
