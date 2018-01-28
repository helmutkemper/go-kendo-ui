package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoActions struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions.text  The text to be shown in the action's button.
  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
         }]
       });
   </script>
  */
  Text                                    String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions.action  The callback function to be called after pressing the action button.
  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
             action: function(e){
                 // e.sender is a reference to the dialog widget object
                 alert("OK action was clicked");
                 // Returning false will prevent the closing of the dialog
                 return true;
             },
         }]
       });
   </script>
  */
  Action                                  String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions.primary  A boolean property indicating whether the action button will be decorated as primary button or not.
  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
             primary: true
         }]
       });
   </script>
  */
  Primary                                 Boolean

}
func(el *KendoActions) IsSet() bool {
  return el != nil
}
func(el *KendoActions) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Actions", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

