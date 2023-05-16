package tickledb

type Query struct {
	fields []*field
}

type field struct {
	fieldName string
	value     interface{}
}

func GetQuery() *Query {
	return &Query{
		fields: make([]*field, 0),
	}
}

func (q *Query) Fields(fields ...*field) *Query {
	q.fields = append(q.fields, fields...)
	return q
}

func FieldWithName(fieldName string) *field {
	return &field{
		fieldName: fieldName,
	}
}

func (f *field) WithValue(value interface{}) *field {
	f.value = value
	return f
}
