package telerik

import "fmt"

func ExampleHtmlElementFormDataList_String() {
  el := HtmlElementFormDataList{
    Global: HtmlGlobalAttributes{
      Id: "dataListId",
    },
    Name: "dataListName",
    Options: map[string]string{
      "key_1":"value_1",
      "key_2":"value_2",
      "key_3":"value_3",
    },
    Value: "key_2",
  }

  fmt.Printf( "%v", el.String() )

  // Output:
  // <input Id="dataListId"  list="dataListName"  value="key_2" ><datalist id="dataListName"><option value="key_1">value_1</option><option value="key_2">value_2</option><option value="key_3">value_3</option></datalist>
}
