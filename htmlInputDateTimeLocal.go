package telerik

// <input> elements of type datetime-local create input controls that let the user easily enter both a date and a time,
// including the year, month, and day as well as the time in hours and minutes. The user's local time zone is used. The
// control's UI varies in general from browser to browser; at the moment support is patchy, with only Chrome/Opera and
// Edge on desktop and most modern versions of mobile browsers having usable implementations. In other browsers, these
// degrade gracefully to simple <input type="text"> controls.
//
// It's worth noting that seconds are not supported.
//
// Because of the limited browser support for datetime-local, and the variations in how the inputs work, it may
// currently still be best to use a framework or library to present these, or to use a custom input of your own. Another
// option is to use separate date and time inputs, each of which is more widely supported than datetime-local.
//
// Some browsers may resort to a text-only input element that validates that the results are legitimate date/time values
// before letting them be delivered to the server, as well, but you shouldn't rely on this behavior since you can't
// easily predict it.
//
// <input id="datetime" type="datetime-local">
type HtmlInputDateTimeLocal struct{
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

  /*
  Identifies a list of pre-defined options to suggest to the user. The value must be the id of a <datalist> element in
  the same document. The browser displays only options that are valid values for this input element. This attribute is
  ignored when the type attribute's value is hidden, checkbox, radio, file, or a button type.
  */
  List                        String

  ValueAsDate                 Boolean
  ValueAsNumber               Boolean
  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputDateTimeLocal)String() string {
  return `<input ` + el.Global.String() + ` type="datetime-local" ` + el.Name.ToAttr("name") + el.List.ToAttr("list") + el.ValueAsDate.ToAttr("valueasdate") + el.ValueAsNumber.ToAttr("valueasnumber") + el.Value.ToAttr("value") + `>`
}