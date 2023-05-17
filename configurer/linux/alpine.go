package linux

import (
	"github.com/k0sproject/k0sctl/configurer"
	"github.com/k0sproject/rig"
	"github.com/k0sproject/rig/os/linux"
	"github.com/k0sproject/rig/os/registry"
)

// BaseLinux for tricking go interfaces
type BaseLinux struct {
	configurer.Linux
}

// Alpine provides OS support for Alpine Linux
type Alpine struct {
	linux.Alpine
	BaseLinux
}

func init() {
	registry.RegisterOSModule(
		func(os rig.OSVersion) bool {
			return os.ID == "alpine"
		},
		func() interface{} {
			linuxType := &Alpine{}
			linuxType.PathFuncs = interface{}(linuxType).(configurer.PathFuncs)
			return linuxType
		},
	)
}
