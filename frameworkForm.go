package telerik

import (
  "fmt"
  "bytes"
  "html/template"
)

type elementProcessed struct {
  Label               string
  Input               string
}

type FrameworkInput struct {
  Id                  String
  Label               String
  Title               String
  PlaceHolder         String
  ValidationMessage   String
  Input               interface{}
}

type FrameworkForm struct {
  Template            string
  Class               String
  Form                HtmlElementForm
  Elements            []FrameworkInput
  elementsProcessed   []elementProcessed
}
func(el *FrameworkForm) ToScript() string {
  var out bytes.Buffer
  for i := 0; i != len( el.Elements ); i += 1 {
    switch elementCast := el.Elements[i].Input.(type) {
    case KendoUiAutoComplete:
      elementCast.HtmlId = el.Elements[i].Id
      out.WriteString( elementCast.String() )
    }
  }
  return out.String()
}
func(el *FrameworkForm) ToForm() string {
  var out bytes.Buffer

  el.elementsProcessed = make([]elementProcessed, 0)

  for i := 0; i != len( el.Elements ); i += 1 {
    element := elementProcessed{}

    label := HtmlElementFormLabel{
      Content: Content{
        el.Elements[i].Label,
      },
      For: el.Elements[i].Id,
      Global:HtmlGlobalAttributes{
        Title: el.Elements[i].Title,
      },
    }
    switch elementCast := el.Elements[i].Input.(type) {
    case KendoUiAutoComplete:
      elementNew := HtmlInputSearch{}
      elementNew.Global.Id     = el.Elements[i].Id
      elementNew.Global.Class  = "k-textbox"
      elementNew.Name          = el.Elements[i].Id
      elementNew.PlaceHolder   = el.Elements[i].PlaceHolder

      if el.Elements[i].ValidationMessage != "" {
        label.Global.Class += " required"

        elementNew.Required = TRUE
        elementNew.Global.Extra = map[string]string{
          "validationMessage":el.Elements[i].ValidationMessage.String(),
        }
      }

      element.Input = elementNew.String()
    case HtmlInputText:
      elementCast.Global.Id    = el.Elements[i].Id
      elementCast.Global.Class = "k-textbox"
      elementCast.Name         = el.Elements[i].Id
      elementCast.PlaceHolder  = el.Elements[i].PlaceHolder

      if el.Elements[i].ValidationMessage != "" {
        label.Global.Class += " required"

        elementCast.Required = TRUE
        elementCast.Global.Extra = map[string]string{
          "validationMessage":el.Elements[i].ValidationMessage.String(),
        }
      }

      element.Input = elementCast.String()
    }

    element.Label = label.String()

    el.elementsProcessed = append( el.elementsProcessed, element )
  }

  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(el.Template))
  err := tmpl.Execute(&out, el.elementsProcessed)
  if err != nil {
    fmt.Println(err.Error())
  }

  form := el.Form
  form.Content = Content{
    out.String(),
  }

  return form.String()
}