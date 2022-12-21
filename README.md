
# DESCRIPTION

Suppose you have some objects that represent a Spanner table, perhaps
a protobuf message.

```go
type Example struct {
  Foo string `protobuf:"string,1,opt,name=foo,proto3" json:"foo,omitempty"
  Bar string `protobuf:"string,2,opt,name=bar,proto3" json:"bar,omitempty"
  Baz string `protobuf:"string,2,opt,name=baz,proto3" json:"baz,omitempty"
}
```

But perhaps this message does not _exactly_ represent a Spanner table:
for example, you might have some extra columns in the Spanner table, or
maybe some column types do not exactly translate to/from the protobuf
message definition.

But it still would be nice to keep the protobuf message as its single
source of truth. Perhaps you would like generate some intermediate object
to handle the translation between the protobuf object and the Spanner
table. Perhaps you would like to generate the Spanner DDL to create the
Spanner table from code.

You can use this package along with the format package to embed Spanner
specific metadata into existing objects, as well as use them to generate
Spanner DDLs.

# SYNOPSIS

To enable generating Spanner DDL for the example above, create another
file alongside the generated `*.pb.go` file:

```go
func (Example) Name() string {
  return `ExampleTable`
}

func (Example) Columns() []format.Column {
  return []format.Column{
    model.StringColumn(`Foo`, 1024),
    model.StringColumn(`Bar`, 1024),
    model.StringColumn(`Baz`, 1024),
  }
}

func (Example) PrimaryKey() *model.PrimaryKeySpec {
  return &model.PrimaryKeySpec{
    Columns: []string{
      `Foo`,
      `Bar`,
      `Baz`,
    },
  }
}
```

Feed the struct into `format.Format()` and you will end up with a DDL:

```go
var buf bytes.Buffer
format.Format(&buf, Example{})

// CREATE TABLE `Example` (
//   Foo STRING(1024),
//   Bar STRING(1024),
//   Baz STRING(1024)
// ) PRIMARY KEY (`Foo`, `Bar`, `Baz`)
```
