package telerik

import (
  log "github.com/helmutkemper/seelog"
  "reflect"
)

type KendoAnimationVertical struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close
  Example - configure the close animation
   <input id="dropdownlist" />
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: ["One", "Two"],
     animation: {
      close: {
        effects: "zoom:out",
        duration: 300
      }
     }
   });
   </script>
  */
  Collapse KendoClose `jsObject:"collapse"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open  The animation played when the suggestion popup is opened.
  Example - configure the open animation
   <input id="dropdownlist" />
   <script>
   $("#dropdownlist").kendoDropDownList({
     animation: {
      open: {
        effects: "zoom:in",
        duration: 300
      }
     }
   });
   </script>
  */
  Expand KendoOpen `jsObject:"expand"`

  *ToJavaScriptConverter
}
func(el *KendoAnimationVertical) ToJavaScript() ([]byte) {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoAnimationVertical.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}
