package resources

import (
	"mikea/declix/interfaces"

	"github.com/pterm/pterm"
)

func (m *Missing) StyledString(r interfaces.Resource) string {
	return pterm.FgRed.Sprint("missing")
}
