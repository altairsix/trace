package trace

type FieldType int

const (
	TypeString FieldType = 1
	TypeInt    FieldType = 2
	TypeFloat  FieldType = 3
	TypeObject FieldType = 4
)

type Field struct {
	Key    string
	Type   FieldType
	String string
	Int    int64
	Float  float64
	Object interface{}
}

func String(key, value string) Field {
	return Field{
		Key:    key,
		Type:   TypeString,
		String: value,
	}
}

func Int(key string, value int64) Field {
	return Field{
		Key:  key,
		Type: TypeInt,
		Int:  value,
	}
}

func Float(key string, value float64) Field {
	return Field{
		Key:   key,
		Type:  TypeFloat,
		Float: value,
	}
}

func Object(key string, value interface{}) Field {
	return Field{
		Key:    key,
		Type:   TypeObject,
		Object: value,
	}
}
