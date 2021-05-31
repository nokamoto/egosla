package test

import (
	"fmt"

	"go.uber.org/zap"
)

// Scenario represents a test.
type Scenario struct {
	// Name is the test name.
	Name string
	// Run executes the test with the current state and returns an updated state.
	Run func(state State, logger *zap.Logger) (State, error)
}

// Scenarios represents ordered scenarios.
type Scenarios []Scenario

// Run executes all scenarios. It halts if one of the scenarios fails.
func (xs Scenarios) Run(logger *zap.Logger) {
	size := len(xs)
	logger.Info("begin", zap.Int("#", size))
	st := make(State)
	for i, x := range xs {
		sl := logger.With(zap.String("scenario", x.Name), zap.Any("state", st))
		s, err := x.Run(st, sl)
		if err != nil {
			sl.Fatal("fatal", zap.Error(err))
		}
		st = s
		sl.Info("ok", zap.String("#", fmt.Sprintf("%d/%d", i+1, size)))
	}
	logger.Info("done", zap.Int("#", size))
}
