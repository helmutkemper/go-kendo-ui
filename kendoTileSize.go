package telerik

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
  Width                                   int

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
  Height                                  int

}
func(el *KendoTileSize) IsSet() bool {
  return el != nil
}

