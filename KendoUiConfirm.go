package telerik

import (
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoUiConfirm struct{
  Html                                  HtmlElementDiv                          `jsObject:"-"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/confirm#configuration-messages

  Defines the text of the labels that are shown within the confirm dialog. Used primarily for localization.

  Example
   <div id="confirm"></div>
   <script>
   $("#confirm").kendoConfirm({
     messages:{
       okText: "OK",
       cancel: "No"
     }
   }).data("kendoConfirm").open();
   </script>
  */
  Messages                                *KendoConfirmMessages                 `jsObject:"messages"`

  *ToJavaScriptConverter
}
func(el *KendoUiConfirm) ToJavaScript() []byte {
  var ret bytes.Buffer

  if el.Html.Global.Id == "" {
    log.Critical("KendoUiConfirm not have a html id for mount JavaScript code.")
    return []byte{}
  }

  element := reflect.ValueOf(el).Elem()
  data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoUiConfirm.Error: %v", err.Error() )
    return []byte{}
  }

  ret.Write( []byte(`$("#` + el.Html.Global.Id + `").kendoConfirm({`) )
  ret.Write( data )
  ret.Write( []byte(`});`) )

  return ret.Bytes()
}
func(el *KendoUiConfirm) ToHtml() []byte{
  return el.Html.ToHtml()
}