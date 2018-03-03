package telerik

import "fmt"

func ExampleKendoDataSource_ToJavaScript_OfflineStorage_1() {
  javaScript := KendoDataSource{
    VarName: "dataSource",
    Aggregate: &[]KendoAggregates{
      {
        Aggregate: AGGREGATE_SUM,
        Field: "age",
      },
      {
        Aggregate: AGGREGATE_MIN,
        Field: "age",
      },
      {
        Aggregate: AGGREGATE_MAX,
        Field: "age",
      },
    },
    AutoSync: TRUE,
    Batch: TRUE,
    Data: []map[string]interface{}{
      {
        "name": "Jane Doe",
        "age": 30,
      },
      {
        "name": "Jane Doe",
        "age": 33,
      },
    },
    InPlaceSort: FALSE,
    OfflineStorage: "products-offline",
  }

  fmt.Printf( "%s", javaScript.ToJavaScript() )

  // Output:
  //
}

func ExampleKendoDataSource_ToJavaScript_OfflineStorage_2() {
  javaScript := KendoDataSource{
    VarName: "dataSource",
    Aggregate: &[]KendoAggregates{
      {
        Aggregate: AGGREGATE_SUM,
        Field: "age",
      },
      {
        Aggregate: AGGREGATE_MIN,
        Field: "age",
      },
      {
        Aggregate: AGGREGATE_MAX,
        Field: "age",
      },
    },
    AutoSync: TRUE,
    Batch: TRUE,
    Data: []map[string]interface{}{
      {
        "name": "Jane Doe",
        "age": 30,
      },
      {
        "name": "Jane Doe",
        "age": 33,
      },
    },
    InPlaceSort: FALSE,
    OfflineStorage: &OfflineStorage{
      GetItem: "products-key",
      SetItem: "products-key",
    },
  }

  fmt.Printf( "%s", javaScript.ToJavaScript() )

  // Output:
  //
}

func ExampleKendoDataSource_ToJavaScript_OfflineStorage_3() {
  javaScript := KendoDataSource{
    VarName: "dataSource",
    Aggregate: &[]KendoAggregates{
      {
        Aggregate: AGGREGATE_SUM,
        Field: "age",
      },
      {
        Aggregate: AGGREGATE_MIN,
        Field: "age",
      },
      {
        Aggregate: AGGREGATE_MAX,
        Field: "age",
      },
    },
    AutoSync: TRUE,
    Batch: TRUE,
    Data: []map[string]interface{}{
      {
        "name": "Jane Doe",
        "age": 30,
      },
      {
        "name": "Jane Doe",
        "age": 33,
      },
    },
    InPlaceSort: FALSE,
    OfflineStorage: JavaScript{
      Code: `{ getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } }`,
    },
  }

  fmt.Printf( "%s", javaScript.ToJavaScript() )

  // Output:
  //
}
