package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoUiColorPicker struct{
  HtmlId                                  String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-buttons

  Specifies whether the widget should display the Apply / Cancel buttons.
  Applicable only for the HSV selector, when a <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-palette"><b><u>pallete</u></b></a> is not specified.

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     buttons: false
   })
   </script>
  */

  Buttons                                 Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-clearButton

  Specifies whether the widget should display the 'Clear color' button.
  Applicable only for the HSV selector, when a <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-palette"><b><u>pallete</u></b></a> is not specified.

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     clearButton: false
   });
   </script>
  */

  ClearButton                             Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-columns

  The number of columns to show in the color dropdown when a <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-palette"><b><u>pallete</u></b></a> is specified. This is automatically initialized for the "basic" and "websafe" palettes. If you use a custom palette then you can set this to some value that makes sense for your colors.

  Example - wrap list of colors on two rows with 3 columns
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     palette: [ "#000", "#333", "#666", "#999", "#ccc", "#fff" ],
     columns: 3
   });
   </script>
  */

  Columns                                 Int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-tileSize

  The size of a color cell.

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     palette: "basic",
     tileSize: 32
   });
   </script>
  */

  TileSize                                *KendoTileSize

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-messages

  Allows localization of the strings that are used in the widget.

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     messages: {
       apply: "Update",
       cancel: "Discard"
     }
   })
   </script>
  */

  Messages                                *KendoColorMessages

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-palette

  When a non-null <b><u>palette</u></b> argument is supplied, the drop-down will be a simple color picker that lists the colors. The following are supported:
  "basic" -- displays 20 basic colors
  "websafe" -- display the "web-safe" color palette
  otherwise, pass a string with colors in HEX representation separated with commas, or an array of colors, and it will display that palette instead. If you pass an array it can contain strings supported by <a href="/kendo-ui/api/javascript/kendo#parseColor">parseColor</a> or <a href="/kendo-ui/api/javascript/kendo#Color">Color</a> objects.
  If <b><u>palette</u></b> is missing or <b><u>null</u></b>, the widget will display the HSV selector.

  Example - use "websafe" palette
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     palette: "websafe"
   });
   </script>
  */

  Palette                                 String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-opacity

  Only for the HSV selector. If <b><u>true</u></b>, the widget will display the opacity slider. Note that currently in HTML5 the <b><u>&lt;input type="color"&gt;</u></b> does not support opacity.

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     opacity: true
   });
   </script>
  */

  Opacity                                 Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-preview

  Only applicable for the HSV selector.
  Displays the color preview element, along with an input field where the end user can paste a color in a CSS-supported notation.

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     preview: false
   });
   </script>
  */

  Preview                                 Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-toolIcon

  A CSS class name to display an icon in the color picker button. If specified, the HTML for the element will look like this:

  Example
   <input id="colorpicker" type="color" />
   <script>
   $("#colorpicker").kendoColorPicker({
     toolIcon: "k-foreColor"
   });
   </script>
  */

  ToolIcon                                String

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-value

  The initially selected color. Note that when initializing the widget from an <b><u>&lt;input&gt;</u></b> element, the initial color will be decided by the field instead.

  Example
   <div id="colorpicker"></div>
   <script>
   $("#colorpicker").kendoColorPicker({
     value: "#b72bba"
   });
   </script>
  */

  Value                                   String
}
func(el *KendoUiColorPicker) IsSet() bool {
  return el != nil
}
func(el *KendoUiColorPicker) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "KendoUiColorPicker", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

