package telerik

import (
  "reflect"
  "fmt"
  "time"
  "bytes"
  "strconv"
)

type ToJavaScriptConverter struct {}

/*
keys := make([]int, 0)
for k, _ := range romanNumeralDict {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println(k, romanNumeralDict[k])
}
*/

func(el *ToJavaScriptConverter) getTagData( tag reflect.StructTag ) (string, string) {
  var tagName, tagValue string

  tagName  = "htmlAttr"
  tagValue = tag.Get(tagName)
  if tagValue != "" {
    return tagName, tagValue
  }

  tagName  = "htmlAttrSet"
  tagValue = tag.Get(tagName)
  if tagValue != "" {
    return tagName, tagValue
  }

  tagName  = "htmlAttrOnOff"
  tagValue = tag.Get(tagName)
  if tagValue != "" {
    return tagName, tagValue
  }

  tagName  = "htmlAttrTrueFalse"
  tagValue = tag.Get(tagName)
  if tagValue != "" {
    return tagName, tagValue
  }

  return "", ""
}

func(el *ToJavaScriptConverter) ToTelerikHtml( element reflect.Value ) []byte {
  var buffer bytes.Buffer
  var tagName, tagValue string

  //typeOfT := element.Type()

  for i := 0; i < element.NumField(); i += 1 {
    field := element.Field(i)
    typeField := element.Type().Field(i)
    tag := typeField.Tag

    tagName, tagValue = el.getTagData( tag )

    if tagValue == "-" {
      continue
    }

    switch field.Type().String() {
      case "string":
        if field.Interface().(string) == "" {
          continue
        }

        buffer.WriteString(` ` + tagValue + `="` + field.Interface().(string) + `"`)

      case "map[string]string":
        if len( field.Interface().(map[string]string) ) == 0 {
          continue
        }

        for k, v := range field.Interface().(map[string]string){
          if tagName == "data" {
            buffer.WriteString(` ` + k + `="` + v + `"`)
          } else {
            buffer.WriteString(` data-` + k + `="` + v + `"`)
          }
        }

      case "map[string]interface {}":
        if len( field.Interface().(map[string]interface{}) ) == 0 {
          continue
        }

        for k, v := range field.Interface().(map[string]interface{}){
          switch v.(type) {
          case string:
            if field.Interface().(string) == "" {
              continue
            }

            buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(string) + `",`)

          case []string:
            if len( field.Interface().([]string) ) == 0 {
              continue
            }

            buffer.WriteString(tag.Get("jsObject") + `: [`)
            for _, v := range field.Interface().([]string) {
              buffer.WriteString(`"`)
              buffer.WriteString( v )
              buffer.WriteString(`",`)
            }
            buffer.WriteString(`],`)

          case int:
            if v.(int) == 0 {
              continue
            }

            buffer.WriteString(` ` + k + `="` + strconv.Itoa( v.(int) ) + `"`)

          case int64:
            if v.(int64) == 0 {
              continue
            }

            buffer.WriteString(` ` + k + `="` + strconv.FormatInt( v.(int64), 64 ) + `,`)

          case float32:
            if v.(float32) == 0 {
              continue
            }

            buffer.WriteString(` ` + k + `="` + strconv.FormatFloat( float64( v.(float32) ), 'E', -1, 32) + `,`)

          case float64:
            if v.(float64) == 0 {
              continue
            }

            buffer.WriteString(` ` + k + `="` + strconv.FormatFloat( v.(float64), 'E', -1, 64) + `,`)
          }
        }

      case "telerik.TypeHtmlDropZone":
        if field.Interface().(TypeHtmlDropZone).String() == "" {
          continue
        }

        buffer.WriteString(` ` + tagValue + `="` + field.Interface().(TypeHtmlDropZone).String() + `"`)

      case "telerik.Warp":
        if field.Interface().(Warp).String() == "" {
          continue
        }

        buffer.WriteString(` ` + tagValue + `="` + field.Interface().(Warp).String() + `"`)

      case "telerik.AutoCapitalize":
        if field.Interface().(AutoCapitalize).String() == "" {
          continue
        }

        buffer.WriteString(` ` + tagValue + `="` + field.Interface().(AutoCapitalize).String() + `"`)

      case "telerik.Boolean":
        if field.Interface().(Boolean) == 0 {
          continue
        }

        if tagName == "htmlAttrSet" && field.Interface().(Boolean) == TRUE {
          buffer.WriteString(` ` + tagValue)
          continue
        }

        if tagName == "htmlAttrOnOff" {
          buffer.WriteString(` ` + tagValue + `="` + field.Interface().(Boolean).OnOff() + `"`)
          continue
        }

        if tagName == "htmlAttrOnOff" {
          buffer.WriteString(` ` + tagValue + `="` + field.Interface().(Boolean).String() + `"`)
          continue
        }

      case "int":
        if field.Interface().(int) == 0 {
          continue
        }

        buffer.WriteString(` ` + tagValue + `="` + strconv.Itoa( field.Interface().(int) ) + `"`)

    default:
      fmt.Printf("\nhtmlTag(): %d: %s - %s = %v\n", i, field.Type(), field.Interface(), tagName)
    }
  }

  return buffer.Bytes()
}

func(el *ToJavaScriptConverter) ToAceJavaScript( element reflect.Value ) ([]byte, error) {
  var buffer bytes.Buffer

  //typeOfT := element.Type()

  for i := 0; i < element.NumField(); i += 1 {
    field := element.Field(i)
    typeField := element.Type().Field(i)
    tag := typeField.Tag

    if tag.Get("jsSetFunction") == "-" {
      continue
    }

    switch field.String() {

    }
  }

  return buffer.Bytes(), nil
}

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

    switch field.Type().String() {
    //------------------------------------------------------------------------------------------------------------------
    case "interface {}":

      if field.Interface() == nil {
        continue
      }

      switch field.Interface().(type) {
      case *KendoDataSource:
        if field.Interface().(*KendoDataSource) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: `)
        buffer.Write( field.Interface().(*KendoDataSource).ToJavaScript() )

      case *KendoAnimation:
        if field.Interface().(*KendoAnimation) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: { `)
        buffer.Write( field.Interface().(*KendoAnimation).ToJavaScript() )
        buffer.WriteString(`},`)

      case *KendoPopup:
        if field.Interface().(*KendoPopup) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: { `)
        buffer.Write( field.Interface().(*KendoPopup).ToJavaScript() )
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

      case []string:
        if len( field.Interface().([]string) ) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: [`)
        for _, v := range field.Interface().([]string) {
          buffer.WriteString(`"`)
          buffer.WriteString( v )
          buffer.WriteString(`",`)
        }
        buffer.WriteString(`],`)

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

      case *KendoOfflineStorage:
        if field.Interface().(*KendoOfflineStorage) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + field.Interface().(*KendoOfflineStorage).String() + `,`)

      case *KendoComplexFilter:
        if field.Interface().(*KendoComplexFilter) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: { `)
        buffer.Write( field.Interface().(*KendoComplexFilter).ToJavaScript() )
        buffer.WriteString(`},`)

      case *[]KendoComplexFilter:
        if field.Interface().(*[]KendoComplexFilter) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: [`)
        for _, v := range *field.Interface().(*[]KendoComplexFilter) {
          buffer.WriteString(`{`)
          buffer.Write( v.ToJavaScript() )
          buffer.WriteString(`},`)
        }
        buffer.WriteString(`],`)

      case *[]KendoGroup:
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

      case *map[string]interface{}:
        if field.Interface().(*map[string]interface{}) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: {`)
        for k, v := range *field.Interface().(*map[string]interface{}) {
          switch v.(type) {
          case string:
            buffer.WriteString(`"` + k + `": ` + `"` + v.(string) + `",`)

          case int:
            buffer.WriteString(`"` + k + `": ` + strconv.Itoa( v.(int) ) + `,`)

          case int64:
            buffer.WriteString(`"` + k + `": ` + strconv.FormatInt( v.(int64), 64 ) + `,`)

          case float32:
            buffer.WriteString(`"` + k + `": ` + strconv.FormatFloat( float64( v.(float32) ), 'E', -1, 32) + `,`)

          case float64:
            buffer.WriteString(`"` + k + `": ` + strconv.FormatFloat( v.(float64), 'E', -1, 64) + `,`)
          }
        }
        buffer.WriteString(`},`)

      case *KendoCreate:
        if field.Interface().(*KendoCreate) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: { `)
        buffer.Write( field.Interface().(*KendoCreate).ToJavaScript() )
        buffer.WriteString(`},`)

      case *[]map[string]interface{}:
        if field.Interface().(*[]map[string]interface{}) == nil {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: [`)
        for _, mapStringInterface := range *field.Interface().(*[]map[string]interface{}) {
          buffer.WriteString(`{`)
          for k, v := range mapStringInterface{

            buffer.WriteString( `"` + k + `": ` )
            switch v.(type) {
              case string:
                buffer.WriteString(`"` + v.(string) + `",`)

              case int:
                buffer.WriteString(strconv.Itoa( v.(int) ) + `,`)

              case int64:
                buffer.WriteString(strconv.FormatInt( v.(int64), 64 ) + `,`)

              case float32:
                buffer.WriteString(strconv.FormatFloat( float64( v.(float32) ), 'E', -1, 32) + `,`)

              case float64:
                buffer.WriteString(strconv.FormatFloat( v.(float64), 'E', -1, 64) + `,`)

              case bool:
                buffer.WriteString( strconv.FormatBool( v.(bool) ) + `,` )

              case []map[string]interface{}:
                buffer.WriteString(`[`)

                for _, mapArr := range v.([]map[string]interface{}){
                  buffer.WriteString(`{`)

                  for k, v := range mapArr{
                    buffer.WriteString( `"` + k + `": ` )
                    switch v.(type){
                      case string:
                        buffer.WriteString(`"` + v.(string) + `",`)

                      case int:
                        buffer.WriteString(strconv.Itoa( v.(int) ) + `,`)

                      case int64:
                        buffer.WriteString(strconv.FormatInt( v.(int64), 64 ) + `,`)

                      case float32:
                        buffer.WriteString(strconv.FormatFloat( float64( v.(float32) ), 'E', -1, 32) + `,`)

                      case float64:
                        buffer.WriteString(strconv.FormatFloat( v.(float64), 'E', -1, 64) + `,`)

                      case bool:
                        buffer.WriteString( strconv.FormatBool( v.(bool) ) + `,` )
                    }
                  }
                  buffer.WriteString(`},`)
                }
                buffer.WriteString(`],`)
            }
          }
          buffer.WriteString(`},`)
        }
        buffer.WriteString(`],`)

      case Boolean:
        if field.Interface().(Boolean) == 0 {
          continue
        }

        buffer.WriteString(tag.Get("jsObject") + `: ` + field.Interface().(Boolean).String() + `,`)

      default:
        fmt.Printf("\n\n\n\ninterface{}: %d: %s %s = %v  template: ''%v''\n\n\n\n\n", i, typeOfT.Field(i).Name, field.Type(), field.Interface(), tag.Get("jsObject"))
      }
    //---------------------------------------------------------------------------------------------------------------

    case "*telerik.KendoConfirmMessages":
      if field.Interface().(*KendoConfirmMessages) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoConfirmMessages).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoColorMessages":
      if field.Interface().(*KendoColorMessages) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoColorMessages).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoTileSize":
      if field.Interface().(*KendoTileSize) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoTileSize).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoCreate":
      if field.Interface().(*KendoCreate) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoCreate).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoRead":
      if field.Interface().(*KendoRead) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoRead).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoDestroy":
      if field.Interface().(*KendoDestroy) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoDestroy).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoUpdate":
      if field.Interface().(*KendoUpdate) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoUpdate).ToJavaScript() )
      buffer.WriteString(`},`)

    case "telerik.HtmlMethod":
      if field.Interface().(HtmlMethod) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(HtmlMethod).String() + `",`)

    case "telerik.KendoTimeDepth":
      if field.Interface().(KendoTimeDepth) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoTimeDepth).String() + `",`)

    case "telerik.KendoTypeDataJSon":
      if field.Interface().(KendoTypeDataJSon) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoTypeDataJSon).String() + `",`)

    case "telerik.KendoTypeData":
      if field.Interface().(KendoTypeData) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoTypeData).String() + `",`)

    case "telerik.KendoContextMenuDirection":
      if field.Interface().(KendoContextMenuDirection) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoContextMenuDirection).String() + `",`)

    case "telerik.HtmlMouseEvent":
      if field.Interface().(HtmlMouseEvent) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(HtmlMouseEvent).String() + `",`)

    case "telerik.KendoOrientation":
      if field.Interface().(KendoOrientation) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoOrientation).String() + `",`)

    case "telerik.KendoAggregate":
      if field.Interface().(KendoAggregate) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoAggregate).String() + `",`)

    case "*[]map[string]interface {}":
      if field.Interface().(*[]map[string]interface{}) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, mapStringInterface := range *field.Interface().(*[]map[string]interface{}) {

        buffer.WriteString(`{`)
        for k, v := range mapStringInterface{

          buffer.WriteString( `"` + k + `": ` )
          switch v.(type) {
          case string:
            buffer.WriteString(`"` + v.(string) + `",`)

          case int:
            buffer.WriteString(strconv.Itoa( v.(int) ) + `,`)

          case int64:
            buffer.WriteString(strconv.FormatInt( v.(int64), 64 ) + `,`)

          case float32:
            buffer.WriteString(strconv.FormatFloat( float64( v.(float32) ), 'E', -1, 32) + `,`)

          case float64:
            buffer.WriteString(strconv.FormatFloat( v.(float64), 'E', -1, 64) + `,`)
          }
        }
        buffer.WriteString(`},`)

      }
      buffer.WriteString(`],`)
      
    case "*[]telerik.KendoActions":
      if len( *( field.Interface().( *[]KendoActions) ) ) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, v := range *( field.Interface().( *[]KendoActions) ) {
        buffer.WriteString(`{`)
        buffer.Write( v.ToJavaScript() )
        buffer.WriteString(`},`)
      }
      buffer.WriteString(`],`)

    case "[]string":
      if len(field.Interface().([]string)) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, v := range field.Interface().([]string) {
        buffer.WriteString(`"`)
        buffer.WriteString( v )
        buffer.WriteString(`",`)
      }
      buffer.WriteString(`],`)

    case "telerik.KendoWeekDays":
      if len(field.Interface().(KendoWeekDays)) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, v := range field.Interface().(KendoWeekDays) {
        buffer.WriteString(`"`)
        buffer.WriteString( v.String() )
        buffer.WriteString(`",`)
      }
      buffer.WriteString(`],`)

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

    case "*[]telerik.KendoFilters":
      if field.Interface().(*[]KendoFilters) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: [`)
      for _, v := range *field.Interface().(*[]KendoFilters) {
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

    case "*telerik.KendoSort":
      if field.Interface().(*KendoSort) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoSort).ToJavaScript() )
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

    case "*telerik.KendoTransport":
      if field.Interface().(*KendoTransport) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: { `)
      buffer.Write( field.Interface().(*KendoTransport).ToJavaScript() )
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

    case "*telerik.KendoSchema":
      if field.Interface().(*KendoSchema) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: {`)
      buffer.Write( field.Interface().(*KendoSchema).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoSignalr":
      if field.Interface().(*KendoSignalr) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: {`)
      buffer.Write( field.Interface().(*KendoSignalr).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoDataModel":
      if field.Interface().(*KendoDataModel) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: {`)
      buffer.Write( field.Interface().(*KendoDataModel).ToJavaScript() )
      buffer.WriteString(`},`)

    case "*telerik.KendoMessages":
      if field.Interface().(*KendoMessages) == nil {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: {`)
      buffer.Write( field.Interface().(*KendoMessages).ToJavaScript() )
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

    case "telerik.KendoLogic":
      if field.Interface().(KendoLogic) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoLogic).String() + `",` )

    case "telerik.KendoType":
      if field.Interface().(KendoType) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoType).String() + `",` )

    case "telerik.KendoButtonLayout":
      if field.Interface().(KendoButtonLayout) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoButtonLayout).String() + `",` )

    case "telerik.KendoDirection":
      if field.Interface().(KendoDirection) == 0 {
        continue
      }

      buffer.WriteString(tag.Get("jsObject") + `: "` + field.Interface().(KendoDirection).String() + `",` )

    case "telerik.Content":
      buffer.Write( field.Interface().(Content).Bytes() )

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

    if tag.Get("jsObject") == "-" {
      continue
    }

    switch field.Type().String() {
      case "telerik.KendoAggregate": continue
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
      case "*[]telerik.KendoFilters": continue
      case "telerik.KendoLogic": continue
      case "telerik.KendoOperator": continue
      case "*[]telerik.KendoAggregates": continue
      case "telerik.KendoDirection": continue
      case "*telerik.KendoSchema": continue
      case "*telerik.KendoDataModel": continue
      case "telerik.KendoTypeData": continue
      case "*telerik.KendoSort": continue
      case "*telerik.KendoTransport": continue
      case "telerik.KendoType": continue
      case "telerik.KendoTypeDataJSon": continue
      case "telerik.HtmlMethod": continue
      case "*[]map[string]interface {}": continue
      case "*telerik.KendoDestroy": continue
      case "*telerik.KendoRead": continue
      case "*telerik.KendoUpdate": continue
      case "*telerik.KendoSignalr": continue
      case "[]string": continue
      case "*telerik.KendoTileSize": continue
      case "*telerik.KendoColorMessages": continue
      case "*telerik.KendoConfirmMessages": continue
      case "telerik.KendoContextMenuDirection": continue
      case "telerik.KendoOrientation": continue
      case "telerik.HtmlMouseEvent": continue
      case "telerik.KendoWeekDays": continue
      case "telerik.KendoTimeDepth": continue
      case "*[]telerik.KendoActions": continue
      case "telerik.KendoButtonLayout": continue
      case "telerik.Content": continue
      case "*telerik.KendoMessages": continue
    }

    fmt.Printf("\n\n\n\n%d: %s %s = %v  template: ''%v''\n\n\n\n\n", i, typeOfT.Field(i).Name, field.Type(), field.Interface(), tag.Get("jsObject"))
  }

  return buffer.Bytes(), nil
}