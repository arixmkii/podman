package provider

import (
	"github.com/containers/podman/v5/pkg/machine/qemu"
	"github.com/containers/podman/v5/pkg/machine/vmconfigs"
)

func getQemuProvider() (vmconfigs.VMProvider, error) {
	return new(qemu.QEMUStubber), nil
}
