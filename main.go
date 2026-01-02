// Package main provides an auto-ping utility for THE FINALS.
// It automatically presses the ping key while aiming (holding right mouse button).
package main

import (
	"fmt"
	"syscall"
	"time"
)

var (
	Version          = "dev"
	user32           = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	keybd_event      = user32.NewProc("keybd_event")
)

// Virtual-Key Codes
// https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
const (
	VK_RBUTTON      = 0x02 // Right mouse button
	VK_LCONTROL     = 0xA2 // Left Control key
	KEYEVENTF_KEYUP = 0x0002
)

// Timing configuration
const (
	PingInterval   = 1 * time.Second
	PollRateActive = 100 * time.Millisecond // Polling when aiming
	PollRateIdle   = 200 * time.Millisecond // Slow polling when idle
)

func isKeyPressed(vk int) bool {
	ret, _, _ := getAsyncKeyState.Call(uintptr(vk))
	return ret&0x8000 != 0
}

func pressKey(vk int) {
	keybd_event.Call(uintptr(vk), 0, 0, 0)
	keybd_event.Call(uintptr(vk), 0, uintptr(KEYEVENTF_KEYUP), 0)
}

func main() {
	fmt.Printf("THE FINALS Auto-Ping %s\n", Version)
	fmt.Println("Close window to exit")

	var nextPingTime time.Time

	for {
		if isKeyPressed(VK_RBUTTON) {
			now := time.Now()
			if nextPingTime.IsZero() || now.After(nextPingTime) {
				pressKey(VK_LCONTROL)
				nextPingTime = now.Add(PingInterval) // Calculate next ping time
			}
			time.Sleep(PollRateActive) // Polling when aiming
		} else {
			nextPingTime = time.Time{} // Reset when not pressing
			time.Sleep(PollRateIdle)   // Slow polling when idle
		}
	}
}
