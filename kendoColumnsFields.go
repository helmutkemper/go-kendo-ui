package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoColumnsFields struct {
  Title string `jsObject:"title"`

  Columns []*KendoColumnsFields `jsObject:"columns"`

  Field string `jsObject:"field"`

  *ToJavaScriptConverter
}
func(el *KendoColumnsFields) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoColumnsFields.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}