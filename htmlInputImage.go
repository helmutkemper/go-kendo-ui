package telerik

// <input> elements of type "image" are used to create graphical submit buttons, i.e. submit buttons that take the form
// of an image rather than text.
/*
 <input id="image" type="image" alt="Login" src="./login.png">
*/
type HtmlInputImage struct{
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
  The alt attribute provides alt text for the image, so screen reader users can get a better idea of what the button is
  used for. It will also display if the image can't be shown for any reason (for example if the path is misspelled). If
  possible, use text which matches the label you'd use if you were using a standard submit button.
  */
  Alt                         String

  /*
  The src attribute is used to specify the path to the image you want to display in the button.
  */
  Src                         String

  /*
  The width and height attributes are used to specify the width and height the image should be shown at, in pixels. The
  button is the same size as the image; if you need the button's hit area to be bigger than the image, you will need to
  use CSS (e.g. padding). Also, if you specify only one dimension, the other is automatically adjusted so that the image
  maintains its original aspect ratio.
  */
  Width                       Int

  /*
  The width and height attributes are used to specify the width and height the image should be shown at, in pixels. The
  button is the same size as the image; if you need the button's hit area to be bigger than the image, you will need to
  use CSS (e.g. padding). Also, if you specify only one dimension, the other is automatically adjusted so that the image
  maintains its original aspect ratio.
  */
  Height                      Int

  /*
  The URI of a program that processes the information submitted by the input element; overrides the action attribute of
  the element's form owner.
  */
  FormAction                  String

  /*
  Specifies the type of content that is used to submit the form to the server. Possible values are:
  > application/x-www-form-urlencoded: The default value if the attribute is not specified.
  > text/plain.
  If this attribute is specified, it overrides the enctype attribute of the element's form owner.
  */
  FormEncType                 String

  /*
  Specifies the HTTP method that the browser uses to submit the form. Possible values are:
  post: The data from the form is included in the body of the form and is sent to the server.
  get: The data from the form is appended to the form attribute URI, with a '?' as a separator, and the resulting URI is
  sent to the server. Use this method when the form has no side-effects and contains only ASCII characters.
  If specified, this attribute overrides the method attribute of the element's form owner.
  */
  FormMethod                  String

  /*
  A Boolean attribute specifying that the form is not to be validated when it is submitted. If this attribute is
  specified, it overrides the novalidate attribute of the element's form owner.
  */
  FormNoValidate              Boolean

  /*
  A name or keyword indicating where to display the response that is received after submitting the form. This is a name
  of, or keyword for, a browsing context (for example, tab, window, or inline frame). If this attribute is specified,
  it overrides the target attribute of the element's form owner. The following keywords have special meanings:
  > _self: Load the response into the same browsing context as the current one. This value is the default if the
  attribute is not specified.
  > _blank: Load the response into a new unnamed browsing context.
  > _parent: Load the response into the parent browsing context of the current one. If there is no parent, this option
  behaves the same way as _self.
  > _top: Load the response into the top-level browsing context (that is, the browsing context that is an ancestor of
  the current one, and has no parent). If there is no parent, this option behaves the same as _self.
  */
  FormTarget                  String

  /*
  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Disabled                    Boolean

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputImage)String() string {
  return `<input ` + el.Global.String() + ` type="image" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Form.ToAttr("form") + el.Disabled.ToAttrSet("disabled") + el.FormAction.ToAttr("formaction") + el.Alt.ToAttr("alt") + el.Src.ToAttr("src") + el.Width.ToAttr("width") + el.Height.ToAttr("height") + el.FormEncType.ToAttr("formenctype") + el.FormMethod.ToAttr("formmethod") + el.FormTarget.ToAttr("formtarget") + el.FormNoValidate.ToAttrSet("formnovalidate") + `>`
}