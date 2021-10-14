package memory

import "syscall"

var pagesize uintptr = uintptr(syscall.Getpagesize())

func GetPageHead(ptr uintptr) uintptr {
	return pagesize * (ptr / pagesize)
}

func NextPageHead(curpage uintptr) uintptr {
	return curpage + pagesize
}

func Mprotect(ptr uintptr, size uintptr, prot int) error {
	return mprotect(ptr, size, prot)
}
