//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package memory

import "syscall"

func mprotect(addr uintptr, len uintptr, prot int) (err error) {
	var data reflect.SliceHeader
	data.Data = addr
	data.Len = int(size)
	data.Cap = int(size)
	var byteslice []byte = *(*[]byte)(unsafe.Pointer(&data))
	err = syscall.Mprotect(byteslice, prot)
	return
}

var ReadOnly = syscall.PROT_READ
var ReadWrite = syscall.PROT_READ | syscall.PROT_WRITE
var Execute = syscall.PROT_EXEC
var ExecuteRead = syscall.PROT_EXEC | syscall.PROT_READ
var ExecuteReadWrite = syscall.PROT_EXEC | syscall.PROT_READ | syscall.PROT_WRITE
