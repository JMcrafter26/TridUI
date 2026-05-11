//go:build !windows

package main

func attachConsole() {
	// Unix-like systems inherently attach std i/o based on shell execution. No-op needed.
}
