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

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/page#page

  The page of data which the data source will return when the view method is invoked or request from the remote service.
  The data source will page the data items client-side unless the serverPaging option is set to <b>true</b>.

  Example - set the current page
  <script>
  var dataSource = new kendo.data.DataSource({
    data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
    ],
    // set the second page as the current page
    page: 2,
    pageSize: 2
  });
  dataSource.fetch(function(){
    var view = dataSource.view();
    console.log(view.length); // displays "1"
    console.log(view[0].name); // displays "Ham"
  });
  </script>
  */
  Page                                    Int                               `jsObject:"page"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/pagesize#pageSize

  The number of data items per page. The property has no default value. <b>That is why to use paging, make sure some pageSize value is set</b>.
  The data source will page the data items client-side unless the serverPaging option is set to <b>true</b>.

  Example - set the page size
  <script>
  var dataSource = new kendo.data.DataSource({
    data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
    ],
    page: 1,
    // a page of data contains two data items
    pageSize: 2
  });
  dataSource.fetch(function(){
    var view = dataSource.view();
    console.log(view.length); // displays "2"
    console.log(view[0].name); // displays "Tea"
    console.log(view[1].name); // displays "Coffee"
  });
  </script>
  */
  PageSize                                Int                               `jsObject:"pageSize"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serveraggregates#serverAggregates

  If set to true, the data source will leave the aggregate calculation to the remote service. By default, the data source calculates aggregates client-side.
  <b>Configure schema.aggregates if you set serverAggregates to true</b>.
  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

  Example - enable server aggregates
  <script>
  var dataSource = new kendo.data.DataSource({
    transport: {
      // transport configuration
    },
    serverAggregates: true,
    aggregate: [
      { field: "age", aggregate: "sum" }
    ],
    schema: {
      aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
    }
  });
  </script>
  */
  ServerAggregates                        Boolean                           `jsObject:"serverAggregates"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serverfiltering#serverFiltering

  If set to true, the data source will leave the filtering implementation to the remote service. By default, the data source performs filtering client-side.
  By default, the filter is sent to the server following jQuery's conventions.

  For example, the filter { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] } is sent as:

  filter[logic]: and
  filter[filters][0][field]: name
  filter[filters][0][operator]: startswith
  filter[filters][0][value]: Jane
  Use the parameterMap option to send the filter option in a different format.

  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

  Example - enable server filtering
  <script>
  var dataSource = new kendo.data.DataSource({
    transport: {
      // transport configuration
    },
    serverFiltering: true,
    filter: { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }
  });
  </script>
  */
  ServerFiltering                         Boolean                           `jsObject:"serverFiltering"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/servergrouping#serverGrouping

  If set to true, the data source will leave the grouping implementation to the remote service. By default, the data source performs grouping client-side. (default: false)
  By default, the group is sent to the server following jQuery's conventions.

  For example, the group { field: "category", dir: "desc" } is sent as:

  group[0][field]: category
  group[0][dir]: desc
  Use the parameterMap option to send the group option in a different format.

  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

  Example - enable server grouping
  <script>
  var dataSource = new kendo.data.DataSource({
    transport: {
      // transport configuration
    },
    serverGrouping: true,
    group: { field: "category", dir: "desc" }
  });
  </script>
  */
  ServerGrouping                          Boolean                           `jsObject:"serverGrouping"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serverpaging#serverPaging

  If set to true, the data source will leave the data item paging implementation to the remote service. By default, the data source performs paging client-side. (default: false)
  Configure schema.total if you set serverPaging to true. In addition, pageSize should be set no matter if paging is performed client-side or server-side.

  The following options are sent to the server when server paging is enabled:

  page - the page of data item to return (1 means the first page).
  pageSize - the number of items to return.
  skip - how many data items to skip.
  take - the number of data items to return (the same as pageSize).
  Use the parameterMap option to send the paging options in a different format.

  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

  Example - enable server paging
  <script>
  var dataSource = new kendo.data.DataSource({
    transport: {
      // transport configuration
    },
    serverPaging: true,
    schema: {
      total: "total" // total is returned in the "total" field of the response
    }
  });
  </script>
  */
  ServerPaging                            Boolean                           `jsObject:"serverPaging"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/type#type

  If set, the data source will use a predefined transport and/or schema.

  The supported values are:

  "odata" which supports the OData v.2 protocol - http://www.odata.org/
  "odata-v4" which partially supports odata version 4 - https://github.com/telerik/ui-for-aspnet-mvc-examples/tree/master/grid/odata-v4-web-api-binding
  "signalr"

  Example - enable OData support
  <script>
  var dataSource= new kendo.data.DataSource({
    type: "odata",
    transport: {
      read: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
    },
    pageSize: 20,
    serverPaging: true
  });
  dataSource.fetch(function() {
    console.log(dataSource.view().length); // displays "20"
  });
  </script>
  */

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