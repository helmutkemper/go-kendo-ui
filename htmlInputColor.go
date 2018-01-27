package telerik

// <input> elements of type "color" provide a user interface element that lets a user specify a color, either by using a
// visual color picker interface or by entering the color into a text field in "#rrggbb" hexadecimal format. Only simple
// colors (with no alpha channel) are allowed.
//
// The element's presentation may vary substantially from one browser and/or platform to another — it might be a simple
// textual input that automatically validates to ensure that the color information is entered in the proper format, or a
// platform-standard color picker, or some kind of custom color picker window.
/*
 <input type="color">
*/
type HtmlInputColor struct{
  /*
  The name of the control, which is submitted with the form data.
  @see typeNamesForAutocomplete.go
  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
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

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputColor)String() string {
  return `<input ` + el.Global.String() + ` type="color" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + `">`
}