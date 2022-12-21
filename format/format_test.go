package format_test

import (
	"bytes"
	"testing"

	model "github.com/lestrrat-go/spanner-model"
	"github.com/lestrrat-go/spanner-model/format"
	"github.com/stretchr/testify/require"
)

type ExampleTable struct {
	Foo int
	Bar string
}

func (ExampleTable) Name() string {
	return "ExampleTable"
}

func (ExampleTable) Columns() []model.Column {
	return []model.Column{
		model.Int64Column(`Foo`),
		model.StringColumn(`Bar`, 1024),
	}
}

func (ExampleTable) Interleave() *model.InterleaveSpec {
	return model.NewInterleaveSpec(`Parent`, model.OnDeleteCascade)
}
func (ExampleTable) PrimaryKey() *model.PrimaryKeySpec {
	return model.NewPrimaryKeySpec().
		AddColumn(`Foo`, model.SortAscending)
}

func TestFormat(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, format.Format(&buf, ExampleTable{}), `format.Format should succeed`)

	t.Logf("%s", buf.String())
}
