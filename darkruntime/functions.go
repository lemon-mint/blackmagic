package darkruntime

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func GetFunctionInfo(name string) *runtime.Func {
	funcs := AllFuncs()
	for _, fn := range funcs {
		if fn.Name() == name {
			return fn
		}
	}
	return nil
}

var ErrFunctionNotFound = fmt.Errorf("function not found")

func GetFunctionByName(name string, fn interface{}) error {
	f := GetFunctionInfo(name)
	if f == nil {
		return ErrFunctionNotFound
	}
	SetFunctionEntry(fn, f.Entry())
	return nil
}

func SetFunctionEntry(fn interface{}, entry uintptr) {
	v := reflect.ValueOf(fn)
	//
	//  *(**uintptr)(unsafe.Pointer(&fn)) = &entry
	//
	*(**uintptr)(unsafe.Pointer(v.Pointer())) = &entry
}
