package systemd_test

import (
	"mikea/declix/interfaces"
	. "mikea/declix/resources/systemd"
)

var _ interfaces.Resource = &ServiceImpl{}
