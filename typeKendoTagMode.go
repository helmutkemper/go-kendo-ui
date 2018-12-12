package telerik

type KendoTagMode int

var KendoTagModes = [...]string{
	"",
	"multiple",
	"single",
}

func (el KendoTagMode) String() string {
	return KendoTagModes[el]
}

const (
	TAG_MODE_MULTIPLE KendoTagMode = iota + 1
	TAG_MODE_SINGLE
)
