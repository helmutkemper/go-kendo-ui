package telerik

type KendoDataSource struct {
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

  Aggregate                               *KendoAggregate
}