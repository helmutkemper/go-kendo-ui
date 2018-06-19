package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoGridColumnsIconClass struct {
  Edit string `jsObject:"edit"`
  Update string `jsObject:"update"`
  Cancel string `jsObject:"cancel"`

  *ToJavaScriptConverter
}
func(el *KendoGridColumnsIconClass) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoGridColumnsIconClass.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}