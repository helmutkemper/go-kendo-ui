package telerik

import (
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoUiColorPalette struct{
  Div                                   HtmlElementDiv                              `jsObject:"-"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpalette#configuration-palette

  When a non-null palette argument is supplied, the drop-down will be a simple color picker that lists the colors. The following are supported: (default: null)

  "basic" -- displays 20 basic colors
  "websafe" -- display the "web-safe" color palette
  otherwise, pass a string with colors in HEX representation separated with commas, or an array of colors, and it will display that palette instead. If you pass an array it can contain strings supported by parseColor or Color objects.
  If palette is missing or null, the widget will display the HSV selector.

  Example - use "websafe" palette
   <div id="palette"></div>
   <script>
   $("#palette").kendoColorPalette({
     palette: "websafe"
   });
   </script>

  Example - use list of colors
  <input id="colorpicker" type="color" />
  <script>
  $("#colorpicker").kendoColorPicker({
    palette: [ "#000", "#333", "#666", "#999", "#ccc", "#fff" ],
    columns: 6
  });
  </script>
  */
  Palette                               interface{}                                 `jsObject:"palette" jsType:"[]string,KendoColorPalette"`

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
  Columns                                 int                                       `jsObject:"columns"`

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
  TileSize                                *KendoTileSize                            `jsObject:"tileSize"`

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
  Value                                   string                                    `jsObject:"value"`

  *ToJavaScriptConverter
}
func(el *KendoUiColorPalette) ToJavaScript() string {
  var ret bytes.Buffer
  if el.Div.Global.Id == "" {
    log.Critical("KendoUiColorPalette not have a html id for mount JavaScript code.")
    return ""
  }

  element := reflect.ValueOf(el).Elem()
  data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoUiColorPalette.Error: %v", err.Error() )
    return ""
  }

  ret.Write( []byte(`$("#` + el.Div.Global.Id + `").kendoColorPalette({`) )
  ret.Write( data )
  ret.Write( []byte(`});`) )

  return ret.String()
}/*
func(el *KendoUiColorPalette) ToHtml() string {
  return el.Div.String()
}*/