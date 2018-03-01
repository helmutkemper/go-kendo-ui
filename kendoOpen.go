package telerik

import (
  "fmt"
  "html/template"
  "bytes"
  log "github.com/helmutkemper/seelog"
  "reflect"
)

type KendoOpen struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open.effects  The effect(s) to use when playing the open animation. Multiple effects should be separated with a space.
  <a href="/kendo-ui/api/javascript/effects/common">Complete list of available animations</a>
  
  */
  Effects                                 String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open.duration  The duration of the open animation in milliseconds.
  
  */
  Duration                                Int

  *ToJavaScriptConverter
}
func(el *KendoOpen) IsSet() bool {
  return el != nil
}
func(el *KendoOpen) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Open", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}
func(el *KendoOpen) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element, "")
  if err != nil {
    log.Criticalf( "KendoOpen.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}
