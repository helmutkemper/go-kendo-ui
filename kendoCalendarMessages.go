package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoCalendarMessages struct {
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar/configuration/messages

  Allows localization of the strings that are used in the widget. (default: "")

  Example
  <div id="calendar"></div>
  <script>
  $("#calendar").kendoCalendar({
      "weekNumber": true,
      "messages": {
          "weekColumnHeader": "W"
      }
   })
  </script>

  Allows customization of the week column header text. Set the value to make the widget compliant with web accessibility standards. (default: "")

  Example
  <div id="calendar"></div>
  <script>
  $("#calendar").kendoCalendar({
      "weekNumber": true,
      "messages": {
          "weekColumnHeader": "W"
      }
   })
  </script>
  */

  WeekColumnHeader                    string                  `jsObject:"weekColumnHeader"`

  *ToJavaScriptConverter
}
func(el *KendoCalendarMessages) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoCalendarMessages.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}