package target

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

func LoadFromText(ctx context.Context, text string) (ret *Target, err error) {
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
	ret, err = Load(ctx, evaluator, pkl.TextSource(text))
	return ret, err
}
