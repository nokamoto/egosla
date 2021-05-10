package main

import (
	"go.uber.org/zap"
)

type state map[string]string

type scenario struct {
	name string
	run  func(state state, logger *zap.Logger) (state, error)
}

type scenarios []scenario

func (xs scenarios) run(logger *zap.Logger) {
	logger.Info("begin")
	st := make(state)
	for _, x := range xs {
		sl := logger.With(zap.String("scenario", x.name), zap.Any("state", st))
		sl.Info("run")
		s, err := x.run(st, sl)
		if err != nil {
			sl.Fatal("fatal", zap.Error(err))
		}
		st = s
		sl.Info("ok")
	}
	logger.Info("done")
}
