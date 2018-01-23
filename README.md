# telerik
telerik kendo ui for golang

```
package main

import (
  "fmt"
  "html/template"
  "github.com/helmutkemper/telerik"
  "os"
)

func main() {

  a := telerik.KendoUiAutoComplete{
    HtmlId: "autocomplete",
    Animation: &telerik.KendoAnimation{
      Open: &telerik.KendoOpen{
        Duration: 800,
      },
      Close: &telerik.KendoClose{
        Duration: 800,
      },
    },
    DataSource: `["Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City"]`,
  }

  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(telerik.GetTemplate()))
  err := tmpl.ExecuteTemplate(os.Stdout, "KendoUiAutoComplete", a)
  if err != nil {
    fmt.Println(err.Error())
  }
}
```
