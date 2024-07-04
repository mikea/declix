package impl

import (
	"fmt"
	"mikea/declix/impl/apt_package"
	"mikea/declix/impl/file"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"

	"github.com/pterm/pterm"
)

func CreateResources(pkl []pkl.Resource) []interfaces.Resource {
	resources := make([]interfaces.Resource, len(pkl))
	for i, res := range pkl {
		resources[i] = CreateResource(res)
	}
	return resources
}

func CreateResource(expected pkl.Resource) interfaces.Resource {
	switch v := expected.(type) {
	case pkl.File:
		return file.New(v)
	case pkl.Package:
		return apt_package.New(v)
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
