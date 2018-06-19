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
  //$("#grid").kendoGrid({allowCopy: true,columns: [{field: "productName",},{field: "category",},],dataSource: [{"productName": "Tea","category": "Beverages",},{"productName": "Coffee","category": "Beverages",},{"category": "Food","productName": "Ham",},{"productName": "Bread","category": "Food",},],selectable: "multiple, cell",});
}

func ExampleKendoUiGrid_ToHtml_ColumnMenu() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name" },
      { Field: "age" },
    },
    ColumnMenu: TRUE,
    DataSource: []map[string]interface{}{
      { "name": "Jane Doe", "age": 30 },
      { "name": "John Doe", "age": 33 },
    },
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",},{field: "age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],});
}

func ExampleKendoUiGrid_ToHtml_ColumnMenu_sortable() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name" },
      { Field: "age" },
    },
    ColumnMenu: KendoGridColumnMenu{
      Columns: FALSE,
    },
    Sortable: TRUE,
    DataSource: []map[string]interface{}{
      { "name": "Jane Doe", "age": 30 },
      { "name": "John Doe", "age": 33 },
    },
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",},{field: "age",},],dataSource: [{"age": 30,"name": "Jane Doe",},{"name": "John Doe","age": 33,},],sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_ColumnMenu_messages() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name" },
      { Field: "age" },
    },
    ColumnMenu: KendoGridColumnMenu{
      Messages: KendoGridColumnMenuMessages{
        Columns: "Choose columns",
        Filter: "Apply filter",
        SortAscending: "Sort (asc)",
        SortDescending: "Sort (desc)",
      },
    },
    Sortable: TRUE,
    Filterable: TRUE,
    DataSource: []map[string]interface{}{
      { "name": "Jane Doe", "age": 30 },
      { "name": "John Doe", "age": 33 },
    },
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columnMenu: {messages: {columns: "Choose columns",filter: "Apply filter",sortAscending: "Sort (asc)",sortDescending: "Sort (desc)",},},columns: [{field: "name",},{field: "age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],filterable: true,sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_ColumnMenu_Columns() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
    },
    DataSource: []map[string]interface{}{
      { "name": "Jane Doe", "age": 30 },
      { "name": "John Doe", "age": 33 },
    },
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"age": 33,"name": "John Doe",},],});
}

func ExampleKendoUiGrid_ToHtml_ColumnMenu_EditorMode() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
      { Command: []kendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
    },

    DataSource: KendoDataSource{
      Data: []map[string]interface{}{
        { "name": "Jane Doe", "age": 30 },
        { "name": "John Doe", "age": 33 },
      },
      Schema: KendoSchema{
        Model: KendoDataModel{
          Id: "id",
          Fields: map[string]KendoField{
            "age": {
              Type: JAVASCRIPT_NUMBER,
            },
          },
        },
      },
    },
    Editable: KENDO_GRID_EDITOR_MODE_POPUP,
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{Command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"age": 30,"name": "Jane Doe",},{"age": 33,"name": "John Doe",},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",});
}

func ExampleKendoUiGrid_ToHtml_Excel_AllPages() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
      { Command: []kendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
    },
    ColumnMenu: TRUE,
    Toolbar: []KendoGridToolBarName{ GRID_TOOLBAR_NAME_SAVE, GRID_TOOLBAR_NAME_EXCEL },
    Excel: KendoGridExcel{
      AllPages: TRUE,
    },
    DataSource: KendoDataSource{
      Data: []map[string]interface{}{
        { "name": "Jane Doe", "age": 30 },
        { "name": "John Doe", "age": 33 },
      },
      Schema: KendoSchema{
        Model: KendoDataModel{
          Id: "id",
          Fields: map[string]KendoField{
            "age": {
              Type: JAVASCRIPT_NUMBER,
            },
          },
        },
      },
    },
    Editable: KENDO_GRID_EDITOR_MODE_POPUP,
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{Command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: ["save","excel",],});
}

func ExampleKendoUiGrid_ToHtml_Toolbar_Excel() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
      { Command: []kendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
    },
    ColumnMenu: TRUE,
    Toolbar: []KendoGridToolbar{
      {
        Name: GRID_TOOLBAR_NAME_SAVE,
        IconClass: "k-icon k-i-copy",
      },
      {
        Name: GRID_TOOLBAR_NAME_EXCEL,
      },
    },
    Excel: KendoGridExcel{
      AllPages: TRUE,
    },
    DataSource: KendoDataSource{
      Data: []map[string]interface{}{
        { "name": "Jane Doe", "age": 30 },
        { "name": "John Doe", "age": 33 },
      },
      Schema: KendoSchema{
        Model: KendoDataModel{
          Id: "id",
          Fields: map[string]KendoField{
            "age": {
              Type: JAVASCRIPT_NUMBER,
            },
          },
        },
      },
    },
    Editable: KENDO_GRID_EDITOR_MODE_POPUP,
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s", el.ToJavaScript() )

  //Output:
  //<div id="grid"></div>
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{Command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}