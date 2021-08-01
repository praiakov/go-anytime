package main

var x struct{
	a bool
	b int16
	c []int
}

// equivalente a pb := &x.b
pb := (*int16)(unsafe.Pointer(
	uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

*pb = 42

fmt.Prinln(x.b) // "42"