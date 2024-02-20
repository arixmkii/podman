package provider

import (
	"github.com/containers/podman/v5/pkg/machine/define"
	"github.com/containers/podman/v5/pkg/machine/hyperv"
	"github.com/containers/podman/v5/pkg/machine/qemu"
	"github.com/containers/podman/v5/pkg/machine/vmconfigs"
	"github.com/containers/podman/v5/pkg/machine/wsl"
)

func getQemuProvider() (vmconfigs.VMProvider, error) {
	return new(qemu.QEMUStubber), nil
}

func isQemuInstalled() bool {
	return true
}

func GetAll() []vmconfigs.VMProvider {
	return []vmconfigs.VMProvider{
		new(wsl.WSLStubber),
		new(hyperv.HyperVStubber),
		new(qemu.QEMUStubber),
	}
}

// SupportedProviders returns the providers that are supported on the host operating system
func SupportedProviders() []define.VMType {
	return []define.VMType{define.HyperVVirt, define.WSLVirt, define.QemuVirt}
}
