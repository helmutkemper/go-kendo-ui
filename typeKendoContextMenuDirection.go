package telerik

type KendoContextMenuDirection int

func (el KendoContextMenuDirection) String() string {
	return KendoContextMenuDirections[el]
}

var KendoContextMenuDirections = [...]string{
	"",
	"top",
	"bottom",
	"left",
	"right",
}

const (
	CONTEXT_MENU_DIRECTION_TOP KendoContextMenuDirection = iota + 1
	CONTEXT_MENU_DIRECTION_BOTTOM
	CONTEXT_MENU_DIRECTION_LEFT
	CONTEXT_MENU_DIRECTION_RIGHT
)
