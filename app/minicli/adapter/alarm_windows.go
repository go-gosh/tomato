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
	go showMessageBox("gomato", "You have completed a tomato!")
	go func() {
		// call function to play sound 3 times
		for i := 0; i < 3; i++ {
			ptr, _ := syscall.UTF16PtrFromString("alert")
			_, _, err = syscall.Syscall6(win32Fun, 3, uintptr(unsafe.Pointer(ptr)), uintptr(0), uintptr(0), 0, 0, 0)
		}
	}()
}

func showMessageBox(title, text string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	proc := user32.NewProc("MessageBoxW")
	textPtr, _ := syscall.UTF16PtrFromString(text)
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	proc.Call(uintptr(0), uintptr(unsafe.Pointer(textPtr)), uintptr(unsafe.Pointer(titlePtr)), uintptr(0))
}
