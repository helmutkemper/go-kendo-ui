package telerik

import (
  "fmt"
  "html/template"
  "bytes"
  log "github.com/helmutkemper/seelog"
  "reflect"
)

type KendoAnimation struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close  
  Example - configure the close animation
   <input id="dropdownlist" />
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: ["One", "Two"],
     animation: {
      close: {
        effects: "zoom:out",
        duration: 300
      }
     }
   });
   </script>
  */
  Close                                   *KendoClose                 `jsObject:"close"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open  The animation played when the suggestion popup is opened.
  Example - configure the open animation
   <input id="dropdownlist" />
   <script>
   $("#dropdownlist").kendoDropDownList({
     animation: {
      open: {
        effects: "zoom:in",
        duration: 300
      }
     }
   });
   </script>
  */
  Open                                    *KendoOpen                  `jsObject:"open"`

  *ToJavaScriptConverter
}
func(el *KendoAnimation) IsSet() bool {
  return el != nil
}
func(el *KendoAnimation) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Animation", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}
func(el *KendoAnimation) ToJavaScript() ([]byte) {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element, "")
  if err != nil {
    log.Criticalf( "KendoAnimation.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}
