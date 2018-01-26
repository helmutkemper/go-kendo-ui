package telerik

// The HTML <fieldset> element is used to group several controls as well as labels (<label>) within a web form.
type HtmlElementFormFieldSet struct{
  /*
  The name of the control, which is submitted with the form data.
  */
  Name                        String

  /*
  Content inside html tag
  */
  Content                     Content

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

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementFormFieldSet)String() string {
  return `<fieldset ` + el.Global.String() + el.Name.ToAttr("name") + el.Form.ToAttr("form") + el.Disabled.ToAttrSet("disabled") + `>` + el.Content.String() + `</fieldset>`
}