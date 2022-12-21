package spannermodel_test

import (
	"testing"

	model "github.com/lestrrat-go/spanner-model"
)

func TestColumn(t *testing.T) {
	col := model.BoolColumn(`name`)
	t.Logf("%s, %s", col.Name(), col.Type())
}
