package telerik

import "fmt"

func ExampleHtmlElementForm_String() {
  el := HtmlElementForm{
    Name: "form",
    Method: "get",
    Action: "./index.cpp",
    Content: Content{
      HtmlElementFormLabel{
        Form: "name",
        Content: Content{
          "label_1",
        },
      },
      HtmlInputText{
        Global: HtmlGlobalAttributes{
          Id: "name",
        },
        Name: NAMES_FOR_AUTOCOMPLETE_NAME,
      },
    },
  }

  fmt.Printf( "%v\n", el.String() )

  // Output:
  // <form  name="form"  action="./index.cpp"  method="get" ><label  form="name" >label_1</label><input Id="name"  type="text"  name="name" ></form>
}
