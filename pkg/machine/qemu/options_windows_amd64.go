package qemu

var (
	QemuCommand = "qemu-system-x86_64"
)

func (v *MachineVM) addArchOptions() []string {
	// TODO add HAXM support. Via env var config (as this is used only during init anyway)?
	opts := []string{"-machine", "q35,accel=whpx:tcg"} // doesn't work on win qemu? , "-cpu", "host"
	return opts
}

func (v *MachineVM) prepare() error {
	return nil
}

func (v *MachineVM) archRemovalFiles() []string {
	return []string{}
}
