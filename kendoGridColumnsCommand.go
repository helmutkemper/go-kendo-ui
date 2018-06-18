package telerik

type kendoGridColumnsCommand int

var kendoGridColumnsCommands = [...]string {
  "",
  "edit",
  "destroy",
  "custom",
}

func (el kendoGridColumnsCommand) String() string {
  return kendoGridColumnsCommands[el]
}

const (
  COLUMNS_COMMAND_EDIT kendoGridColumnsCommand = iota + 1
  COLUMNS_COMMAND_DESTROY
  COLUMNS_COMMAND_CUSTOM
)
