package telerik

type TypeKendoGridColumnsCommand int

var TypeKendoGridColumnsCommands = [...]string {
  "",
  "edit",
  "destroy",
  "view",
  "custom",
}

func (el TypeKendoGridColumnsCommand) String() string {
  return TypeKendoGridColumnsCommands[el]
}

const (
  COLUMNS_COMMAND_EDIT TypeKendoGridColumnsCommand = iota + 1
  COLUMNS_COMMAND_DESTROY
  COLUMNS_COMMAND_VIEW
  COLUMNS_COMMAND_CUSTOM
)
