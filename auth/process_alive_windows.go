//go:build windows

package auth

import (
	"errors"
	"syscall"
)

// processAliveOS reports whether pid names a process that is still running.
//
// A Windows process object outlives the process while any handle to it stays
// open, so os.FindProcess succeeds for a pid that has already exited. A wait with
// a zero timeout returns at once: WAIT_OBJECT_0 means exited, anything else means
// still running. GetExitCodeProcess against STILL_ACTIVE is not used, because 259
// is also a legal exit code.
//
// Only SYNCHRONIZE is requested, the one access right the wait needs, and every
// refusal except "no such process" reads as alive: ERROR_ACCESS_DENIED is what a
// live process owned by another user answers.
func processAliveOS(pid int) bool {
	// Declared here because syscall does not export it on Windows, and go.mod
	// has no require block for golang.org/x/sys.
	const errorInvalidParameter = syscall.Errno(87) // no such process object
	handle, err := syscall.OpenProcess(syscall.SYNCHRONIZE, false, uint32(pid))
	if err != nil {
		return !errors.Is(err, errorInvalidParameter)
	}
	defer func() { _ = syscall.CloseHandle(handle) }()
	state, err := syscall.WaitForSingleObject(handle, 0)
	if err != nil {
		// A lock is never taken over on a guess; the mtime gate still recovers it.
		return true
	}
	return state != syscall.WAIT_OBJECT_0
}
