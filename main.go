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
	VK_RBUTTON  = 0x02
	VK_LCONTROL = 0xA2

	INPUT_KEYBOARD  = 1
	KEYEVENTF_KEYUP = 0x0002

	// Configuration - adjust these as needed
	PingInterval   = 1 * time.Second
	PollRateActive = 50 * time.Millisecond  // Fast polling when aiming
	PollRateIdle   = 200 * time.Millisecond // Slow polling when idle
	KeyPressDelay  = 10 * time.Millisecond
)

type keyboardInput struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type input struct {
	Type uint32
	_    [8 - unsafe.Sizeof(uint32(0))]byte // Padding for 64-bit alignment
	Ki   keyboardInput
}

func isKeyPressed(vk int) bool {
	ret, _, _ := getAsyncKeyState.Call(uintptr(vk))
	return ret&0x8000 != 0
}

func pressKey(vk int) {
	var inp input
	inp.Type = INPUT_KEYBOARD
	inp.Ki.Vk = uint16(vk)

	// Key down
	sendInput.Call(1, uintptr(unsafe.Pointer(&inp)), uintptr(unsafe.Sizeof(inp)))

	time.Sleep(KeyPressDelay)

	// Key up
	inp.Ki.Flags = KEYEVENTF_KEYUP
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
				pressKey(VK_LCONTROL)
			} else if time.Since(lastPing) >= PingInterval {
				pressKey(VK_LCONTROL)
				lastPing = time.Now()
			}
			time.Sleep(PollRateActive)
		} else {
			pressed = false
			time.Sleep(PollRateIdle)
		}
	}
}
