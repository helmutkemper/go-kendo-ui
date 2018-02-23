# telerik

[![Build Status](https://travis-ci.org/helmutkemper/KendoUI.svg?branch=master)](https://travis-ci.org/helmutkemper/KendoUI)

telerik kendo ui for golang

```
package main

import (
  "fmt"
  "html/template"
  html "github.com/helmutkemper/telerik"
  "os"
)

func main() {

  a := html.KendoUiAutoComplete{
    HtmlId: "autocomplete",
    DataSource: `["Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City"]`,
  }

  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(html.GetTemplate()))
  err := tmpl.ExecuteTemplate(os.Stdout, "KendoUiAutoComplete", a)
  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Printf("\n\n%v\n", &html.HtmlInputButton{
    Name: "click",
    Value:"Click me",
    Global: html.HtmlGlobalAttributes{
      Id:"html",
      Data: map[string]string{"um":"un value", "dois":"dois value"},
      DropZone: html.COPY,
    },
  })

  fmt.Printf("\n\n%v\n", &html.HtmlInputCheckBox{
    Name: "box",
    Value:"box checked",
    Global: html.HtmlGlobalAttributes{
      Id:"box",
    },
  })

  fmt.Printf("\n\n%v\n", &html.HtmlInputDate{
    Name: "box",
    Value:"23/1/2017",
    ValueAsNumber: html.TRUE,
    Global: html.HtmlGlobalAttributes{
      Id:"date",
    },
  })

  fmt.Printf("\n\n%v\n", &html.HtmlInputEmail{
    Name: html.NAMES_FOR_AUTOCOMPLETE_NAME,
    AutoComplete: html.TRUE,
    PlaceHolder: "Digite o seu e-mail",
    Value:"eu@eu.com",
    Size: 20,
    MinLength: 20,
    MaxLength: 100,
    Global: html.HtmlGlobalAttributes{
      Id:"email",
    },
  })

  fmt.Printf("\n\n%v\n\n", &html.HtmlInputFile{
    Name: "file",
    Accept: html.StringArr{".png", ".jpg", ".gif"},
    Required: html.TRUE,
    Global: html.HtmlGlobalAttributes{
      Id:"file",
    },
  })





  f := html.FrameworkForm{
    Template: `<ul>{{range $k, $v := .}}<li>{{$v.Label | safeHTML}}:{{$v.Input | safeHTML}}</li>{{end}}</ul>`,
    Form: html.HtmlElementForm{
      Action: "./cadastrar",
      Method: "get",
    },
    Elements: []html.FrameworkInput{
      {
        Id: "nome",
        Label: "Nome",
        ValidationMessage: "Por favor, digite o seu nome completo.",
        Title: "Digite o seu nome",
        PlaceHolder: "Nome completo",
        Input: html.HtmlInputText{},
      },
      {
        Id: "endereco",
        Label: "Endereço",
        ValidationMessage: "Por favor, digite seu nome completo.",
        PlaceHolder: "Endereço",
        Title: "Digite o seu endereço",
        Input: html.KendoUiAutoComplete{
          DataSource: []string{"Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City"},
        },
      },
    },
  }

  fmt.Printf("%v\n", f.ToForm())
  fmt.Printf("<script>%v</script>\n", f.ToScript())
}
```
