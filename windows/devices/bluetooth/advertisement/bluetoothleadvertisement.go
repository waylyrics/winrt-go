// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint
package advertisement

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type BluetoothLEAdvertisement struct {
	iBluetoothLEAdvertisement
}

func NewBluetoothLEAdvertisement() (*BluetoothLEAdvertisement, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisement")
	if err != nil {
		return nil, err
	}
	return (*BluetoothLEAdvertisement)(unsafe.Pointer(inspectable)), nil
}

const GUIDiBluetoothLEAdvertisement string = "066fb2b7-33d1-4e7d-8367-cf81d0f79653"

type iBluetoothLEAdvertisement struct {
	ole.IInspectable
}

type iBluetoothLEAdvertisementVtbl struct {
	ole.IInspectableVtbl

	GetFlags                       uintptr
	SetFlags                       uintptr
	GetLocalName                   uintptr
	SetLocalName                   uintptr
	GetServiceUuids                uintptr
	GetManufacturerData            uintptr
	GetDataSections                uintptr
	GetManufacturerDataByCompanyId uintptr
	GetSectionsByType              uintptr
}

func (v *iBluetoothLEAdvertisement) VTable() *iBluetoothLEAdvertisementVtbl {
	return (*iBluetoothLEAdvertisementVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothLEAdvertisement) GetLocalName() (string, error) {
	var out string
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetLocalName,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	return out, nil
}
