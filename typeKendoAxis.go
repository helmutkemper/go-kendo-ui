package telerik

type KendoAxis int

var KendoAxes = [...]string{
	"",
	"x",
	"y",
}

func (el KendoAxis) String() string {
	return KendoAxes[el]
}

const (
	AXIS_X KendoAxis = iota + 1
	AXIS_Y
)
