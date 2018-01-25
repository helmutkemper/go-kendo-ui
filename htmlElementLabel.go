package telerik

type HtmlElementFormLabel struct{
  /*
  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
  just as descendants of their form elements. An input can only be associated with one form.
  */
  Form                        String

  /*
  The id of a labelable form-related element in the same document as the label element. The first such element in the
  document with an ID matching the value of the for attribute is the labeled control for this label element.
  A label element can have both a for attribute and a contained control element, as long as the for attribute points to
  the contained control element.
  */
  For                         String

  /*
  Content inside html tag
  */
  Content                     string

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementFormLabel)String() string {
  return `<label ` + el.Global.String() + el.For.ToAttr("for") + el.Form.ToAttr("form") + `>` + el.Content + `</label>`
}