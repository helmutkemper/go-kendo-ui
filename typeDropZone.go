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
  HTML_DROP_ZONE_COPY TypeHtmlDropZone = 1
  HTML_DROP_ZONE_MOVE
  HTML_DROP_ZONE_LINK
)
