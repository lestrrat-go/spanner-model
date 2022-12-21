package format

import (
	"bytes"
	"fmt"
	"io"

	model "github.com/lestrrat-go/spanner-model"
)

func Format(dst io.Writer, in interface{}) error {
	var f Formatter
	switch in := in.(type) {
	case model.Table:
		return f.FormatTable(dst, in)
	case model.Column:
		return f.FormatColumn(dst, in)
	default:
		return fmt.Errorf(`invalid argument to Format() (%T)`, in)
	}
}

type Formatter struct{}

func (f *Formatter) FormatTable(dst io.Writer, in model.Table) error {
	var buf bytes.Buffer

	buf.WriteString("CREATE TABLE `")
	buf.WriteString(in.Name())
	buf.WriteString("` (")
	for i, col := range in.Columns() {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
		if err := f.FormatColumn(&buf, col); err != nil {
			return fmt.Errorf("failed to format column %d: %w", i, err)
		}
	}
	buf.WriteString("\n) PRIMARY KEY (")

	pkey := in.PrimaryKey()
	if pkey == nil || len(pkey.Columns()) == 0 {
		return fmt.Errorf("primary keys may not be empty (%T)", in)
	}

	for i, col := range pkey.Columns() {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("`")
		buf.WriteString(col.Name())
		buf.WriteString("`")

		switch col.SortOrder() {
		case model.SortAscending:
			buf.WriteString(` ASC`)
		case model.SortDescending:
			buf.WriteString(` DESC`)
		}
	}
	buf.WriteString(")")

	if ispec := in.Interleave(); ispec != nil {
		buf.WriteString(",  \nINTERLEAVE IN PARENT `")
		buf.WriteString(ispec.Parent())
		buf.WriteString("`")
	}

	buf.WriteTo(dst)
	return nil
}

func (f *Formatter) FormatColumn(dst io.Writer, in model.Column) error {
	fmt.Fprintf(dst, "`%s` %s", in.Name(), in.Type().Name())
	if wcol, ok := in.Type().(model.ColumnWidther); ok {
		fmt.Fprintf(dst, " (%d)", wcol.Width())
	}
	return nil
}
