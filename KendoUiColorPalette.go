package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoUiColorPalette struct{
  HtmlId                                  String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpalette#configuration-palette

  Specifies the color palette to display. It can be a string with comma-separated colors in hex representation, an array of <a href="/kendo-ui/api/javascript/color"><b><u>kendo.Color</u></b> object</a> objects or of strings that <a href="/kendo-ui/api/javascript/kendo#parseColor">parseColor</a> understands. As a shortcut, you can pass "basic" to get the simple palette (this is the default) or "websafe" to get the Web-safe palette.

  Example - use "websafe" palette
   <div id="palette"></div>
   <script>
   $("#palette").kendoColorPalette({
     palette: "websafe"
   });
   </script>
  */

  Palette                                 String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpalette#configuration-columns

  The number of columns to display. When you use the "websafe" palette, this will automatically default to 18.

  Example - wrap list of colors on two rows with 3 columns
   <div id="palette"></div>
   <script>
   $("#palette").kendoColorPalette({
     palette: [ "#000", "#333", "#666", "#999", "#ccc", "#fff" ],
     columns: 3
   });
   </script>
  */

  Columns                                 Int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpalette#configuration-tileSize

  The size of a color cell.

  Example
   <div id="palette"></div>
   <script>
   $("#palette").kendoColorPalette({
     palette: "basic",
     tileSize: 32
   });
   </script>
  */

  TileSize                                *KendoTileSize

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpalette#configuration-value

  Specifies the initially selected color.

  Example
   <div id="palette"></div>
   <script>
   $("#palette").kendoColorPalette({
     palette: "basic",
     value: "#fff"
   });
   </script>
  */

  Value                                   String
}
func(el *KendoUiColorPalette) IsSet() bool {
  return el != nil
}
func(el *KendoUiColorPalette) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "KendoUiColorPalette", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

