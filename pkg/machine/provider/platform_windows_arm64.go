package provider

import (
	"fmt"

	"github.com/containers/podman/v5/pkg/machine/define"
	"github.com/containers/podman/v5/pkg/machine/hyperv"
	"github.com/containers/podman/v5/pkg/machine/vmconfigs"
	"github.com/containers/podman/v5/pkg/machine/wsl"
)

func getQemuProvider() (vmconfigs.VMProvider, error) {
	return nil, fmt.Errorf("unsupported virtualization provider: `%s`", define.QemuVirt.String())
}

func isQemuInstalled() bool {
	return false
}

func GetAll() []vmconfigs.VMProvider {
	return []vmconfigs.VMProvider{
		new(wsl.WSLStubber),
		new(hyperv.HyperVStubber),
	}
}

// SupportedProviders returns the providers that are supported on the host operating system
func SupportedProviders() []define.VMType {
	return []define.VMType{define.HyperVVirt, define.WSLVirt}
}
