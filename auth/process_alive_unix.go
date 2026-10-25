//go:build !windows

package auth

import (
	"errors"
	"os"
	"syscall"
)

// processAliveOS reports whether pid names a running process. Signal 0 is
// delivered to nothing; the kernel answers only whether it could have been.
//
// EPERM reads as alive: a live process owned by another user refuses the signal,
// and reading that as death would take the lock from a holder still working. A
// zombie accepts signal 0 for ever, which breakStaleLock leaves to the mtime gate.
func processAliveOS(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil || errors.Is(err, syscall.EPERM)
}
