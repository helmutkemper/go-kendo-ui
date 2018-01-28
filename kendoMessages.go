package telerik

import (
  "fmt"
  "html/template"
  "bytes"
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
  Close                                   String

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
  PromptInput                             String

}
func(el *KendoMessages) IsSet() bool {
  return el != nil
}
func(el *KendoMessages) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Messages", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

