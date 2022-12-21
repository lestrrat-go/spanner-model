package spannermodel

const (
	MinWidth = 1
	MaxWidth = 2621440
)

type ColumnWidther interface {
	Width() int64
}

type ColumnType interface {
	Name() string
}

type BoolType struct{}

func (BoolType) Name() string {
	return `BOOL`
}

type BytesType struct {
	width int64
}

func (t *BytesType) Width() int64 {
	return t.width
}
func (t BytesType) Name() string {
	return `BYTES`
}

type DateType struct{}

func (DateType) Name() string {
	return `DATE`
}

type Float64Type struct{}

func (Float64Type) Name() string {
	return `FLOAT64`
}

type Int64Type struct{}

func (Int64Type) Name() string {
	return `INT64`
}

type JSONType struct{}

func (JSONType) Name() string {
	return `JSON`
}

type NumericType struct{}

func (NumericType) Name() string {
	return `NUMERIC`
}

type StringType struct {
	width int64
}

func (t *StringType) Width() int64 {
	return t.width
}

func (t StringType) Name() string {
	return `STRING`
}

type TimestampType struct{}

func (TimestampType) Name() string {
	return `TIMESTAMP`
}

func BoolColumn(name string) Column {
	return &column{
		name: name,
		typ:  BoolType{},
	}
}

func Int64Column(name string) Column {
	return &column{
		name: name,
		typ:  Int64Type{},
	}
}

func StringColumn(name string, width int64) Column {
	return &column{
		name: name,
		typ: &StringType{
			width: width,
		},
	}
}
