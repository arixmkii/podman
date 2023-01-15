//go:build (amd64 && !windows) || (arm64 && !windows)
// +build amd64,!windows arm64,!windows

package machine

import (
	"os"
	"strings"

	"github.com/containers/podman/v4/pkg/machine"
	"github.com/containers/podman/v4/pkg/machine/qemu"
	"github.com/sirupsen/logrus"
)

func GetSystemProvider(name string) machine.Provider {
	provider := name
	providerOverride, found := os.LookupEnv("CONTAINERS_MACHINE_PROVIDER")
	if found {
		provider = providerOverride
	}
	provider = strings.ToUpper(strings.TrimSpace(provider))
	switch provider {
	case "QEMU":
		return qemu.GetVirtualizationProvider()
	default:
		if len(provider) > 0 {
			logrus.Warnf("Unuspported provider specified %q. Will use default provider.", provider)
		}
		return qemu.GetVirtualizationProvider()
	}
}
