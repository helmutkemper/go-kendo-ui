package telerik

import (
  //"fmt"
  "fmt"
  "bytes"
  "sort"
)

type Content []interface{}

func(el Content) ToHtml() []byte {
  var buffer bytes.Buffer

  keys := make([]int, 0)
  for k := range el {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    switch outConverted := el[k].(type) {
    case string:
      buffer.WriteString( el[k].(string) )
    case HtmlElementDiv:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputGeneric:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputNumber:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementSpan:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormLabel:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputText:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormLegend:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormFieldSet:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementForm:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormSelect:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormTextArea:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormMeter:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormButton:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormDataList:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiNumericTextBox:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiComboBox:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiDialog:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementScript:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiMultiSelect:
      buffer.Write( outConverted.ToHtml() )
    case HtmlScriptType:
      buffer.WriteString( outConverted.String() )
    /*case HtmlElementFormOutput:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormProgress:
      buffer.Write( outConverted.ToHtml() )*/

    default:
      fmt.Printf("\n\n\nToHtml() - typeContent.go: error: type %T not found in typeContent\n\n\n", outConverted)
    }
  }

  return buffer.Bytes()
}
func(el Content) ToJavaScript() []byte {
  var buffer bytes.Buffer
      //buffer.WriteString( "\n" )

  keys := make([]int, 0)
  for k := range el {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    switch outConverted := el[k].(type) {
    case string:
    case KendoUiDialog:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDateTimePicker:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDatePicker:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDateInput:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiContextMenu:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiConfirm:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiComboBox:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiColorPicker:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiColorPalette:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiCalendar:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiButton:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiMultiSelect:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiAutoComplete:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDraggable:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDropDownList:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDropTarget:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDropTargetArea:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementDiv:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementFormLabel:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiNumericTextBox:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementFormSelect:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementScript:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementFormButton:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementSpan:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    default:
      fmt.Printf("\n\n\nToJavaScript() - typeContent.go: error: type %T not found in typeContent\n\n\n", outConverted)
    }

    if buffer.Len() > 0 {
      buffer.WriteString( "\n" )
    }
  }

  return buffer.Bytes()
}