package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoTileSize struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-tileSize.width  The width of the color cell.
  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     palette: "basic",
     tileSize: { width: 40 }
   });
   </script>
  */
  Width                                   Int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-tileSize.height  The height of the color cell.
  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     palette: "basic",
     tileSize: { height: 40 }
   });
   </script>
  */
  Height                                  Int

}
func(el *KendoTileSize) IsSet() bool {
  return el != nil
}
func(el *KendoTileSize) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "TileSize", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

