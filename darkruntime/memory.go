package darkruntime

import "unsafe"

func MemMove(to, from unsafe.Pointer, n uintptr)

//go:linkname MemMove runtime.memmove
