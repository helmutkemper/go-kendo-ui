package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
  "bytes"
)

type KendoDataSource struct {
  VarName                                 String              `jsObject:"htmlId"`

  /*
  @sse https://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate

  The aggregates which are calculated when the data source populates with data.

  The supported aggregates are:

  > "average" - Only for Number.
  > "count" - String, Number and Date.
  > "max" - Number and Date.
  > "min" - Number and Date.
  > "sum" - Only for Number.

  The data source calculates aggregates client-side unless the serverAggregates option is set to true.

  EXAMPLE - SPECIFY AGGREGATES
  <script>
  var dataSource = new kendo.data.DataSource({
    data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
    ],
    aggregate: [
      { field: "age", aggregate: "sum" },
      { field: "age", aggregate: "min" },
      { field: "age", aggregate: "max" }
    ]
  });
  dataSource.fetch(function(){
    var results = dataSource.aggregates().age;
    console.log(results.sum, results.min, results.max); // displays "63 30 33"
  });
  </script>
  */

  Aggregate                               *[]KendoAggregate                   `jsObject:"aggregate"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/autosync#autoSync

  If set to true the data source would automatically save any changed data items by calling the sync method. By default, changes are not automatically saved. (default: false)

  Example - enable auto sync
  <script>
  var dataSource = new kendo.data.DataSource({
    autoSync: true,
    transport: {
      read:  {
        url: "https://demos.telerik.com/kendo-ui/service/products",
        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
      },
      update: {
        url: "https://demos.telerik.com/kendo-ui/service/products/update",
        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
      }
    },
    schema: {
      model: { id: "ProductID" }
    }
  });
  dataSource.fetch(function() {
    var product = dataSource.at(0);
    product.set("UnitPrice", 20); // auto-syncs and makes request to https://demos.telerik.com/kendo-ui/service/products/update
  });
  </script>
  */
  AutoSync                                Boolean                           `jsObject:"autoSync"`

  /*
  @sse https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/batch#batch

  If set to <b>true</b>, the data source will batch CRUD operation requests. For example, updating two data items would cause one HTTP request instead of two. By default, the data source makes a HTTP request for every CRUD operation. (default: false)
  The changed data items are sent as <b>models</b> by default. This can be changed via the parameterMap option.

  Example - enable the batch mode
  <script>
  var dataSource = new kendo.data.DataSource({
    batch: true,
    transport: {
      read:  {
        url: "https://demos.telerik.com/kendo-ui/service/products",
        dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
      },
      update: {
        url: "https://demos.telerik.com/kendo-ui/service/products/update",
        dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
      }
    },
    schema: {
      model: { id: "ProductID" }
    }
  });
  dataSource.fetch(function() {
    var product = dataSource.at(0);
    product.set("UnitPrice", 20);
    var anotherProduct = dataSource.at(1);
    anotherProduct.set("UnitPrice", 20);
    dataSource.sync(); // causes only one request to "https://demos.telerik.com/kendo-ui/service/products/update"
  });
  </script>
  */
  Batch                                   Boolean                           `jsObject:"batch"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/data#data

  The array of data items which the data source contains. The data source will wrap those items as kendo.data.ObservableObject or kendo.data.Model (if schema.model is set).

  Example - set the data items of a data source
  <script>
  var dataSource = new kendo.data.DataSource({
    data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
    ]
  });
  dataSource.fetch(function(){
    var janeDoe = dataSource.at(0);
    console.log(janeDoe.name); // displays "Jane Doe"
  });
  </script>
  */
  //fixme: data pode ser um xml
  Data                                    []MapStringArr                    `jsObject:"data"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/inplacesort#inPlaceSort

  If set to <b>true</b> the original <b>Array</b> used as data will be sorted when sorting operation is performed. This setting supported only with local data, bound to a JavaScript array via the data option. (default: false)
  */
  InPlaceSort                             Boolean                           `jsObject:"inPlaceSort"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/offlinestorage#offlineStorage

  Example - set offline storage key
  <script>
  var dataSource = new kendo.data.DataSource({
      offlineStorage: "products-offline",
      transport: {
          read: {
              url: "https://demos.telerik.com/kendo-ui/service/products",
              type: "jsonp"
          }
      }
  });
  </script>
  */
  OfflineStorage                          interface{}                       `jsObject:"offlineStorage"`

  *ToJavaScriptConverter
}
func(el *KendoDataSource) ToJavaScript() []byte {
  var ret bytes.Buffer
  if el.VarName == "" {
    log.Critical("KendoDataSource not have a VarName for mount JavaScript code.")
    return []byte{}
  }

  element := reflect.ValueOf(el).Elem()
  data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoDataSource.Error: %v", err.Error() )
    return []byte{}
  }

  ret.Write( []byte(`var ` + el.VarName + ` = new kendo.data.DataSource({`) )
  ret.Write( data )
  ret.Write( []byte(`});`) )


  return ret.Bytes()
}