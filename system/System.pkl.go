// Code generated from Pkl module `mikea.declix.System`. DO NOT EDIT.
package system

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type System struct {
	Ssh *SshConfig `pkl:"ssh"`

	Packages []*Package `pkl:"packages"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a System
func LoadFromPath(ctx context.Context, path string) (ret *System, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a System
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*System, error) {
	var ret System
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
