package telerik

import (
  "bytes"
  "reflect"
)

// The HTML <datalist> element contains a set of <option> elements that represent the values available for other
// controls.
type HtmlElementFormDataList struct{
  /*
  The name of the control, which is submitted with the form data.
  */
  Name                        string                      `htmlAttr:"name"`

  /*
  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
  checkbox.
  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
  changed before the reload.
  */
  Value                       string                      `htmlAttr:"value"`

  /*
  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
  just as descendants of their form elements. An input can only be associated with one form.
  */
  Form                        string                      `htmlAttr:"form"`

  /*
  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Disabled                    Boolean                     `htmlAttrSet:"disabled"`

  Options                     map[string]string           `htmlAttrSet:"options"`

  Global                      HtmlGlobalAttributes        `htmlAttrSet:"-"`

  *ToJavaScriptConverter                                  `htmlAttrSet:"-"`
}
func(el *HtmlElementFormDataList)ToHtml() []byte {
  var buffer bytes.Buffer

  element := reflect.ValueOf(el).Elem()
  data := el.ToJavaScriptConverter.ToTelerikHtml(element)

  buffer.Write( []byte( `<input` ) )
  buffer.Write( el.Global.ToHtml() )
  buffer.Write( data )
  buffer.Write( []byte( `>` ) )


  buffer.WriteString(`<datalist id="` + el.Name + `">` )
  for k, v := range el.Options{
    if v == "" {
      buffer.WriteString(`<option value="` + k + `">` )
    } else {
      buffer.WriteString(`<option value="` + k + `">` + v + `</option>` )
    }
  }
  buffer.WriteString(`</datalist>` )

  return buffer.Bytes()
}