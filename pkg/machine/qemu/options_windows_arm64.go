package qemu

var (
	QemuCommand = "qemu-system-aarch64"
)

func (v *MachineVM) addArchOptions() []string {
	// stub to fix compilation issues
	opts := []string{}
	return opts
}

func (v *MachineVM) prepare() error {
	return nil
}

func (v *MachineVM) archRemovalFiles() []string {
	return []string{}
}
