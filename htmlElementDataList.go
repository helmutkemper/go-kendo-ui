package telerik

type HtmlElementFormDataList struct{
  /*
  The name of the control, which is submitted with the form data.
  */
  Name                        String

  /*
  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
  checkbox.
  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
  changed before the reload.
  */
  Value                       String

  /*
  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
  just as descendants of their form elements. An input can only be associated with one form.
  */
  Form                        String

  /*
  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Disabled                    Boolean

  /*
  The HTML <option> element is used to define an item contained in a <select>, an <optgroup>, or a <datalist> element.
  As such, <option> can represent menu items in popups and other lists of items in an HTML document.
  */
  Options                     map[string]string

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementFormDataList)String() string {
  out := `<input ` + el.Global.String() + ` list="` + el.Name.String() + `" ` + el.Value.ToAttr("value") + el.Form.ToAttr("form") + el.Disabled.ToAttrSet("disabled") + `>`
  out += `<datalist id="` + el.Name.String() + `">`
  for k, v := range el.Options{
    if v == "" {
      out += `<option value="` + k + `">`
    } else {
      out += `<option value="` + k + `">` + k + `</option>`
    }
  }
  out += `</datalist>`

  return out
}