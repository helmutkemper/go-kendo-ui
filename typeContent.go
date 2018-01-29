package telerik

import (
  "fmt"
)

type Content []interface{}

func(el Content) String() string {
  out := ``
  for _, v := range el{
    switch outConverted := v.(type) {
    case string:
      out += v.(string)
    case String:
      out += v.(String).String()
    case HtmlElementFormLegend:
      out += outConverted.String()
    case HtmlElementDiv:
      out += outConverted.String()
    case HtmlElementFormFieldSet:
      out += outConverted.String()
    case HtmlElementForm:
      out += outConverted.String()
    case HtmlElementFormSelect:
      out += outConverted.String()
    case HtmlElementFormTextArea:
      out += outConverted.String()
    case HtmlElementFormLabel:
      out += outConverted.String()
    case HtmlElementFormMeter:
      out += outConverted.String()
    case HtmlElementFormOptGroup:
      out += outConverted.String()
    case HtmlElementFormOutput:
      out += outConverted.String()
    case HtmlElementFormProgress:
      out += outConverted.String()
    case HtmlElementFormButton:
      out += outConverted.String()
    case HtmlElementFormDataList:
      out += outConverted.String()
    case HtmlInputText:
      out += outConverted.String()
    default:
      fmt.Printf("error: type %T not found in typeContent\n", outConverted)
    }
  }

  return out
}