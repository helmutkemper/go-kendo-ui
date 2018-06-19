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
  //$("#grid").kendoGrid({allowCopy: true,columns: [{field: "productName",},{field: "category",},],dataSource: [{"category": "Beverages","productName": "Tea",},{"productName": "Coffee","category": "Beverages",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],selectable: "multiple, cell",});
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
  //$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",},{field: "age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],sortable: true,});
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
  //$("#grid").kendoGrid({columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],});
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
      { Command: []TypeKendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
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
  //$("#grid").kendoGrid({columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"age": 30,"name": "Jane Doe",},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",});
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
      { Command: []TypeKendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
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
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"age": 30,"name": "Jane Doe",},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: ["save","excel",],});
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
      { Command: []TypeKendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
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
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"age": 33,"name": "John Doe",},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_Columns_Command_Destroy() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
      { Command: []TypeKendoGridColumnsCommand{ COLUMNS_COMMAND_EDIT } },
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
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_Columns_Command_2() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
      { Command: []KendoGridColumnsCommand{
        {
          Name: COLUMNS_COMMAND_CUSTOM,
          Text: "details",
          Click: JavaScript{ Code: `function(e){ /* command button click handler */ }` },
          ClassName: "btn-destroy",
          IconClass: "k-icon k-i-copy",
        },
        {
          Name: COLUMNS_COMMAND_DESTROY,
        },
      } },
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
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",click: function(e){ /* command button click handler */ },iconClass: "k-icon k-i-copy",name: "custom",text: "details",},{name: "destroy",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_Columns_Command_3() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },
    Columns: []KendoGridColumns{
      { Field: "name", Title: "Name" },
      { Field: "age", Title: "Age" },
      { Command: []KendoGridColumnsCommand{
        {
          Name: COLUMNS_COMMAND_DESTROY,
          Text: "remove",
          ClassName: "btn-destroy",
          IconClass: KendoGridColumnsIconClass{
            Edit: "k-icon k-i-edit",
            Update: "k-icon k-i-copy",
            Cancel: "k-icon k-i-arrow-60-up",
          },
        },
        {
          Name: COLUMNS_COMMAND_EDIT,
        },
      } },
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
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",iconClass: {edit: "k-icon k-i-edit",update: "k-icon k-i-copy",cancel: "k-icon k-i-arrow-60-up",},name: "destroy",text: "remove",},{name: "edit",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_Filterable_Cell() {
  el := KendoUiGrid{
    Html: HtmlElementDiv{
      Global: HtmlGlobalAttributes{
        Id: "grid",
      },
    },

    Columns: []KendoGridColumns{
      {
        Field: "name",
        Title: "Name",
        Filterable: KendoGridColumnsFilterable{
          Cell: KendoGridColumnsCell{
            DataSource: KendoDataSource{
              Data: []map[string]interface{}{
                { "someField": "Jane" },
                { "someField": "Jake" },
                { "someField": "John" },
              },
            },
            DataTextField: "someField",
          },
        },
      },
      {
        Field: "age",
        Title: "Age",
      },
      {
        Command: []KendoGridColumnsCommand{
        {
          Name: COLUMNS_COMMAND_DESTROY,
          Text: "remove",
          ClassName: "btn-destroy",
        },
        {
          Name: COLUMNS_COMMAND_EDIT,
        },
      } },
    },
    Filterable: KendoGridFilterable{
      Mode: KENDO_GRID_FILTERABLE_MODE_ROW,
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
  //$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",filterable: {cell: {dataSource:  new kendo.data.DataSource({data: [{"someField": "Jane",},{"someField": "Jake",},{"someField": "John",},],}),dataTextField: "someField",},},title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},{name: "edit",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},filterable: {mode: "row",},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}