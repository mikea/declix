package resources

import "github.com/pterm/pterm"

func (m *Missing) GetStyledString() string {
	return pterm.FgRed.Sprint("missing")
}
