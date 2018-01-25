package telerik

type HtmlElementFormSelect struct{
  /*
  The name of the control, which is submitted with the form data.
  */
  Name                        String

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
  This Boolean attribute indicates whether the user can enter more than one value. This attribute applies when the type
  attribute is set to email or file, otherwise it is ignored.
  */
  Multiple                    String

  /*
  This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type
  attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS
  pseudo-classes will be applied to the field as appropriate.
  */
  Required                    Boolean

  /*
  The initial size of the control. This value is in pixels unless the value of the type attribute is text or password,
  in which case it is an integer number of characters. Starting in HTML5, this attribute applies only when the type
  attribute is set to text, search, tel, url, email, or password, otherwise it is ignored. In addition, the size must be
  greater than zero. If you do not specify a size, a default value of 20 is used. HTML5 simply states "the user agent
  should ensure that at least that many characters are visible", but different characters can have different widths in
  certain fonts. In some browsers, a certain string with x characters will not be entirely visible even if size is
  defined to at least x.
  */
  Size                        Int

  /*
  The HTML <option> element is used to define an item contained in a <select>, an <optgroup>, or a <datalist> element.
  As such, <option> can represent menu items in popups and other lists of items in an HTML document.
  */
  Options                     map[string]string

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementFormSelect)String() string {
  out := `<select ` + el.Global.String() + el.Name.ToAttr("name") + el.Form.ToAttr("form") + el.Multiple.ToAttr("multiple") + el.Size.ToAttr("size") + el.Required.ToAttrSet("required") + el.Disabled.ToAttrSet("disabled") + `>`
  for k, v := range el.Options{
    if v == "" {
      out += `<option value="` + k + `">`
    } else {
      out += `<option value="` + k + `">` + k + `</option>`
    }
  }
  out += `</select>`

  return out
}