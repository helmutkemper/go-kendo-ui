package telerik

type KendoButtonLayout int

var KendoButtonLayouts = [...]string{
	"",
	"normal",
	"stretched",
}

func (el KendoButtonLayout) String() string {
	return KendoButtonLayouts[el]
}

const (
	BUTTON_LAYOUT_NORMAL KendoButtonLayout = iota + 1
	BUTTON_LAYOUT_STRETCHED
)
