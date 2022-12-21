package spannermodel

type Table interface {
	Name() string
	Columns() []Column
	PrimaryKey() *PrimaryKeySpec
	Interleave() *InterleaveSpec
}

type Column interface {
	Name() string
	Type() ColumnType
}

type column struct {
	name string
	typ  ColumnType
}

func (c column) Name() string {
	return c.name
}

func (c column) Type() ColumnType {
	return c.typ
}

type SortOrder int

const (
	SortDefault SortOrder = iota
	SortAscending
	SortDescending
)

type PrimaryKeyColumn struct {
	name      string
	sortOrder SortOrder
}

func (c *PrimaryKeyColumn) Name() string {
	return c.name
}

func (c *PrimaryKeyColumn) SortOrder() SortOrder {
	return c.sortOrder
}

type PrimaryKeySpec struct {
	columns []*PrimaryKeyColumn
}

func (s *PrimaryKeySpec) Columns() []*PrimaryKeyColumn {
	return s.columns
}

func NewPrimaryKeySpec() *PrimaryKeySpec {
	return &PrimaryKeySpec{}
}

func (s *PrimaryKeySpec) AddColumn(name string, sortOrder SortOrder) *PrimaryKeySpec {
	s.columns = append(s.columns, &PrimaryKeyColumn{
		name:      name,
		sortOrder: sortOrder,
	})
	return s
}

type OnDelete int

const (
	OnDeleteDefault OnDelete = iota
	OnDeleteNoAction
	OnDeleteCascade
)

type InterleaveSpec struct {
	parent   string
	onDelete OnDelete
}

func NewInterleaveSpec(parent string, policy OnDelete) *InterleaveSpec {
	return &InterleaveSpec{
		parent:   parent,
		onDelete: policy,
	}
}

func (s *InterleaveSpec) Parent() string {
	return s.parent
}
