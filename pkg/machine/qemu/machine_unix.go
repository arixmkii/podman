//go:build unix
// +build unix

package qemu

import (
	"bytes"
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

func isProcessAlive(pid int) bool {
	err := unix.Kill(pid, syscall.Signal(0))
	if err == nil || err == unix.EPERM {
		return true
	}
	return false
}

func isQemuAlive(pid int, stderrBuf *bytes.Buffer) error {
	// check if qemu is still alive
	var status syscall.WaitStatus
	pid, err := syscall.Wait4(pid, &status, syscall.WNOHANG, nil)
	if err != nil {
		return fmt.Errorf("failed to read qemu process status: %w", err)
	}
	if pid > 0 {
		// child exited
		return fmt.Errorf("qemu exited unexpectedly with exit code %d, stderr: %s", status.ExitStatus(), stderrBuf.String())
	}
	return nil
}
