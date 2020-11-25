package provision

import (
	"testing"

	"docker-machine/drivers/fakedriver"
	"docker-machine/libmachine/auth"
	"docker-machine/libmachine/engine"
	"docker-machine/libmachine/provision/provisiontest"
	"docker-machine/libmachine/swarm"
)

func TestUbuntuSystemdCompatibleWithHost(t *testing.T) {
	info := &OsRelease{
		ID:        "ubuntu",
		VersionID: "15.04",
	}
	p := NewUbuntuSystemdProvisioner(nil)
	p.SetOsReleaseInfo(info)

	compatible := p.CompatibleWithHost()

	if !compatible {
		t.Fatalf("expected to be compatible with ubuntu 15.04")
	}

	info.VersionID = "14.04"

	compatible = p.CompatibleWithHost()

	if compatible {
		t.Fatalf("expected to NOT be compatible with ubuntu 14.04")
	}

}

func TestUbuntuSystemdDefaultStorageDriver(t *testing.T) {
	p := NewUbuntuSystemdProvisioner(&fakedriver.Driver{}).(*UbuntuSystemdProvisioner)
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "overlay2" {
		t.Fatal("Default storage driver should be overlay2")
	}
}
