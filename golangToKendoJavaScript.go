package telerik

import (
  "reflect"
  "fmt"
  "time"
  "bytes"
  "strconv"
)

type ToJavaScriptConverter struct {}

func(el *ToJavaScriptConverter) ToTelerikJavaScript( element reflect.Value ) ([]byte, error) {
  var buffer bytes.Buffer

  typeOfT := element.Type()

  for i := 0; i < element.NumField(); i += 1 {
    field := element.Field(i)
    typeField := element.Type().Field(i)
    tag := typeField.Tag

    if tag.Get("jsObject") == "-" {
      continue
    }

    if tag.Get("jsObject") == "htmlId" {
      continue
    }

    switch field.Type().String() {
    case "interface {}":

      if field.Interface() == nil {
        continue
      }

      switch field.Interface().(type) {
      case *KendoAnimation:
        if field.Interface().(*KendoAnimation) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: { `)

        buffer.Write( field.Interface().(*KendoAnimation).ToJavaScript() )

        buffer.WriteString(`},`)

      case *KendoVirtual:
        if field.Interface().(*KendoVirtual) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: { `)

        buffer.Write( field.Interface().(*KendoVirtual).ToJavaScript() )

        buffer.WriteString(`},`)

      case *JavaScript:
        if field.Interface().(*JavaScript) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + field.Interface().(*JavaScript).Code + `,`)

      case string:
        if field.Interface().(string) == "" {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(string) + `",`)

      case int:
        if field.Interface().(int) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + strconv.Itoa( field.Interface().(int) ) + `,`)

      case int64:
        if field.Interface().(int64) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + strconv.FormatInt( field.Interface().(int64), 64 ) + `,`)

      case float32:
        if field.Interface().(float32) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + strconv.FormatFloat( float64( field.Interface().(float32) ), 'E', -1, 32) + `,`)

      case float64:
        if field.Interface().(float64) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + strconv.FormatFloat( field.Interface().(float64), 'E', -1, 64) + `,`)

      case *OfflineStorage:
        if field.Interface().(*OfflineStorage) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + field.Interface().(*OfflineStorage).String() + `,`)

      case Boolean:
        if field.Interface().(Boolean) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + field.Interface().(Boolean).String() + `,`)

      default:
        fmt.Printf("\n\n\n\ninterface{}: %d: %s %s = %v  template: ''%v''\n\n\n\n\n", i, typeOfT.Field(i).Name, field.Type(), field.Interface(), tag.Get("jsObject"))
      }

    case "telerik.KendoAggregate":
      if field.Interface().(KendoAggregate) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoAggregate).String() + `",`)

    case "*[]telerik.KendoGroup":
      if field.Interface().(*[]KendoGroup) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, v := range *field.Interface().(*[]KendoGroup) {
        buffer.WriteString(`{`)
        buffer.Write( v.ToJavaScript() )
        buffer.WriteString(`},`)
      }
      buffer.WriteString(`],`)

    case "*[]telerik.KendoAggregates":
      if field.Interface().(*[]KendoAggregates) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, v := range *field.Interface().(*[]KendoAggregates) {
        buffer.WriteString(`{`)
        buffer.Write( v.ToJavaScript() )
        buffer.WriteString(`},`)
      }
      buffer.WriteString(`],`)

    case "int":
      if field.Interface().(int) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: ` + strconv.Itoa( field.Interface().(int) ) + `,`)

    case "*telerik.JavaScript":
      if field.Interface().(*JavaScript) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: ` + field.Interface().(*JavaScript).Code + `,`)

    case "string":
      if field.Interface().(string) == "" {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(string) + `",`)

    case "[]map[string]interface {}":
      if len( field.Interface().([]map[string]interface{}) ) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)

      for _, mapData := range field.Interface().([]map[string]interface{}){

        buffer.WriteString(`{`)
        for k, v := range mapData {
          switch v.(type) {
          case string:
            buffer.WriteString( `"` + k + `": "` + v.(string) + `",` )
          case int:
            buffer.WriteString( `"` + k + `": "` + strconv.Itoa(v.(int)) + `",` )
          case int64:
            buffer.WriteString( `"` + k + `": "` + strconv.FormatInt(v.(int64), 64) + `",` )
          case float64:
            buffer.WriteString( `"` + k + `": "` + strconv.FormatFloat(v.(float64), 'E', -1, 64) + `",` )
          case float32:
            buffer.WriteString( `"` + k + `": "` + strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32) + `",` )
          }
        }
        buffer.WriteString(`},`)

      }

      buffer.WriteString(`],`)

    case "telerik.StringArr":
      length := len( field.Interface().(StringArr) )
      if length == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)

      for k, v := range field.Interface().(StringArr) {
        buffer.WriteString( `"` + v + `"` )

        if k != length -1 {
          buffer.WriteString(`,`)
        }
      }

      buffer.WriteString(`],`)

    case "*telerik.KendoGroup":
      if field.Interface().(*KendoGroup) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoGroup).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoAnimation":
      if field.Interface().(*KendoAnimation) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoAnimation).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoClose":
      if field.Interface().(*KendoClose) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoClose).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoOpen":
      if field.Interface().(*KendoOpen) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoOpen).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoPopup":
      if field.Interface().(*KendoPopup) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoPopup).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoVirtual":
      if field.Interface().(*KendoVirtual) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoVirtual).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoMonth":
      if field.Interface().(*KendoMonth) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)

      buffer.Write( field.Interface().(*KendoMonth).ToJavaScript() )

      buffer.WriteString(`},`)

    case "*telerik.KendoCalendarMessages":
      if field.Interface().(*KendoCalendarMessages) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: {`)
      buffer.Write( field.Interface().(*KendoCalendarMessages).ToJavaScript() )
      buffer.WriteString(`},`)

    case "telerik.Boolean":
      if field.Interface().(Boolean) == 0 {
        continue
      }

      if field.Interface().(Boolean) == -1 {
        buffer.WriteString(tag.Get("jsObject") + `: false,`)
      } else {
        buffer.WriteString(tag.Get("jsObject") + `: true,`)
      }

    case "telerik.KendoEffect":
      if field.Interface().(KendoEffect) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoEffect).String() + `",` )

    case "telerik.KendoOperator":
      if field.Interface().(KendoOperator) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoOperator).String() + `",` )

    case "telerik.KendoFilter":
      if field.Interface().(KendoFilter) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoFilter).String() + `",` )

    case "telerik.KendoOrigin":
      if field.Interface().(KendoOrigin) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoOrigin).String() + `",` )

    case "telerik.KendoPosition":
      if field.Interface().(KendoPosition) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoPosition).String() + `",` )

    case "telerik.KendoTagMode":
      if field.Interface().(KendoTagMode) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoTagMode).String() + `",` )

    case "telerik.KendoMapValueTo":
      if field.Interface().(KendoMapValueTo) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoMapValueTo).String() + `",` )

    case "time.Time":
      if field.Interface().(time.Time).String() == `0001-01-01 00:00:00 +0000 UTC` {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: `)

      if int( field.Interface().(time.Time).Month() ) == 0 && field.Interface().(time.Time).Day() == 0 && field.Interface().(time.Time).Hour() == 0 && field.Interface().(time.Time).Minute() == 0 && field.Interface().(time.Time).Second() == 0 {
        buffer.WriteString( `new Date(` + strconv.Itoa( field.Interface().(time.Time).Year() ) + `),`)
        //fmt.Printf("new Date(%v, %v, %v, %v, %v, %v, 0)", v.Year(), int( v.Month() ), v.Day(), v.Hour(), v.Minute(), v.Second())
      } else if field.Interface().(time.Time).Day() == 0 && field.Interface().(time.Time).Hour() == 0 && field.Interface().(time.Time).Minute() == 0 && field.Interface().(time.Time).Second() == 0 {
        buffer.WriteString(`new Date(` + strconv.Itoa( field.Interface().(time.Time).Year() ) + `, ` + strconv.Itoa( int( field.Interface().(time.Time).Month() ) ) + `),` )
      } else if field.Interface().(time.Time).Hour() == 0 && field.Interface().(time.Time).Minute() == 0 && field.Interface().(time.Time).Second() == 0 {
        buffer.WriteString(`new Date(` + strconv.Itoa( field.Interface().(time.Time).Year() ) + `, ` + strconv.Itoa( int( field.Interface().(time.Time).Month() ) ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Day() ) + `),`)
      } else if field.Interface().(time.Time).Minute() == 0 && field.Interface().(time.Time).Second() == 0 {
        buffer.WriteString(`new Date(` + strconv.Itoa( field.Interface().(time.Time).Year() ) + `, ` + strconv.Itoa( int( field.Interface().(time.Time).Month() ) ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Day() ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Hour() ) + `),`)
      } else if field.Interface().(time.Time).Second() == 0 {
        buffer.WriteString(`new Date(` + strconv.Itoa( field.Interface().(time.Time).Year() ) + `, ` + strconv.Itoa( int( field.Interface().(time.Time).Month() ) ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Day() ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Hour() ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Minute() ) + `),`)
      } else {
        buffer.WriteString(`new Date(` + strconv.Itoa( field.Interface().(time.Time).Year() ) + `, ` + strconv.Itoa( int( field.Interface().(time.Time).Month() ) ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Day() ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Hour() ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Minute() ) + `, ` + strconv.Itoa( field.Interface().(time.Time).Second() ) + `),`)
      }

    case "[]time.Time":
      length := len( field.Interface().([]time.Time) )
      if length == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + ": [")

      for k, v := range field.Interface().([]time.Time) {
        if int( v.Month() ) == 0 && v.Day() == 0 && v.Hour() == 0 && v.Minute() == 0 && v.Second() == 0 {
          buffer.WriteString(`new Date(` + strconv.Itoa( v.Year() ) + `)`)
          //fmt.Printf("new Date(%v, %v, %v, %v, %v, %v, 0)", v.Year(), int( v.Month() ), v.Day(), v.Hour(), v.Minute(), v.Second())
        } else if v.Day() == 0 && v.Hour() == 0 && v.Minute() == 0 && v.Second() == 0 {
          buffer.WriteString(`new Date(` + strconv.Itoa( v.Year() ) + `, ` + strconv.Itoa( int( v.Month() ) ) + `)`)
        } else if v.Hour() == 0 && v.Minute() == 0 && v.Second() == 0 {
          buffer.WriteString(`new Date(` + strconv.Itoa( v.Year() ) + `, ` + strconv.Itoa( int( v.Month() ) ) + `, ` + strconv.Itoa( v.Day() ) + `)`)
        } else if v.Minute() == 0 && v.Second() == 0 {
          buffer.WriteString(`new Date(` + strconv.Itoa( v.Year() ) + `, ` + strconv.Itoa( int( v.Month() ) ) + `, ` + strconv.Itoa( v.Day() ) + `, ` + strconv.Itoa( v.Hour() ) + `)`)
        } else if v.Second() == 0 {
          buffer.WriteString(`new Date(` + strconv.Itoa( v.Year() ) + `, ` + strconv.Itoa( int( v.Month() ) ) + `, ` + strconv.Itoa( v.Day() ) + `, ` + strconv.Itoa( v.Hour() ) + `, ` + strconv.Itoa( v.Minute() ) + `)`)
        } else {
          buffer.WriteString(`new Date(` + strconv.Itoa( v.Year() ) + `, ` + strconv.Itoa( int( v.Month() ) ) + `, ` + strconv.Itoa( v.Day() ) + `, ` + strconv.Itoa( v.Hour() ) + `, ` + strconv.Itoa( v.Minute() ) + `, ` + strconv.Itoa( v.Second() ) + `)`)
        }

        if k != length -1 {
          buffer.WriteString(",")
        }
      }

      buffer.WriteString("],")
    }
  }

  // fixme: remover isto - inicio
  //if kendoElement != "" {
    fmt.Print("\n\n\n\n\n\n\n")
  //}
  // fixme: remover isto - fim

  for i := 0; i < element.NumField(); i += 1 {
    field := element.Field(i)
    typeField := element.Type().Field(i)
    tag := typeField.Tag

    switch field.Type().String() {
      case "interface {}": continue
      case "telerik.TypeAggregate": continue
      case "*[]telerik.KendoAggregate": continue
      case "*telerik.KendoCalendarMessages": continue
      case "[]time.Time": continue
      case "time.Time": continue
      case "string": continue
      case "telerik.Boolean": continue
      case "telerik.KendoEffect": continue
      case "*telerik.ToJavaScriptConverter": continue
      case "*telerik.KendoMonth": continue
      case "*telerik.JavaScript": continue
      case "[]map[string]interface {}": continue
      case "*telerik.KendoAnimation": continue
      case "*telerik.KendoClose": continue
      case "*telerik.KendoOpen": continue
      case "*telerik.KendoPopup": continue
      case "*telerik.KendoVirtual": continue
      case "telerik.KendoFilter": continue
      case "telerik.KendoOrigin": continue
      case "telerik.KendoPosition": continue
      case "telerik.KendoTagMode": continue
      case "telerik.KendoMapValueTo": continue
      case "int": continue
      case "*[]telerik.KendoGroup": continue
    }

    fmt.Printf("\n\n\n\n%d: %s %s = %v  template: ''%v''\n\n\n\n\n", i, typeOfT.Field(i).Name, field.Type(), field.Interface(), tag.Get("jsObject"))
  }

  return buffer.Bytes(), nil
}