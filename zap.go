package trace

import "github.com/uber-go/zap"

func ZapFields(err error, fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	if err != nil {
		zapFields = make([]zap.Field, len(fields)+1)
		zapFields[len(fields)] = zap.Error(err)
	}
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
	return zapFields
}
