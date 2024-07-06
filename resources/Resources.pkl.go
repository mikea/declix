// Code generated from Pkl module `mikea.declix.Resources`. DO NOT EDIT.
package resources

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Resources struct {
	Resources []Resource `pkl:"resources"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Resources
func LoadFromPath(ctx context.Context, path string) (ret *Resources, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Resources
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Resources, error) {
	var ret Resources
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
