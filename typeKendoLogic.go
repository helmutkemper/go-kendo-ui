package telerik

type KendoLogic int

var KendoLogics = [...]string{
	"",
	"and",
	"or",
}

func (el KendoLogic) String() string {
	return KendoLogics[el]
}

const (
	LOGIC_AND KendoLogic = iota + 1
	LOGIC_OR
)
