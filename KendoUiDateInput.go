package telerik

import "time"

type KendoUiDateInput struct{
  HtmlId                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-format

  Specifies the format, which is used to format the value of the DateInput displayed in the input. The format also will be used to parse the input.

  Example - specify a custom date format
   <input id="dateinput" />
   <script>
   $("#dateinput").kendoDateInput({
       format: "yyyy/MM/dd"
   });
   </script>
  */

  Format                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-max

  Specifies the maximum date which can be entered in the input.

  Example - specify the maximum date
   <input id="dateinput" />
   <script>
   $("#dateinput").kendoDateInput({
       max: new Date(2013, 0, 1) // sets max date to Jan 1st, 2013
   });
   </script>
  */

  Max                                     time.Time

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-min

  Specifies the minimum date that which be entered in the input.

  Example - specify the minimum date
   <input id="dateinput" />
   <script>
   $("#dateinput").kendoDateInput({
       min: new Date(2011, 0, 1) // sets min date to Jan 1st, 2011
   });
   </script>
  */

  Min                                     time.Time

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-value

  Specifies the selected date.

  Example
   <input id="dateinput" />
   <script>
   $("#dateinput").kendoDateInput({
       value: new Date(2011, 0, 1)
   });
   </script>
  */

  Value                                   time.Time

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-messages

  The messages that DateInput uses. Use it to customize or localize the placeholders of each date/time part.

  Example - customize column menu messages
   <input id="dateinput" />
   <script>
   $("#dateinput").kendoDateInput({
       messages:{
           "year": "year",
           "month": "month",
           "day": "day",
           "weekday": "day of the week",
           "hour": "hours",
           "minute": "minutes",
           "second": "seconds",
           "dayperiod": "AM/PM"
       }
   });
   </script>
  */

  Messages                                *KendoMessages
}
func(el *KendoUiDateInput) IsSet() bool {
  return el != nil
}

