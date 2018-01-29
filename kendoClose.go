package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoClose struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close.effects  The effect(s) to use when playing the close animation. Multiple effects should be separated with a space.
  <a href="/kendo-ui/api/javascript/effects/common">Complete list of available animations</a>
  
  */
  Effects                                 String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close.duration  The duration of the close animation in milliseconds.
  
  */
  Duration                                Int

}
func(el *KendoClose) IsSet() bool {
  return el != nil
}
func(el *KendoClose) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Close", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

