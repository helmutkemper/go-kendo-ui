package telerik

import (
  "reflect"
  "fmt"
  "errors"
  "time"
)

type ToJavaScriptConverter struct {}

func(el *ToJavaScriptConverter) ToTelerikJavaScript( element reflect.Value, kendoElement string ) (string, error) {
  typeOfT := element.Type()

  passHtmlId := false
  for i := 0; i < element.NumField(); i += 1 {
    field := element.Field(i)
    typeField := element.Type().Field(i)
    tag := typeField.Tag

    if tag.Get("template") == "htmlId" {
      passHtmlId = true
      fmt.Printf(`$("#%v")%v({`, tag.Get("template"), kendoElement)

      continue
    }

    switch field.Type().String() {
    case "telerik.String":
      if field.Interface().(String) == "" {
        continue
      }

      fmt.Printf(`%v: "%v",`, tag.Get("template"), field.Interface())
      //fmt.Printf(`{{if (.%v) ne "" }}%v: {{.%v}},{{end}}`, typeOfT.Field(i).Name, tag.Get("template"), typeOfT.Field(i).Name)

    case "[]time.Time":
      length := len( field.Interface().([]time.Time) )
      if length > 0 {
        fmt.Printf("%v: [", tag.Get("template"))
      }
      for k, v := range field.Interface().([]time.Time) {
        fmt.Printf("new Date(%v, %v, %v, %v, %v, %v, 0)", v.Year(), int( v.Month() ), v.Day(), v.Hour(), v.Minute(), v.Second())

        if k != length -1 {
          fmt.Printf(",")
        }
      }
      if length > 0 {
        fmt.Printf("]")
      }
    }
  }

  fmt.Printf("});")

  if passHtmlId == false {
    return "", errors.New("no html element id found")
  }



  fmt.Print("\n\n\n\n\n\n\n")

  for i := 0; i < element.NumField(); i += 1 {
    field := element.Field(i)
    typeField := element.Type().Field(i)
    tag := typeField.Tag
    fmt.Printf("%d: %s %s = %v  template: ''%v''\n", i, typeOfT.Field(i).Name, field.Type(), field.Interface(), tag.Get("template"))
  }

  return "", nil
}