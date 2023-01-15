//go:build (amd64 && !remote) || (arm64 && !remote)
// +build amd64,!remote arm64,!remote

package system

import (
	"github.com/containers/common/pkg/config"
	cmdMach "github.com/containers/podman/v4/cmd/podman/machine"
)

func resetMachine() error {
	cfg, err := config.ReadCustomConfig()
	if err != nil {
		return err
	}

	provider := cmdMach.GetSystemProvider(cfg.Machine.Provider)
	return provider.RemoveAndCleanMachines()
}
