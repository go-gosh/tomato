//go:build windows
// +build windows

package adapter

import (
	"syscall"
	"unsafe"
)

func (d dataManager) Alarm() {
	funInDllFile, err := syscall.LoadLibrary("Winmm.dll") // call dll file
	if err != nil {
		print("cant not call : syscall.LoadLibrary , errorInfo :" + err.Error())
	}
	defer syscall.FreeLibrary(funInDllFile)
	// function name in dll
	funName := "PlaySound"
	// register handle
	win32Fun, err := syscall.GetProcAddress(syscall.Handle(funInDllFile), funName)
	// call function to play sound 3 times
	for i := 0; i < 3; i++ {
		ptr, _ := syscall.UTF16PtrFromString("alert")
		_, _, err = syscall.Syscall6(win32Fun, 3, uintptr(unsafe.Pointer(ptr)), uintptr(0), uintptr(0), 0, 0, 0)
	}
}
