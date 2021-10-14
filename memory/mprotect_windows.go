//go:build windows

package memory

import (
	"fmt"
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var procVirtualProtect = kernel32.NewProc("VirtualProtect")

var ErrVirtualProtect = fmt.Errorf("virtualprotect failed")

func mprotect(addr uintptr, size uintptr, prot int) (err error) {
	var oldprot uint32

	ret, _, _ := procVirtualProtect.Call(
		uintptr(addr),
		uintptr(size),
		uintptr(prot),
		uintptr(unsafe.Pointer(&oldprot)),
	)
	if ret > 0 {
		return nil
	}
	return ErrVirtualProtect
}

var ReadOnly = 0x02
var ReadWrite = 0x04
var Execute = 0x10
var ExecuteRead = 0x20
var ExecuteReadWrite = 0x40
