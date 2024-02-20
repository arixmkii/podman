package provider

import (
	"fmt"

	"github.com/containers/podman/v5/pkg/machine/define"
	"github.com/containers/podman/v5/pkg/machine/vmconfigs"
)

func getQemuProvider() (vmconfigs.VMProvider, error) {
	return nil, fmt.Errorf("unsupported virtualization provider: `%s`", define.QemuVirt.String())
}
