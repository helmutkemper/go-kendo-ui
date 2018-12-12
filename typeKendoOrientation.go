package telerik

type KendoOrientation int

var KendoOrientations = [...]string{
	"",
	"horizontal",
	"vertical",
}

func (el KendoOrientation) String() string {
	return KendoOrientations[el]
}

const (
	ORIENTATION_HORIZONTAL KendoOrientation = iota + 1
	ORIENTATION_VERTICAL
)
