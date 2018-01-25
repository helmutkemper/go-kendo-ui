package telerik

import "fmt"

func ExampleHtmlElementDiv_String() {
  el := HtmlElementDiv{
    Global: HtmlGlobalAttributes{
      Id: "divId",
      Class: "test",
    },
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

  fmt.Printf( "%v", el.String() )

  // Output:
  // <div Class="test" Id="divId" ><label  form="name" >label_1</label><input Id="name"  type="text"  name="name" ></div>
}