package telerik

// <input> elements of type "search" are text fields designed for the user to enter search queries into. These are
// functionally identical to text inputs, but may be styled differently by the user agent.
//
// <input type="search">
type HtmlInputSearch struct{
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
  If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the maximum
  number of characters (in UTF-16 code units) that the user can enter. For other control types, it is ignored. It can
  exceed the value of the size attribute. If it is not specified, the user can enter an unlimited number of characters.
  Specifying a negative number results in the default behavior (i.e. the user can enter an unlimited number of
  characters). The constraint is evaluated only when the value of the attribute has been changed.
  */
  MaxLength                   Int

  /*
  If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the minimum
  number of characters (in Unicode code points) that the user can enter. For other control types, it is ignored.
  */
  MinLength                   Int

  /*
  A regular expression that the control's value is checked against. The pattern must match the entire value, not just
  some subset. Use the title attribute to describe the pattern to help the user. This attribute applies when the value
  of the type attribute is text, search, tel, url, email, or password, otherwise it is ignored. The regular expression
  language is the same as JavaScript RegExp algorithm, with the 'u' parameter that makes it treat the pattern as a
  sequence of unicode code points. The pattern is not surrounded by forward slashes.
  */
  Pattern                     String

  /*
  A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or
  line-feeds.
  */
  PlaceHolder                 String

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

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputSearch)String() string {
  return `<input ` + el.Global.String() + ` type="search" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Form.ToAttr("form") + el.AutoComplete.ToAttr("autocomplete") + el.List.ToAttr("list") + el.MaxLength.ToAttr("maxlength") + el.MinLength.ToAttr("minlength") + el.Pattern.ToAttr("pattern") + el.PlaceHolder.ToAttr("placeholder") + el.Size.ToAttr("size") + el.Required.ToAttrSet("required") + el.Disabled.ToAttrSet("disabled") + `>`
}