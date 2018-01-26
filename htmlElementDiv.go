package telerik

// The HTML Content Division element (<div>) is the generic container for flow content. It has no effect on the content
// or layout until styled using CSS. As a "pure" container, the <div> element does not inherently represent anything.
// Instead, it's used to group content so it can be easily styled using the class or id attributes, marking a section of
// a document as being written in a different language (using the lang attribute), and so on.
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