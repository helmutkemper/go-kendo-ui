package telerik

import (
  "fmt"
  "html/template"
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoAggregate struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/aggregate#aggregate

  The aggregates which are calculated when the data source populates with data.

  The supported aggregates are:

  <b>"average"</b> - Only for Number.
  <b>"count"</b> - String, Number and Date.
  <b>"max"</b> - Number and Date.
  <b>"min"</b> - Number and Date.
  <b>"sum"</b> - Only for Number.
  The data source calculates aggregates client-side unless the serverAggregates option is set to <b>true</b>.

  Example - specify aggregates
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

  Example - specify aggregate function
  <script>
  var dataSource = new kendo.data.DataSource({
    data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
    ],
    aggregate: [
      { field: "age", aggregate: "sum" }
    ]
  });
  dataSource.fetch(function(){
    var results = dataSource.aggregates().age;
    console.log(results.sum); // displays "63"
  });
  </script>
  */
  Aggregate                                    TypeAggregate            `jsObject:"aggregate"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/aggregate#aggregate.field

  The data item field which will be used to calculate the aggregates.

  Example - specify aggregate field
  <script>
  var dataSource = new kendo.data.DataSource({
    data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
    ],
    aggregate: [
      { field: "age", aggregate: "sum" }
    ]
  });
  dataSource.fetch(function(){
    var results = dataSource.aggregates().age;
    console.log(results.sum); // displays "63"
  });
  </script>
  */
  Field                                        String                   `jsObject:"field"`

  *ToJavaScriptConverter
}
func(el *KendoAggregate) IsSet() bool {
  return el != nil
}
func(el *KendoAggregate) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Actions", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }

  return buffer.String()
}
func(el *KendoAggregate) ToJavaScript() ([]byte) {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element, "")
  if err != nil {
    log.Criticalf( "KendoAggregate.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}