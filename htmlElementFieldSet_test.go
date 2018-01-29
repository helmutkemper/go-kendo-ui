package telerik

import "fmt"

func ExampleHtmlElementFormFieldSet_String() {
  el := HtmlElementFormFieldSet{
    Global: HtmlGlobalAttributes{
      Id: "fieldSetId",
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
  // <fieldset  id="fieldSetId" ><label  form="name" >label_1</label><input  id="name"  type="text"  name="name" ></fieldset>
}
