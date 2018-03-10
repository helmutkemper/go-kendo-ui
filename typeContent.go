package telerik

import (
  //"fmt"
  "fmt"
  "bytes"
)

type Content []interface{}

func(el Content) Bytes() []byte {
  var buffer bytes.Buffer

  for _, v := range el{
    switch outConverted := v.(type) {
    case string:
      buffer.WriteString( v.(string) )
    case HtmlElementDiv:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormLabel:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputText:
      buffer.Write( outConverted.ToHtml() )
      /*case String:
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
        out += outConverted.String()*/
    default:
      fmt.Printf("typeContent.go: error: type %T not found in typeContent\n", outConverted)
    }
  }

  return buffer.Bytes()
}
/*
func(el Content) String() string {
  out := ``

  for _, v := range el{
    switch outConverted := v.(type) {
    case string:
      out += v.(string)
    case HtmlElementDiv:
      out += outConverted.ToHtml()
    /*case String:
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
}*/