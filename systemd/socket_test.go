package systemd_test

import (
	"mikea/declix/interfaces"
	. "mikea/declix/systemd"
)

var _ interfaces.Resource = &SocketImpl{}
