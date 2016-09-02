package trace

import (
	"context"

	"github.com/uber-go/zap"
)

type Tracer interface {
	Ok(event string, fields ...Field) error
	Wrap(err error, event string, fields ...Field) error
}

type tracer struct {
	logger zap.Logger
}

func (t tracer) Ok(event string, fields ...Field) error {
	zapFields := ZapFields(nil, fields...)
	t.logger.Info(event, zapFields...)
	return nil
}

func (t tracer) Wrap(err error, event string, fields ...Field) error {
	zapFields := ZapFields(err, fields...)
	t.logger.Info(event, zapFields...)
	return err
}

func FromContext(ctx context.Context) Tracer {
	logger := zap.New(zap.NewJSONEncoder(zap.NoTime()))

	return tracer{
		logger: logger,
	}
}
