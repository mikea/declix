package impl

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"
	"mikea/declix/resources/apt"
	"mikea/declix/resources/filesystem"

	"github.com/pterm/pterm"
)

func CreateResources(pkl []resources.Resource) []interfaces.Resource {
	resources := make([]interfaces.Resource, len(pkl))
	for i, res := range pkl {
		resources[i] = CreateResource(res)
	}
	return resources
}

func CreateResource(expected resources.Resource) interfaces.Resource {
	switch v := expected.(type) {
	case filesystem.File:
		return filesystem.New(v)
	case apt.Package:
		return apt.New(v)
	default:
		panic(fmt.Sprintf("unexpected pkl.Resource: %#v", v))
	}
}

func DetermineStatuses(resources []interfaces.Resource, executor interfaces.CommandExcutor, progress pterm.ProgressbarPrinter) ([]interfaces.Status, []error) {
	statuses := make([]interfaces.Status, len(resources))
	errors := make([]error, len(resources))
	for i, res := range resources {
		progress.UpdateTitle(res.Id())
		progress.Increment()
		statuses[i], errors[i] = res.DetermineStatus(executor)
	}

	return statuses, errors
}
