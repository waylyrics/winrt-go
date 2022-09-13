// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package advertisement

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureBluetoothLEAdvertisementReceivedEventArgs string = "rc(Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementReceivedEventArgs;{27987ddf-e596-41be-8d43-9e6731d4a913})"

type BluetoothLEAdvertisementReceivedEventArgs struct {
	ole.IUnknown
}

func (impl *BluetoothLEAdvertisementReceivedEventArgs) GetRawSignalStrengthInDBm() (int16, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisementReceivedEventArgs))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisementReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetRawSignalStrengthInDBm()
}

func (impl *BluetoothLEAdvertisementReceivedEventArgs) GetBluetoothAddress() (uint64, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisementReceivedEventArgs))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisementReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetBluetoothAddress()
}

func (impl *BluetoothLEAdvertisementReceivedEventArgs) GetAdvertisement() (*BluetoothLEAdvertisement, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisementReceivedEventArgs))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisementReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetAdvertisement()
}

const GUIDiBluetoothLEAdvertisementReceivedEventArgs string = "27987ddf-e596-41be-8d43-9e6731d4a913"
const SignatureiBluetoothLEAdvertisementReceivedEventArgs string = "{27987ddf-e596-41be-8d43-9e6731d4a913}"

type iBluetoothLEAdvertisementReceivedEventArgs struct {
	ole.IInspectable
}

type iBluetoothLEAdvertisementReceivedEventArgsVtbl struct {
	ole.IInspectableVtbl

	GetRawSignalStrengthInDBm uintptr
	GetBluetoothAddress       uintptr
	GetAdvertisementType      uintptr
	GetTimestamp              uintptr
	GetAdvertisement          uintptr
}

func (v *iBluetoothLEAdvertisementReceivedEventArgs) VTable() *iBluetoothLEAdvertisementReceivedEventArgsVtbl {
	return (*iBluetoothLEAdvertisementReceivedEventArgsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothLEAdvertisementReceivedEventArgs) GetRawSignalStrengthInDBm() (int16, error) {
	var out int16
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetRawSignalStrengthInDBm,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out int16
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEAdvertisementReceivedEventArgs) GetBluetoothAddress() (uint64, error) {
	var out uint64
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBluetoothAddress,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint64
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEAdvertisementReceivedEventArgs) GetAdvertisement() (*BluetoothLEAdvertisement, error) {
	var out *BluetoothLEAdvertisement
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAdvertisement,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out BluetoothLEAdvertisement
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiBluetoothLEAdvertisementReceivedEventArgs2 string = "12d9c87b-0399-5f0e-a348-53b02b6b162e"
const SignatureiBluetoothLEAdvertisementReceivedEventArgs2 string = "{12d9c87b-0399-5f0e-a348-53b02b6b162e}"

type iBluetoothLEAdvertisementReceivedEventArgs2 struct {
	ole.IInspectable
}

type iBluetoothLEAdvertisementReceivedEventArgs2Vtbl struct {
	ole.IInspectableVtbl

	GetBluetoothAddressType    uintptr
	GetTransmitPowerLevelInDBm uintptr
	GetIsAnonymous             uintptr
	GetIsConnectable           uintptr
	GetIsScannable             uintptr
	GetIsDirected              uintptr
	GetIsScanResponse          uintptr
}

func (v *iBluetoothLEAdvertisementReceivedEventArgs2) VTable() *iBluetoothLEAdvertisementReceivedEventArgs2Vtbl {
	return (*iBluetoothLEAdvertisementReceivedEventArgs2Vtbl)(unsafe.Pointer(v.RawVTable))
}
