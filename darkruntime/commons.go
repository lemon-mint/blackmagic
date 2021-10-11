package darkruntime

import (
	"reflect"
	"unsafe"

	"github.com/lemon-mint/blackmagic/darkreflect"
)

var modulesSlice *[]unsafe.Pointer //*[]*moduledata

//go:linkname modulesSlice runtime.modulesSlice

var moduledata_t reflect.Type = darkreflect.MustGetType("*runtime.moduledata").Elem()
var moduledata_ftab_off uintptr = darkreflect.MustGetOffset(moduledata_t, "ftab")           // []functab
var moduledata_pclntable_off uintptr = darkreflect.MustGetOffset(moduledata_t, "pclntable") // []byte

type functab struct {
	entry   uintptr
	funcoff uintptr
}

func Modulesinit()

//go:linkname Modulesinit runtime.modulesinit

var _ int = func() int {
	Modulesinit()
	return 0
}()
