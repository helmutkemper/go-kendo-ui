package telerik

// The HTML <optgroup> element creates a grouping of options within a <select> element.
type HtmlElementFormOptGroup struct{
  /*
  Content inside html tag
  */
  Content                     Content

  /*
  The name of the group of options, which the browser can use when labeling the options in the user interface. This
  attribute is mandatory if this element is used.
  */
  Label                       String

  /*
  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Disabled                    Boolean

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementFormOptGroup)String() string {
  return `<label ` + el.Global.String() + el.Label.ToAttr("label") + el.Disabled.ToAttrSet("disabled") + `>` + el.Content.String() + `</label>`
}