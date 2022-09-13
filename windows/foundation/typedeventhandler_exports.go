// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package foundation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

/*
#include <stdlib.h>
*/
import "C"

//export winrt_TypedEventHandler_QueryInterface
func winrt_TypedEventHandler_QueryInterface(instancePtr, iidPtr unsafe.Pointer, ppvObject *unsafe.Pointer) uintptr {
	// Checkout these sources for more information about the QueryInterface method.
	//   - https://docs.microsoft.com/en-us/cpp/atl/queryinterface
	//   - https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)

	if ppvObject == nil {
		// If ppvObject (the address) is nullptr, then this method returns E_POINTER.
		return ole.E_POINTER
	}

	// This function must adhere to the QueryInterface defined here:
	// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	iid := (*ole.GUID)(iidPtr)
	instance := (*TypedEventHandler)(instancePtr)
	if ole.IsEqualGUID(iid, instance.IID) || ole.IsEqualGUID(iid, ole.IID_IUnknown) || ole.IsEqualGUID(iid, ole.IID_IInspectable) {
		*ppvObject = unsafe.Pointer(instance)
	} else {
		*ppvObject = nil
		// Return E_NOINTERFACE if the interface is not supported
		return ole.E_NOINTERFACE
	}

	// If the COM object implements the interface, then it returns
	// a pointer to that interface after calling IUnknown::AddRef on it.
	(*ole.IUnknown)(*ppvObject).AddRef()

	// Return S_OK if the interface is supported
	return ole.S_OK
}

//export winrt_TypedEventHandler_Invoke
func winrt_TypedEventHandler_Invoke(instancePtr unsafe.Pointer, senderPtr unsafe.Pointer, argsPtr unsafe.Pointer) uintptr {
	// See the quote above.
	instance := (*TypedEventHandler)(instancePtr)
	sender := (unsafe.Pointer)(senderPtr)
	args := (unsafe.Pointer)(argsPtr)
	instance.Callback(instance, sender, args)
	return ole.S_OK
}

//export winrt_TypedEventHandler_AddRef
func winrt_TypedEventHandler_AddRef(instancePtr unsafe.Pointer) uint64 {
	instance := (*TypedEventHandler)(instancePtr)
	return instance.RefCount.AddRef()
}

//export winrt_TypedEventHandler_Release
func winrt_TypedEventHandler_Release(instancePtr unsafe.Pointer) uint64 {
	instance := (*TypedEventHandler)(instancePtr)
	rem := instance.RefCount.Release()
	if rem == 0 {
		// We're done.
		instance.Callback = nil
		C.free(instancePtr)
	}
	return rem
}
