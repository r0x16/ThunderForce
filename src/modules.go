package src

import (
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/module"

	devices "github.com/r0x16/ThunderForce/src/devices/infraestructure/module"
)

func ProvideModules(bundle *drivers.ApplicationBundle) []domain.ApplicationModule {
	return []domain.ApplicationModule{
		&module.MainModule{Bundle: bundle},
		&devices.DevicesModule{Bundle: bundle},
	}
}
