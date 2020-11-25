package provision

import (
	"docker-machine/libmachine/drivers"
)

func init() {
	Register("OracleLinux", &RegisteredProvisioner{
		New: NewOracleLinuxProvisioner,
	})
}

func NewOracleLinuxProvisioner(d drivers.Driver) Provisioner {
	return &OracleLinuxProvisioner{
		NewRedHatProvisioner("ol", d),
	}
}

type OracleLinuxProvisioner struct {
	*RedHatProvisioner
}

func (provisioner *OracleLinuxProvisioner) String() string {
	return "ol"
}
