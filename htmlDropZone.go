package telerik

type TypeHtmlDropZone int

var TypeHtmlDropZones = [...]string {
  "copy",
  "move",
  "link",
}

func (el TypeHtmlDropZone) String() string {
  return TypeHtmlDropZones[el]
}

const (
  COPY TypeHtmlDropZone = 1
  MOVE
  LINK
)
