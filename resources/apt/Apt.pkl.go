// Code generated from Pkl module `mikea.declix.resources.Apt`. DO NOT EDIT.
package apt

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Apt struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Apt
func LoadFromPath(ctx context.Context, path string) (ret *Apt, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Apt
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Apt, error) {
	var ret Apt
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
