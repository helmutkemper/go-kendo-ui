package telerik

import "fmt"

func ExampleKendoUiComboBox_ToHtml() {
  html := KendoUiComboBox{
    Html: HtmlInputText{
      Global: HtmlGlobalAttributes{
        Id: "combobox",
      },
    },
    Animation: KendoAnimation{
      Open: KendoOpen{
        Effects: EFFECT_EXPAND_OUT,
        Duration: 300,
      },
      Close: KendoClose{
        Effects: EFFECT_EXPAND_IN,
        Duration: 300,
      },
    },
    AutoBind: FALSE,
    AutoWidth: TRUE,
    CascadeFrom: "parent",
    CascadeFromField: "id",
    ClearButton: FALSE,
    DataSource: []map[string]interface{}{
      { "name": "Child1", "id": 1, "parentId": 1 },
      { "name": "Child2", "id": 2, "parentId": 2 },
      { "name": "Child3", "id": 3, "parentId": 1 },
      { "name": "Child4", "id": 4, "parentId": 2 },
    },
    DataTextField: "name",
    DataValueField: "id",
    Delay: 500,
    Enable: TRUE,
    EnforceMinLength: TRUE,
    Filter: FILTER_CONTAINS,
    Height: 500,
    HighlightFirst: TRUE,
    IgnoreCase: TRUE,
    Index: 0,
    MinLength: 1,
    NoDataTemplate: "No Data!",
    Placeholder: "Select...",
    Suggest: TRUE,
    SyncValueAndText: FALSE,
    Text: "Chai",
    Value: "Child1",
    ValuePrimitive: TRUE,
  }

  fmt.Printf( "%s", html.ToJavaScript() )

  // Output:
  //
}