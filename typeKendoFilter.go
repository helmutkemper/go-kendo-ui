package telerik

type KendoFilter int

var KendoFilters = [...]string{
	"",
	"startswith.",
	"endswith",
	"contains",
}

func (el KendoFilter) String() string {
	return KendoFilters[el]
}

const (
	FILTER_STARTS_WITH KendoFilter = iota + 1
	FILTER_ENDS_WITH
	FILTER_CONTAINS
)
