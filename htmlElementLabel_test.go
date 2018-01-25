package telerik

import "fmt"

func ExampleHtmlElementFormLabel_String() {
  el := HtmlElementFormLabel{
    Form: "name",
    Content: Content{
      "Name",
    },
  }

  fmt.Printf("%v\n", el.String())

  // Output:
  // <label  form="name" >Name</label>
}
