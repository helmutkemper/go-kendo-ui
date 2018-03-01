package telerik

import (
  "fmt"
  "html/template"
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoMonth struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-month.content  Template to be used for rendering the cells in the calendar "month" view, which are in range.
  Example - specify cell template as a string
   <input id="datetimepicker" />
   <script id="cell-template" type="text/x-kendo-template">
       <div class="#= data.value < 10 ? 'exhibition' : 'party' #"></div>
       #= data.value #
   </script>
   <script>
   $("#datetimepicker").kendoDateTimePicker({
       month: {
          content: $("#cell-template").html()
       }
   });
   </script>
  */
  Content                                 String                  `jsObject:"content"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-month.weekNumber  The template to be used for rendering the cells in "week" column. By default, the widget renders the calculated week of the year. The properties available in the data object are:
  weekNumber - calculated week number.
  These properties can be used in the template to make additional calculations. 
  Example - specify week number template as a string
   <style>
     .italic{
       font-style: italic;
     }
   </style>
   <body>
   
   <input id="datetimepicker1" />
   <script id="week-template" type="text/x-kendo-template">
      <a class="italic">#= data.weekNumber #</a>
   </script>
   <script>
     $("#datetimepicker1").kendoDateTimePicker({
       weekNumber: true,
       month: {
         weekNumber: $("#week-template").html()
       }
     });
   </script>
  */
  WeekNumber                              String                  `jsObject:"weekNumber"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-month.empty  The template used for rendering cells in the calendar "month" view, which are outside the min/max range.
  Example - specify an empty cell template as a string
   <input id="datetimepicker1" />
   <script>
   $("#datetimepicker1").kendoDateTimePicker({
       month: {
          empty: '-'
       }
   });
   </script>
  */
  Empty                                   String                  `jsObject:"empty"`

  *ToJavaScriptConverter

}
func(el *KendoMonth) IsSet() bool {
  return el != nil
}
func(el *KendoMonth) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(s.(string))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "Month", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}
func(el *KendoMonth) ToJavaScript() ([]byte) {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoMonth.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}