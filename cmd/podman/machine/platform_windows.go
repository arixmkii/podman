package machine

import (
	"fmt"
	"os"

	"github.com/containers/common/pkg/config"
	"github.com/containers/podman/v4/pkg/machine"
	"github.com/containers/podman/v4/pkg/machine/hyperv"
	"github.com/containers/podman/v4/pkg/machine/qemu"
	"github.com/containers/podman/v4/pkg/machine/wsl"
)

func GetSystemProvider(vmType *machine.VMType) (machine.VirtProvider, error) {
	resolvedVMType := machine.QemuVirt
	if vmType != nil {
		resolvedVMType = *vmType
	} else {
		cfg, err := config.Default()
		if err != nil {
			return nil, err
		}
		provider := cfg.Machine.Provider
		if providerOverride, found := os.LookupEnv("CONTAINERS_MACHINE_PROVIDER"); found {
			provider = providerOverride
		}
		resolvedVMType, err = machine.ParseVMType(provider, machine.WSLVirt)
		if err != nil {
			return nil, err
		}
	}

	switch resolvedVMType {
	case machine.WSLVirt:
		return wsl.GetWSLProvider(), nil
	case machine.HyperVVirt:
		return hyperv.GetVirtualizationProvider(), nil
	case machine.QemuVirt:
		return qemu.GetVirtualizationProvider(), nil
	default:
		return nil, fmt.Errorf("unsupported virtualization provider: `%s`", resolvedVMType.String())
	}
}
