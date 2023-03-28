package application

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type CommandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger logrus.FieldLogger
}

func NewCommandLoggingDecorator[C any](base CommandHandler[C], logger logrus.FieldLogger) *CommandLoggingDecorator[C] {
	return &CommandLoggingDecorator[C]{
		base:   base,
		logger: logger,
	}
}

func (d *CommandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	logger := d.logger.WithFields(logrus.Fields{
		"command":      generateActionName(cmd),
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing command")

	defer func() {
		if err != nil {
			logger.WithError(err).Error("Failed to execute command")
		} else {
			logger.Info("Command executed successfully")
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type QueryLoggingDecorator[Q any, R any] struct {
	base   QueryHandler[Q, R]
	logger logrus.FieldLogger
}

func NewQueryLoggingDecorator[Q any, R any](base QueryHandler[Q, R], logger logrus.FieldLogger) *QueryLoggingDecorator[Q, R] {
	return &QueryLoggingDecorator[Q, R]{
		base:   base,
		logger: logger,
	}
}

func (d *QueryLoggingDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	logger := d.logger.WithFields(logrus.Fields{
		"query":      generateActionName(query),
		"query_body": fmt.Sprintf("%#v", query),
	})

	logger.Debug("Executing query")

	defer func() {
		if err != nil {
			logger.WithError(err).Error("Failed to execute query")
		} else {
			logger.Info("Query executed successfully")
		}
	}()

	return d.base.Handle(ctx, query)
}
