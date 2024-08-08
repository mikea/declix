package impl

import (
	"fmt"

	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

func HaveStyleStrings(expected string, current string, action string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(r *ResourceState) (bool, error) {
		if r.Error != nil {
			return false, r.Error
		}
		if expected != r.Expected.GetStyledString() {
			return false, fmt.Errorf(
				"expected state doesn't match\n    expected styled string: %q\n    actual styled string: %q",
				expected, r.Expected.GetStyledString())
		}
		if current != r.Current.GetStyledString() {
			return false, fmt.Errorf(
				"current state doesn't match\n    current styled string: %q\n    actual styled string: %q",
				current, r.Current.GetStyledString())
		}
		if action != r.Action.GetStyledString(r.Resource) {
			return false, fmt.Errorf(
				"action doesn't match\n    action styled string: %q\n    actual styled string: %q",
				action, r.Action.GetStyledString(r.Resource))
		}
		return true, nil
	}).WithTemplate(
		"Expected:\n{{.FormattedActual}}\n{{.To}} match\n{{format .Data 1}}",
		[]string{expected, current, action})
}
