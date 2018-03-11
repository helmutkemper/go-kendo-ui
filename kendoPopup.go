package telerik

import (
  "fmt"
  "html/template"
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

//fixme: incompleto - https://docs.telerik.com/kendo-ui/api/javascript/ui/popup
type KendoPopup struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.appendTo  Defines a jQuery selector that will be used to find a container element, where the popup will be appended to.
  Example - append the popup to a specific element
   <div id="container">
       <input id="dropdownlist" />
   </div>
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: [
       { id: 1, name: "Apples" },
       { id: 2, name: "Oranges" }
     ],
     dataTextField: "name",
     dataValueField: "id",
     popup: {
       appendTo: $("#container")
     }
   });
   </script>
  */
  AppendTo                               *JavaScript                              `jsObject:"appendTo"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.origin  Specifies how to position the popup element based on anchor point. The value is space separated "y" plus "x" position.
  The available "y" positions are:
  The available "x" positions are:
  Example - append the popup to a specific element
   <div id="container">
       <input id="dropdownlist" />
   </div>
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: [
       { id: 1, name: "Apples" },
       { id: 2, name: "Oranges" }
     ],
     dataTextField: "name",
     dataValueField: "id",
     popup: {
       origin: "top left"
     }
   });
   </script>
  */
  Origin                                  KendoOrigin                             `jsObject:"origin"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.position  Specifies which point of the popup element to attach to the anchor's origin point. The value is space separated "y" plus "x" position.
  The available "y" positions are:
  The available "x" positions are:
  Example - append the popup to a specific element
   <div id="container">
       <input id="dropdownlist" />
   </div>
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: [
       { id: 1, name: "Apples" },
       { id: 2, name: "Oranges" }
     ],
     dataTextField: "name",
     dataValueField: "id",
     popup: {
       origin: "top left"
     }
   });
   </script>
  */
  Position                                KendoPosition                           `jsObject:"position"`

  *ToJavaScriptConverter
}
func(el *KendoPopup) IsSet() bool {
  return el != nil
}
func(el *KendoPopup) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Popup", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}
func(el *KendoPopup) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoPopup.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}
