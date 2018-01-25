package telerik

type HtmlElementDiv struct{

  /*
  Content inside html tag
  */
  Content                     string

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementDiv)String() string {
  return `<div ` + el.Global.String() + `>` + el.Content + `</div>`
}