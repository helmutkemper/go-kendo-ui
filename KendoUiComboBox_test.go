package telerik

import "fmt"

func ExampleKendoUiComboBox_ToHtml() {
  html := KendoUiComboBox{
    Html: HtmlInputText{
      Global: HtmlGlobalAttributes{
        Id: "combobox",
      },
    },
    Animation: &KendoAnimation{
      Open: &KendoOpen{
        Duration: 300,
        Effects: EFFECT_EXPAND_IN,
      },
      Close: &KendoClose{
        Duration: 300,
        Effects:EFFECT_EXPAND_OUT,
      },
    },
    AutoBind: FALSE,
    AutoWidth: TRUE,
    CascadeFrom: "parent",
    CascadeFromField: "dataTextField",
    ClearButton: FALSE,
    DataSource: &[]map[string]interface{}{
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
    Placeholder: "Please, select one",
    Suggest: TRUE,
  }

  fmt.Printf("%s", html.ToJavaScript())

  // Output:
  // $("#combobox").KendoUiComboBox({animation: { close: { effects: "expand:out",duration: 300,},open: { effects: "expand:in",duration: 300,},},autoBind: false,autoWidth: true,cascadeFrom: "parent",cascadeFromField: "dataTextField",clearButton: false,dataSource: [{"name": "Child1","id": 1,"parentId": 1,},{"name": "Child2","id": 2,"parentId": 2,},{"name": "Child3","id": 3,"parentId": 1,},{"name": "Child4","id": 4,"parentId": 2,},],dataTextField: "name",dataValueField: "id",delay: 500,enable: true,enforceMinLength: true,filter: "contains",height: 500,highlightFirst: true,ignoreCase: true,minLength: 1,placeholder: "Please, select one",suggest: true,});
}
