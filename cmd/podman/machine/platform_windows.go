package machine

import (
	"os"

	"github.com/containers/podman/v4/pkg/machine"
	"github.com/containers/podman/v4/pkg/machine/wsl"
	"github.com/containers/podman/v4/pkg/machine/qemu"
)

func GetSystemDefaultProvider() machine.Provider {
	switch os.Getenv("CONTAINERS_MACHINE_PROVIDER") {
	case "WSL":
		return wsl.GetWSLProvider()
	case "QEMU":
		return qemu.GetVirtualizationProvider()
	default:
		return wsl.GetWSLProvider()
	}
}
