package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kkrav3ts/thefinals-autoping/internal/keyboard"
	"github.com/kkrav3ts/thefinals-autoping/internal/statistics"
)

func main() {
	fmt.Println("THE FINALS Auto-Ping Tool.")

	// PREDEFINED INPUTS
	aimKey := 0x02 // Virtual-Key Code for Right Mouse Button used as aiming key.
	PollRateActive := 10 * time.Millisecond
	PollRateIdle := 200 * time.Millisecond
	delaysCount := 500
	mean := 80.0
	stdDev := 5.0
	minVal := 70.0
	maxVal := 100.0

	// USER-BASED INPUT
	fmt.Printf("Press the key you want to use for ping.\n")
	pingKey := keyboard.DetectKeyPress(keyboard.KeyNames)
	fmt.Printf("Auto-ping enabled using [%s] key. Start aiming with right mouse button...\n", keyboard.KeyNames[pingKey])

	// Generate a pool of delays to cycle through
	delays := statistics.GenerateClickDelays(delaysCount, mean, stdDev, minVal, maxVal)
	fmt.Printf("Generated %v human-like delays to be used in the loop.\n", delaysCount)

	// Graceful shutdown on Ctrl+C
	fmt.Println("Close window or press Ctrl+C to exit")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	// Infinite loop for the process
	delayIndex := 0
	var nextPingTime time.Time
	for {
		if keyboard.IsKeyPressed(aimKey) {
			now := time.Now()
			if nextPingTime.IsZero() {
				// RMB just pressed — first ping fires after a random delay only
				nextPingTime = now.Add(delays[delayIndex])
				delayIndex = (delayIndex + 1) % len(delays)
			}
			if now.After(nextPingTime) {
				keyboard.PressKey(pingKey, 10*time.Millisecond)
				// Subsequent pings at 1 second + random delay
				nextPingTime = time.Now().Add(1*time.Second + delays[delayIndex])
				delayIndex = (delayIndex + 1) % len(delays)
			}
			time.Sleep(PollRateActive)
		} else {
			nextPingTime = time.Time{}
			time.Sleep(PollRateIdle)
		}
	}
}
