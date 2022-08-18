package qemu

import (
	"bytes"
	"fmt"
	"os/exec"
)

func isProcessAlive(pid int) bool {
	// TODO do we want /dev/null it?
	cmd := exec.Command("cmd", "/C", fmt.Sprintf("tasklist /FO csv /NH | findstr /L ,\"%d\",", pid))
	err := cmd.Run()
	if err == nil {
		return true
	}
	return false
}

func isQemuAlive(pid int, stderrBuf *bytes.Buffer) error {
	return nil
}
