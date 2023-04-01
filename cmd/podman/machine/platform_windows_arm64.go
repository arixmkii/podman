package machine

import (
	"fmt"

	"github.com/containers/podman/v4/pkg/machine"
)

func getQemuProvider() (machine.VirtProvider, error) {
	return nil, fmt.Errorf("unsupported virtualization provider: `%s`", machine.QemuVirt.String())
}
