package telerik

type HtmlElementFormMeter struct{
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
  The minimum (numeric or date-time) value for this item, which must not be greater than its maximum (max attribute)
  value.
  */
  Min                         Int

  /*
  The maximum (numeric or date-time) value for this item, which must not be less than its minimum (min attribute) value.
  */
  Max                         Int

  /*
  The lower numeric bound of the high end of the measured range. This must be less than the maximum value (max
  attribute), and it also must be greater than the low value and minimum value (low attribute and min attribute,
  respectively), if any are specified. If unspecified, or if greater than the maximum value, the high value is equal to
  the maximum value.
  */
  High                        Int

  /*
  This attribute indicates the optimal numeric value. It must be within the range (as defined by the min attribute and
  max attribute). When used with the low attribute and high attribute, it gives an indication where along the range is
  considered preferable. For example, if it is between the min attribute and the low attribute, then the lower range is
  considered preferred.
  */
  Optimum                     Int

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementFormMeter)String() string {
  return `<meter ` + el.Global.String() + el.Form.ToAttr("form") + el.Min.ToAttr("min") + el.Max.ToAttr("max") + el.High.ToAttr("high") + el.Optimum.ToAttr("optimum") + `>` + el.Content.String() + `</meter>`
}