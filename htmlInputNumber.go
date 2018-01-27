package telerik

// <input> elements of type "number" are used to let the user enter a number. They include built-in validation to reject
// non-numerical entries. The browser may opt to provide stepper arrows to let the user increase and decrease the value
// using their mouse or by simply tapping with a fingertip.
//
// Note: Browsers that don't support type "number" fall back to using a standard "text" input.
//
// <input id="number" type="number">
type HtmlInputNumber struct{
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
  This attribute indicates whether the value of the control can be automatically completed by the browser.
  Possible values are:
  off: The user must explicitly enter a value into this field for every use, or the document provides its own
  auto-completion method. The browser does not automatically complete the entry.
  on: The browser is allowed to automatically complete the value based on values that the user has entered during
  previous uses, however on does not provide any further information about what kind of data the user might be expected
  to enter.
  @see typeNamesForAutocomplete.go
  */
  AutoComplete                Boolean

  /*
  Identifies a list of pre-defined options to suggest to the user. The value must be the id of a <datalist> element in
  the same document. The browser displays only options that are valid values for this input element. This attribute is
  ignored when the type attribute's value is hidden, checkbox, radio, file, or a button type.
  */
  List                        String

  /*
  A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or
  line-feeds.
  */
  PlaceHolder                 String

  /*
  This attribute indicates that the user cannot modify the value of the control. The value of the attribute is
  irrelevant. If you need read-write access to the input value, do not add the "readonly" attribute. It is ignored if
  the value of the type attribute is hidden, range, color, checkbox, radio, file, or a button type (such as button or
  submit).
  */
  Readonly                    Boolean

  ValueAsNumber               Boolean

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputNumber)String() string {
  return `<input ` + el.Global.String() + ` type="number" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Form.ToAttr("form") + el.Disabled.ToAttrSet("disabled") + el.AutoComplete.ToAttr("autocomplete") + el.List.ToAttr("list") + el.PlaceHolder.ToAttr("placeholder") + el.Readonly.ToAttrSet("readonly") + el.ValueAsNumber.ToAttr("valueasnumber") +`>`
}