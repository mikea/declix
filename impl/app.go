package impl

import (
	"context"
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"
	"mikea/declix/target"

	"github.com/pterm/pterm"
)

type App struct {
	Target    *target.SshConfig
	Resources []*ResourceState
	Executor  interfaces.CommandExecutor
}

type ResourceState struct {
	Resource interfaces.Resource
	Error    error

	Current  interfaces.State
	Expected interfaces.State
	Action   interfaces.Action
}

func (app *App) Dispose() {
	if app.Executor != nil {
		app.Executor.Close()
	}
}

func (app *App) HasErrors() bool {
	for _, r := range app.Resources {
		if r.Error != nil {
			return true
		}
	}
	return false
}

func (app *App) HasActions() bool {
	for _, r := range app.Resources {
		if r.Action != nil {
			return true
		}
	}
	return false
}

func (app *App) ApplyActions() error {
	totalActions := 0
	for _, r := range app.Resources {
		if r.Action != nil {
			totalActions += 1
		}
	}

	progress, err := pterm.DefaultProgressbar.WithTotal(totalActions).WithRemoveWhenDone(true).Start()
	if err != nil {
		return err
	}

	for _, r := range app.Resources {
		if r.Action == nil {
			continue
		}
		progress.UpdateTitle(r.Action.StyledString(r.Resource))
		progress.Increment()

		r.Error = r.Resource.RunAction(app.Executor, r.Action, r.Current, r.Expected)
		if err != nil {
			pterm.Println(pterm.BgRed.Sprint("E ", r.Action.StyledString(r.Resource)), err)
		} else {
			pterm.Println(pterm.FgGreen.Sprint("\u2713"), r.Action.StyledString(r.Resource))
		}
	}
	progress.Stop()
	return nil
}

func (app *App) DetermineActions() error {
	for _, r := range app.Resources {
		if r.Error != nil {
			continue
		}
		r.Action, r.Error = r.Resource.DetermineAction(r.Current, r.Expected)
	}

	return nil
}

func (app *App) DetermineStates() error {
	progress, err := pterm.DefaultProgressbar.WithTotal(len(app.Resources)).WithRemoveWhenDone(true).Start()
	if err != nil {
		return err
	}

	progress.UpdateTitle("Connecting...")
	app.Executor, err = SshExecutor(*app.Target)
	if err != nil {
		return err
	}

	for _, r := range app.Resources {
		progress.UpdateTitle(r.Resource.GetId())
		progress.Increment()

		r.Expected, r.Error = r.Resource.ExpectedState()
		if r.Error != nil {
			continue
		}
		r.Current, r.Error = r.Resource.DetermineState(app.Executor)
	}

	progress.Stop()
	return nil
}

func (app *App) LoadResources(fileName string) error {
	resourcesPkl, err := resources.LoadFromPath(context.Background(), fileName)
	if err != nil {
		return fmt.Errorf("loading resources file: %w", err)
	}

	app.Resources = make([]*ResourceState, len(resourcesPkl.Resources))
	for i, pkl := range resourcesPkl.Resources {
		app.Resources[i] = &ResourceState{
			Resource: CreateResource(pkl),
		}
	}

	return nil
}

func (app *App) LoadTargetFromText(text string) error {
	targetPkl, err := target.LoadFromText(context.Background(), text)
	if err != nil {
		return fmt.Errorf("loading target file: %w", err)
	}
	app.Target = targetPkl.Target
	pterm.Printfln("Target: %v", app.Target)

	return nil

}

func (app *App) LoadTarget(fileName string) error {
	targetPkl, err := target.LoadFromPath(context.Background(), fileName)
	if err != nil {
		return fmt.Errorf("loading target file: %w", err)
	}
	app.Target = targetPkl.Target
	pterm.Printfln("Target: %v", app.Target)

	return nil
}

func (app *App) PrintErrors() {
	pterm.Println(pterm.FgRed.Sprint("Errors:"))
	for _, r := range app.Resources {
		if r.Error != nil {
			pterm.Println(pterm.BgRed.Sprint(r.Resource.GetId()), pterm.FgRed.Sprint(r.Error))
		}
	}
}
