//go:build windows

package main

import (
	"os"
	"syscall"
)

func attachConsole() {
	modkernel32 := syscall.NewLazyDLL("kernel32.dll")
	procAttachConsole := modkernel32.NewProc("AttachConsole")

	// ATTACH_PARENT_PROCESS is (DWORD)-1
	procAttachConsole.Call(^uintptr(0))

	// Re-open standard output to the console.
	// Windows GUI apps start with disconnected stdout/stderr.
	hout, err := syscall.Open("CONOUT$", syscall.O_RDWR, 0)
	if err == nil {
		os.Stdout = os.NewFile(uintptr(hout), "stdout")
		os.Stderr = os.NewFile(uintptr(hout), "stderr")
	}
}
