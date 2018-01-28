package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoUiDropTarget struct{
  HtmlId                                  String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/droptarget#configuration-group

  Used to group sets of draggable and drop targets. A draggable with the same group value as a drop target will be accepted by the drop target.

  Example
   <div class="orange"></div>
   <div class="orange"></div>
   <div class="purple"></div>
   <div class="purple"></div>
   <div id="orangeArea"></div>
   <div id="purpleArea"></div>
   
   <script>
     $(".orange").kendoDraggable({
       group: "orangeGroup",
       hint: function(element) {
         return element.clone();
       }
     });
   
     $(".purple").kendoDraggable({
       group: "purpleGroup",
       hint: function(element) {
         return element.clone();
       }
     });
   
     $("#orangeArea").kendoDropTarget({ group: "orangeGroup", drop: onDrop });
     $("#purpleArea").kendoDropTarget({ group: "purpleGroup", drop: onDrop });
   
     function onDrop(e) {
       e.draggable.destroy();
       e.draggable.element.remove();
     }
   </script>
   <style>
     .orange, .purple{
       width: 50px;
       height: 50px;
       border: 2px solid green;
       margin: 5px;
     }
     #orangeArea, #purpleArea {
       width: 200px;
       height: 200px;
       border: 2px solid green;
       margin: 5px;
     }
     .orange, #orangeArea { background-color: orange; }
     .purple, #purpleArea { background-color: purple; }
   </style>
  */

  Group                                   String
}
func(el *KendoUiDropTarget) IsSet() bool {
  return el != nil
}
func(el *KendoUiDropTarget) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "KendoUiDropTarget", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

