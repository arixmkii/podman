package qemu

import (
	"fmt"
	"os"
	"strings"
)

var (
	QemuCommand = "qemu-system-x86_64"
)

func (v *MachineVM) addArchOptions() []string {
	var accel string
	switch strings.ToLower(os.Getenv("CONTAINERS_QEMU_ACCEL")) {
	case "whpx":
		accel = "whpx"
	case "hax":
		accel = "hax"
	default:
		accel = "whpx"
	}
	opts := []string{"-machine", fmt.Sprintf("q35,accel=%s:tcg", accel), "-cpu", "Skylake-Client"}
	return opts
}

func (v *MachineVM) prepare() error {
	return nil
}

func (v *MachineVM) archRemovalFiles() []string {
	return []string{}
}
