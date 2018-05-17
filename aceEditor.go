package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
)

/*
@see https://github.com/ajaxorg/ace/wiki/Embedding-API
*/
type AceEditor struct{
  Html                                    *HtmlElementDiv         `jsSetFunction:"-"`
  Theme                                   AceTheme                `jsSetFunction:"setTheme"`
  TabSize                                 int                     `jsSetFunction:"setTabSize"`
  Content                                 Content                 `jsSetFunction:"setValue"`
  SoftTabs                                bool                    `jsSetFunction:"setUseSoftTabs"`
  FontSize                                string                  `jsSetFunction:"setUseSoftTabs"`
  WrapMode                                bool                    `jsSetFunction:"setUseWrapMode"`
  HighlightLine                           bool                    `jsSetFunction:"setHighlightActiveLine"`
  ShowPrintMargin                         bool                    `jsSetFunction:"setShowPrintMargin"`
  ReadOnly                                bool                    `jsSetFunction:"setReadOnly"`

  *ToJavaScriptConverter
}
func(el *AceEditor) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoClose.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}