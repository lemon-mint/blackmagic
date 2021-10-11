package darkruntime

import (
	"reflect"
	"unsafe"
)

func SetFunctionEntry(fn interface{}, entry uintptr) {
	v := reflect.ValueOf(fn)
	//
	//  *(**uintptr)(unsafe.Pointer(&fn)) = &entry
	//
	*(**uintptr)(unsafe.Pointer(v.Pointer())) = &entry
}
