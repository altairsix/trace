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
	zapFields := make([]zap.Field, len(fields))
	for index, field := range fields {
		switch field.Type {
		case TypeFloat:
			zapFields[index] = zap.Float64(field.Key, field.Float)
		case TypeInt:
			zapFields[index] = zap.Int64(field.Key, field.Int)
		case TypeObject:
			zapFields[index] = zap.Object(field.Key, field.Object)
		case TypeString:
			zapFields[index] = zap.String(field.Key, field.String)
		}
	}
	t.logger.Info(event, zapFields...)
	return nil
}

func (t tracer) Wrap(err error, event string, fields ...Field) error {
	zapFields := make([]zap.Field, len(fields)+1)
	for index, field := range fields {
		switch field.Type {
		case TypeFloat:
			zapFields[index] = zap.Float64(field.Key, field.Float)
		case TypeInt:
			zapFields[index] = zap.Int64(field.Key, field.Int)
		case TypeObject:
			zapFields[index] = zap.Object(field.Key, field.Object)
		case TypeString:
			zapFields[index] = zap.String(field.Key, field.String)
		}
	}
	zapFields[len(fields)] = zap.Error(err)
	t.logger.Info(event, zapFields...)
	return err
}

func FromContext(ctx context.Context) Tracer {
	logger := zap.New(zap.NewJSONEncoder(zap.NoTime()))

	return tracer{
		logger: logger,
	}
}
