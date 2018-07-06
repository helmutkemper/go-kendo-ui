package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type PanelBarMessages struct {
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/messages#messages.loading

  The text message shown while the root level items are loading. (default: "Loading...")
  */
  //    Example - customize loading message
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    $("#panelbar").kendoPanelBar({
  //      dataSource: {
  //        transport: {
  //          read: function(options) {
  //            // request always fails after 1s
  //            setTimeout(function() {
  //              options.error({});
  //            }, 1000);
  //          }
  //        }
  //      },
  //      messages: {
  //        loading: "Laden..."
  //      }
  //    });
  //    </script>
  Loading string `jsObject:"loading"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/messages#messages.requestFailed

  The text message shown when an error occurs while fetching the content. (default: "Request failed.")
  */
  /*
      Example - customize requestFailed message

      <div id="panelbar"></div>
      <script>
      $("#panelbar").kendoPanelBar({
        dataSource: {
          transport: {
            read: function(options) {
              // request always fails after 1s
              setTimeout(function() {
                options.error({});
              }, 1000);
            }
          }
        },
        messages: {
          requestFailed: "Anforderung fehlgeschlagen."
        }
      });
      </script>
  */
  RequestFailed string `jsObject:"requestFailed"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/messages#messages.retry

  The text message shown in the retry button. (default: "Retry")
  */
  /*
      Example - customize retry message

      <div id="panelbar"></div>
      <script>
      $("#panelbar").kendoPanelBar({
        dataSource: {
          transport: {
            read: function(options) {
              // request always fails after 1s
              setTimeout(function() {
                options.error({});
              }, 1000);
            }
          }
        },
        messages: {
          retry: "Wiederholen"
        }
      });
      </script>
  */
  Retry string `jsObject:"retry"`

  *ToJavaScriptConverter
}
func(el *PanelBarMessages) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "PanelBarMessages.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}
