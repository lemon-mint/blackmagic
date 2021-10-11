package darkreflect

import (
	"reflect"
	"unsafe"
)

func ToType(t unsafe.Pointer) reflect.Type

//go:linkname ToType reflect.toType

func Typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname Typelinks reflect.typelinks

func RtypeOff(section unsafe.Pointer, off int32) unsafe.Pointer

//go:linkname RtypeOff reflect.rtypeOff

func TypesByString(s string) []unsafe.Pointer

//go:linkname TypesByString reflect.typesByString
