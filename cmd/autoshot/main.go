package main

import (
	"fmt"

	"os"
	"os/signal"
	"os/exec"
	"syscall"
	"time"

	"github.com/kkrav3ts/thefinals-autoping/internal/keyboard"
	"github.com/kkrav3ts/thefinals-autoping/internal/statistics"
)

type autoShotConfig struct {
	userFiringKey int
	shotKey       int
	delays        []time.Duration
}

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func configureAutoShot(debug bool, promptForTrigger bool, defaultTriggerKey int, delaysCount int, mean float64, stdDev float64, minVal float64, maxVal float64) autoShotConfig {
	debugf := func(format string, args ...any) {
		if !debug {
			return
		}
		fmt.Printf("[DEBUG %s] %s\n", time.Now().Format("15:04:05.000"), fmt.Sprintf(format, args...))
	}

	fmt.Println("THE FINALS Auto-Shooting Tool.")
	userFiringKey := defaultTriggerKey
	if promptForTrigger {
		fmt.Printf("Press the key or mouse button you want to use as the trigger for auto-shooting.\n")
		debugf("Waiting for trigger key selection from %d configured inputs", len(keyboard.TriggerKeyNames))
		userFiringKey = keyboard.DetectKeyPress(keyboard.TriggerKeyNames)
		fmt.Printf("Auto-shoot trigger enabled using [%s] input.\n", keyboard.TriggerKeyNames[userFiringKey])
		debugf("Selected trigger key: vk=0x%X name=%s", userFiringKey, keyboard.TriggerKeyNames[userFiringKey])
		keyboard.WaitForKeyRelease(userFiringKey)
	} else {
		fmt.Printf("Auto-shoot trigger defaulted to [%s]. Press F1 to change it.\n", keyboard.TriggerKeyNames[userFiringKey])
		debugf("Using default trigger key: vk=0x%X name=%s", userFiringKey, keyboard.TriggerKeyNames[userFiringKey])
	}

	fmt.Printf("Press the key you want to use for shooting.\n")
	debugf("Waiting for shooting key selection from %d configured keys", len(keyboard.KeyNames))
	shotKey := keyboard.DetectKeyPress(keyboard.KeyNames)
	fmt.Printf("Auto-shooting enabled using [%s] key. Hold [%s] to simulate repeated clicks...\n", keyboard.KeyNames[shotKey], keyboard.TriggerKeyNames[userFiringKey])
	debugf("Selected shot key: vk=0x%X name=%s", shotKey, keyboard.KeyNames[shotKey])
	keyboard.WaitForKeyRelease(shotKey)
	debugf("Delay generation config: count=%d mean=%.2fms stdDev=%.2fms min=%.2fms max=%.2fms", delaysCount, mean, stdDev, minVal, maxVal)

	delays := statistics.GenerateClickDelays(delaysCount, mean, stdDev, minVal, maxVal)
	fmt.Printf("Generated %v human-like key presses to be used in the loop.\n", delaysCount)

	if debug && len(delays) > 0 {
		minDelay := delays[0]
		maxDelay := delays[0]
		var total time.Duration
		for _, delay := range delays {
			if delay < minDelay {
				minDelay = delay
			}
			if delay > maxDelay {
				maxDelay = delay
			}
			total += delay
		}
		debugf("Delay pool summary: min=%v max=%v avg=%v", minDelay, maxDelay, total/time.Duration(len(delays)))
		sampleCount := min(10, len(delays))
		debugf("Delay pool sample (%d/%d): %v", sampleCount, len(delays), delays[:sampleCount])
	}

	fmt.Println("Close window or press Ctrl+C to exit")
	fmt.Println("Press F1 to reconfigure trigger/shoot keys")
	fmt.Println("Press F13 to pause/resume auto-shooting")

	return autoShotConfig{
		userFiringKey: userFiringKey,
		shotKey:       shotKey,
		delays:        delays,
	}
}

func main() {
	// PREDEFINED INPUTS
	debug := false          // for debugging
	userFiringKey := 0x01   // Trigger key held by the user to activate auto-shooting.
	reconfigureKey := 0x70  // Virtual-Key Code for F1 used to restart setup.
	pauseToggleKey := 0x7C  // Virtual-Key Code for F13 used to pause/resume the script.
	delaysCount := 1000     // number of delays to generate
	mean := 60.0            // midpoint of delay cluster
	stdDev := 5.0           // standard deviation to create the delay cluster
	minVal := 50.0          // minimum delay
	maxVal := 80.0          // maximum delay

	debugf := func(format string, args ...any) {
		if !debug {
			return
		}
		fmt.Printf("[DEBUG %s] %s\n", time.Now().Format("15:04:05.000"), fmt.Sprintf(format, args...))
	}

	config := configureAutoShot(debug, false, userFiringKey, delaysCount, mean, stdDev, minVal, maxVal)
	userFiringKey = config.userFiringKey
	shotKey := config.shotKey
	delays := config.delays

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	// Infinite loop for the main process
	delayIndex := 0
	shotCount := 0
	paused := false
	f1WasPressed := false
	f13WasPressed := false
	triggerWasPressed := false
	lastShotStartedAt := time.Time{}
	for {
		f1Pressed := keyboard.IsKeyPressed(reconfigureKey)
		if f1Pressed && !f1WasPressed {
			debugf("Reconfiguration requested with F1")
			keyboard.WaitForKeyRelease(reconfigureKey)
			clearTerminal()
			config = configureAutoShot(debug, true, userFiringKey, delaysCount, mean, stdDev, minVal, maxVal)
			userFiringKey = config.userFiringKey
			shotKey = config.shotKey
			delays = config.delays
			delayIndex = 0
			shotCount = 0
			paused = false
			triggerWasPressed = false
			lastShotStartedAt = time.Time{}
			f1WasPressed = false
			f13WasPressed = false
			continue
		}
		f1WasPressed = f1Pressed

		f13Pressed := keyboard.IsKeyPressed(pauseToggleKey)
		if f13Pressed && !f13WasPressed {
			paused = !paused
			if paused {
				fmt.Println("Auto-shooting paused")
				debugf("Pause toggled on with F13")
			} else {
				fmt.Println("Auto-shooting resumed")
				debugf("Pause toggled off with F13")
			}
		}
		f13WasPressed = f13Pressed

		triggerPressed := keyboard.IsKeyPressed(userFiringKey)
		if triggerPressed != triggerWasPressed {
			if triggerPressed {
				debugf("Trigger key pressed")
			} else {
				debugf("Trigger key released after %d simulated shots", shotCount)
			}
			triggerWasPressed = triggerPressed
		}

		if paused {
			time.Sleep(10 * time.Millisecond)
		} else if triggerPressed {
			holdDelay := delays[delayIndex]

			// Press key with human-like key pressed time
			pressStartedAt := time.Now()
			keyboard.PressKey(shotKey, holdDelay)
			pressElapsed := time.Since(pressStartedAt)
			delayIndex = (delayIndex + 1) % len(delays)

			// Human-like Delay between key pressed
			betweenDelay := delays[delayIndex]
			sleepStartedAt := time.Now()
			time.Sleep(betweenDelay)
			sleepElapsed := time.Since(sleepStartedAt)
			shotCount++
			intervalSinceLastStart := time.Duration(0)
			if !lastShotStartedAt.IsZero() {
				intervalSinceLastStart = pressStartedAt.Sub(lastShotStartedAt)
			}
			lastShotStartedAt = pressStartedAt
			debugf(
				"Shot #%d hold=%v actualHold=%v gap=%v actualGap=%v sincePrevStart=%v",
				shotCount,
				holdDelay,
				pressElapsed,
				betweenDelay,
				sleepElapsed,
				intervalSinceLastStart,
			)
			delayIndex = (delayIndex + 1) % len(delays)
		} else {
			// Small polling delay to avoid excessive CPU usage when idle
			time.Sleep(10 * time.Millisecond)
		}
	}
}
