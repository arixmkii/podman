package qemu

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/containers/podman/v4/pkg/machine"
)

func isProcessAlive(pid int) bool {
	if checkProcessStatus("process", pid, nil) == nil {
		return true
	}
	return false
}

func checkProcessStatus(processHint string, pid int, stderrBuf *bytes.Buffer) error {
	active, exitCode := machine.GetProcessState(pid)
	if !active {
		if stderrBuf != nil {
			return fmt.Errorf("%s exited unexpectedly, exit code: %d stderr: %s", processHint, exitCode, stderrBuf.String())
		} else {
			return fmt.Errorf("%s exited unexpectedly, exit code: %d", processHint, exitCode)
		}
	}
	return nil
}

func forwardPipeArgs(name string, destPath string, identityPath string, user string) []string {
	args := []string{}
	args = append(args, []string{"-forward-sock", fmt.Sprintf("npipe:////./pipe/%s", toPipeName(name))}...)
	args = append(args, []string{"-forward-dest", destPath}...)
	args = append(args, []string{"-forward-user", user}...)
	args = append(args, []string{"-forward-identity", identityPath}...)
	return args
}

func podmanPipe(name string) *machine.VMFile {
	return &machine.VMFile{Path: `\\.\pipe\` + toPipeName(name)}
}

func toPipeName(name string) string {
	if !strings.HasPrefix(name, "qemu") {
		if !strings.HasPrefix(name, "podman") {
			name = "podman-" + name
		}
		name = "qemu-" + name
	}
	return name
}

func useFdVLan() bool {
	return false
}

func pathsFromVolume(volume string) []string {
	paths := strings.SplitN(volume, ":", 3)
	driveLetterMatcher := regexp.MustCompile(`^(?:\\\\[.?]\\)?[a-zA-Z]$`)
	if len(paths) > 1 && driveLetterMatcher.MatchString(paths[0]) {
		paths = strings.SplitN(volume, ":", 4)
		paths = append([]string{paths[0] + ":" + paths[1]}, paths[2:]...)
	}
	return paths
}

func extractTargetPath(paths []string) string {
	if len(paths) > 1 {
		return paths[1]
	}
	target := strings.ReplaceAll(paths[0], "\\", "/")
	target = strings.ReplaceAll(target, ":", "/")
	if strings.HasPrefix(target, "//./") || strings.HasPrefix(target, "//?/") {
		target = target[4:]
	}
	dedup := regexp.MustCompile(`//+`)
	return dedup.ReplaceAllLiteralString("/"+target, "/")
}
