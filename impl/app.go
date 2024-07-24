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
	Resources []interfaces.Resource
	States    []interfaces.State
	Expected  []interfaces.State
	Errors    []error
	Actions   []interfaces.Action
	Executor  interfaces.CommandExecutor
}

func (app *App) Dispose() {
	if app.Executor != nil {
		app.Executor.Close()
	}
}

func (app *App) HasErrors() bool {
	for _, err := range app.Errors {
		if err != nil {
			return true
		}
	}
	return false
}

func (app *App) ApplyActions() error {
	totalActions := 0
	for _, a := range app.Actions {
		if a != nil {
			totalActions += 1
		}
	}

	progress, err := pterm.DefaultProgressbar.WithTotal(totalActions).WithRemoveWhenDone(true).Start()
	if err != nil {
		return err
	}

	for i, action := range app.Actions {
		if action == nil {
			continue
		}
		progress.UpdateTitle(action.StyledString(app.Resources[i]))
		progress.Increment()

		err = app.Resources[i].RunAction(app.Executor, app.Actions[i], app.States[i], app.Expected[i])
		if err != nil {
			pterm.Println(pterm.BgRed.Sprint("E ", app.Actions[i].StyledString(app.Resources[i])), err)
		} else {
			pterm.Println(pterm.FgGreen.Sprint("\u2713"), app.Actions[i].StyledString(app.Resources[i]))
		}
	}
	progress.Stop()
	return nil
}

func (app *App) DetermineActions() error {
	app.Actions = make([]interfaces.Action, len(app.Resources))
	for i, res := range app.Resources {
		if app.Errors[i] != nil {
			continue
		}
		app.Actions[i], app.Errors[i] = res.DetermineAction(app.States[i], app.Expected[i])
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

	app.States, app.Expected, app.Errors = DetermineStates(app.Resources, app.Executor, *progress)
	progress.Stop()
	return nil
}

func (app *App) LoadResources(fileName string) error {
	resourcesPkl, err := resources.LoadFromPath(context.Background(), fileName)
	if err != nil {
		return fmt.Errorf("loading resources file: %w", err)
	}

	app.Resources = CreateResources(resourcesPkl.Resources)
	app.Errors = make([]error, len(app.Resources))
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
