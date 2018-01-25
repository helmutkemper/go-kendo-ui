package telerik

type HtmlInputFile struct{
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
  This Boolean attribute indicates whether the user can enter more than one value. This attribute applies when the type
  attribute is set to email or file, otherwise it is ignored.
  */
  Multiple                    Boolean

  /*
  This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type
  attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS
  pseudo-classes will be applied to the field as appropriate.
  */
  Required                    Boolean

  /*
  If the value of the type attribute is file, then this attribute will indicate the types of files that the server
  accepts, otherwise it will be ignored. The value must be a comma-separated list of unique content type specifiers:
  A file extension starting with the STOP character (U+002E). (e.g. .jpg, .png, .doc).
  A valid MIME type with no extensions.
  audio/* representing sound files.
  video/* representing video files.
  image/* representing image files.
  */
  Accept                      StringArr

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputFile)String() string {
  return `<input ` + el.Global.String() + ` type="file" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Accept.ToAttr("accept") + el.Required.ToAttrSet("required") + el.Multiple.ToAttrSet("multiple") + `>`
}