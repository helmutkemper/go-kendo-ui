package telerik

import "bytes"

type CustomTemplate struct {
  Id                                      string                                  `jsObject:"-"`
  Template                                string                                  `jsObject:"-"`
  Footer                                  string                                  `jsObject:"-"`
  NoData                                  string                                  `jsObject:"-"`
  Dialog                                  KendoUiDialog                           `jsObject:"-"`
}
func(el *CustomTemplate) GetTemplate() []byte {
  return []byte(`<script id="` + el.Id + `Template" type="text/x-kendo-tmpl">` + el.Template + `</script>`)
}
func(el *CustomTemplate) GetFooterTemplate() []byte {
  return []byte(`<script id="` + el.Id + `FooterTemplate" type="text/x-kendo-tmpl">` + el.Footer + `</script>`)
}
func(el *CustomTemplate) GetNoDataTemplate() []byte {
  return []byte(`<script id="` + el.Id + `NoDataTemplate" type="text/x-kendo-tmpl">` + el.NoData + `</script>`)
}
func(el *CustomTemplate) GetFooterTemplateButton(buttonLabel, buttonCss string) []byte {
  return []byte(`<button class="k-button k-primary ` + buttonCss + `" onclick="` + el.Id + `CustomButtonClick()">` + buttonLabel + `</button>`)
}
func(el *CustomTemplate) GetFooterTemplateButtonAsString(buttonLabel, buttonCss string) string {
  ret := el.GetFooterTemplateButton(buttonLabel, buttonCss)
  return string( ret )
}
func(el *CustomTemplate) GetFooterTemplateButtonOpenDialogFunction() []byte {
  var buffer bytes.Buffer

  buffer.Write( []byte(`function ` + el.Id + `CustomButtonClick(){` ) )
  buffer.Write( el.Dialog.ToJavaScript() )
  buffer.Write( el.Dialog.MethodOpen() )
  buffer.Write( []byte(`}` ) )

  return buffer.Bytes()
}