package telerik

import "fmt"

func ExampleKendoUiGrid_ToHtml_AllowCopy() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Selectable: KENDO_GRID_SELECTABLE_MULTIPLE_CELL,
    AllowCopy: TRUE,
    Columns: []KendoGridColumns{
      { Field: "productName" },
      { Field: "category" },
    },
    DataSource: []map[string]interface{}{
      { "productName": "Tea", "category": "Beverages" },
      { "productName": "Coffee", "category": "Beverages" },
      { "productName": "Ham", "category": "Food" },
      { "productName": "Bread", "category": "Food" },
    },
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({allowCopy: true,columns: [{field: "productName",},{field: "category",},],dataSource: [{"productName": "Tea","category": "Beverages",},{"productName": "Coffee","category": "Beverages",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],selectable: "multiple, cell",});
}