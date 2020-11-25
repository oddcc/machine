package provision

import (
	"testing"

	"docker-machine/drivers/fakedriver"
	"docker-machine/libmachine/auth"
	"docker-machine/libmachine/engine"
	"docker-machine/libmachine/provision/provisiontest"
	"docker-machine/libmachine/swarm"
)

func TestArchDefaultStorageDriver(t *testing.T) {
	p := NewArchProvisioner(&fakedriver.Driver{}).(*ArchProvisioner)
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "overlay2" {
		t.Fatal("Default storage driver should be overlay2")
	}
}
