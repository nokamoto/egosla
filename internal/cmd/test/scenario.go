package test

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Scenario represents a test.
type Scenario struct {
	// Name is the test name.
	Name string
	// Run executes the test with the current state and returns an updated state.
	Run func(state State, logger *zap.Logger) error
}

// Scenarios represents ordered scenarios.
type Scenarios []Scenario

func (xs Scenarios) run(logger *zap.Logger, setup, teardown *Scenario) {
	size := len(xs)
	logger.Info("begin", zap.Int("#", size))
	st := make(State)

	var failed = false

	if setup != nil {
		logger.Info(setup.Name)
		err := setup.Run(st, logger)
		if err != nil {
			failed = true
			logger.Error(setup.Name, zap.Error(err))
		}
	}

	for i, x := range xs {
		if failed {
			break
		}
		sl := logger.With(zap.String("scenario", x.Name), zap.Any("state", st))
		err := x.Run(st, sl)
		if err != nil {
			failed = true
			sl.Error("failed", zap.Error(err))
			continue
		}
		sl.Info("ok", zap.String("#", fmt.Sprintf("%d/%d", i+1, size)))
	}

	if !failed {
		logger.Info("done", zap.Int("#", size))
	}

	if teardown != nil {
		logger.Info(teardown.Name)
		err := teardown.Run(st, logger)
		if err != nil {
			failed = true
			logger.Error(teardown.Name, zap.Error(err))
		}
	}

	if failed {
		os.Exit(1)
	}
}
