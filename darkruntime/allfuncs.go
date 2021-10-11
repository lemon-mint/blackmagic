package darkruntime

import (
	"runtime"
	"unsafe"
)

func AllFuncs() []*runtime.Func {
	var funcs []*runtime.Func
	for _, m := range *modulesSlice {
		ftab := *(*[]functab)(unsafe.Add(m, moduledata_ftab_off))
		pclntable := *(*[]byte)(unsafe.Add(m, moduledata_pclntable_off))
		for _, f := range ftab {
			_ = f.entry // ignore unused
			if f.funcoff < uintptr(len(pclntable)) {
				function := (*runtime.Func)(unsafe.Pointer(&pclntable[f.funcoff]))
				funcs = append(funcs, function)
			}
		}
	}
	return funcs
}
