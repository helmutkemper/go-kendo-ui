package telerik

type HtmlElementDiv struct{

  /*
  Content inside html tag
  */
  Content                     Content

  Global                      HtmlGlobalAttributes
}
func(el *HtmlElementDiv)String() string {
  return `<div ` + el.Global.String() + `>` + el.Content.String() + `</div>`
}