package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	sendInput        = user32.NewProc("SendInput")
)

const (
	// https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
	VK_RBUTTON = 0x02

	// Input types
	INPUT_MOUSE = 0

	// Mouse event flags
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040

	// Configuration - adjust these as needed
	PingInterval   = 1 * time.Second
	PollRateActive = 50 * time.Millisecond  // Fast polling when aiming
	PollRateIdle   = 200 * time.Millisecond // Slow polling when idle
	ClickDelay     = 10 * time.Millisecond
)

type mouseInput struct {
	Dx        int32
	Dy        int32
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type mouseInputWrapper struct {
	Type uint32
	_    [8 - unsafe.Sizeof(uint32(0))]byte // Padding for 64-bit alignment
	Mi   mouseInput
}

func isKeyPressed(vk int) bool {
	ret, _, _ := getAsyncKeyState.Call(uintptr(vk))
	return ret&0x8000 != 0
}

func clickMiddleMouse() {
	var inp mouseInputWrapper
	inp.Type = INPUT_MOUSE

	// Mouse down
	inp.Mi.Flags = MOUSEEVENTF_MIDDLEDOWN
	sendInput.Call(1, uintptr(unsafe.Pointer(&inp)), uintptr(unsafe.Sizeof(inp)))

	time.Sleep(ClickDelay)

	// Mouse up
	inp.Mi.Flags = MOUSEEVENTF_MIDDLEUP
	sendInput.Call(1, uintptr(unsafe.Pointer(&inp)), uintptr(unsafe.Sizeof(inp)))
}

func main() {
	// Graceful shutdown on Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		os.Exit(0)
	}()

	var pressed bool
	var lastPing time.Time

	for {
		state := isKeyPressed(VK_RBUTTON)

		if state {
			if !pressed {
				pressed = true
				lastPing = time.Now()
				clickMiddleMouse()
			} else if time.Since(lastPing) >= PingInterval {
				clickMiddleMouse()
				lastPing = time.Now()
			}
			time.Sleep(PollRateActive)
		} else {
			pressed = false
			time.Sleep(PollRateIdle)
		}
	}
}
