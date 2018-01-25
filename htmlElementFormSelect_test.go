package telerik

import "fmt"

func ExampleHtmlElementFormSelect_String() {
  el := HtmlElementFormSelect{
    Name: "selectName",
    Options: map[string]string{
      "key_1": "value_1",
      "key_2": "value_2",
      "key_3": "value_3",
    },
  }

  fmt.Printf( "%v\n", el.String() )

  // Output:
  // <select  name="selectName" ><option value="key_1">key_1</option><option value="key_2">key_2</option><option value="key_3">key_3</option></select>
}
