package impl

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"

	"github.com/pterm/pterm"

	_ "mikea/declix/resources/apt"
	_ "mikea/declix/resources/dpkg"
	_ "mikea/declix/resources/filesystem"
	_ "mikea/declix/resources/systemd"
	_ "mikea/declix/resources/users"
)

func CreateResources(pkl []resources.Resource) []interfaces.Resource {
	resources := make([]interfaces.Resource, len(pkl))
	for i, res := range pkl {
		resources[i] = CreateResource(res)
	}
	return resources
}

func CreateResource(r resources.Resource) interfaces.Resource {
	if res, ok := r.(interfaces.Resource); ok {
		return res
	} else {
		panic(fmt.Sprintf("%#v does not implement interfaces.Resource", r))
	}
}

func DetermineStates(resources []interfaces.Resource, executor interfaces.CommandExecutor, progress pterm.ProgressbarPrinter) ([]interfaces.State, []interfaces.State, []error) {
	states := make([]interfaces.State, len(resources))
	expected := make([]interfaces.State, len(resources))
	errors := make([]error, len(resources))
	for i, res := range resources {
		progress.UpdateTitle(res.GetId())
		progress.Increment()

		expected[i], errors[i] = res.ExpectedState()
		if errors[i] != nil {
			continue
		}
		states[i], errors[i] = res.DetermineState(executor)
	}

	return states, expected, errors
}
