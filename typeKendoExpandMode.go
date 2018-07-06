package telerik

type TypeExpandMode int

var TypeExpandModes = [...]string {
  "",
  "single",
  "multiple",
}

func (el TypeExpandMode) String() string {
  return TypeExpandModes[el]
}

const (
  PANELBAR_EXPAND_MODE_SINGLE TypeExpandMode = iota + 1
  PANELBAR_EXPAND_MODE_MULTIPLE
)
