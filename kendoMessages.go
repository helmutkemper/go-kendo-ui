package telerik

import (
  "fmt"
  "html/template"
  "bytes"
  log "github.com/helmutkemper/seelog"
  "reflect"
)

type KendoMessages struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages.close  The title of the close button.
  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     messages:{
       close: "Close Me!"
     }
   });
   </script>
  */
  Close                                   String                  `jsObject:"close"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages.promptInput  The title of the prompt input.
  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     messages:{
       promptInput: "Input!"
     }
   });
   </script>
  */
  PromptInput                             String                  `jsObject:"promptinput"`

  *ToJavaScriptConverter
}
func(el *KendoMessages) IsSet() bool {
  return el != nil
}
func(el *KendoMessages) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Messages", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}
func(el *KendoMessages) ToJavaScript() ([]byte) {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element, "")
  if err != nil {
    log.Criticalf( "KendoMessages.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}
