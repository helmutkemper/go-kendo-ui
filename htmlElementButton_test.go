package telerik

import "fmt"

func ExampleHtmlElementFormButton_String() {
  el := HtmlElementFormButton{
    Global: HtmlGlobalAttributes{
      Id: "button",
    },
    Name: "button",
    Content: Content{
      "click-me",
    },
    Disabled: TRUE,
  }

  fmt.Printf( "%v", el.String() )

  // Output:
  // <button  id="button"  name="button"  disabled >click-me</button>
}
