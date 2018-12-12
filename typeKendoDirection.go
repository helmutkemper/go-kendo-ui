package telerik

type KendoDirection int

func (el KendoDirection) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoDirection) String() string {
	return TypeDirections[el]
}

var TypeDirections = [...]string{
	"",
	"asc",
	"desc",
}

const (
	DIRECTION_ASC KendoDirection = iota + 1
	DIRECTION_DESC
)
