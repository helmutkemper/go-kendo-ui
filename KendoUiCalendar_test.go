package telerik

import (
  "fmt"
  "time"
)

func ExampleKendoUiCalendar_String() {
  html := HtmlElementDiv{
    Global: HtmlGlobalAttributes{
      Id: "calendar",
    },
  }
  javaScript := KendoUiCalendar{
    HtmlId: "calendar",
    Value: time.Now(),
  }

  fmt.Printf( "%v\n", html.String() )
  fmt.Printf("<script>%v</script>\n", javaScript.String() )

  // Output:
  //
}

func ExampleKendoUiCalendar_ToTelerikTemplate() {
  javaScript := KendoUiCalendar{
    HtmlId: "calendar",
    Culture: "en-US",
    Dates: []time.Time{
      time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
      time.Date(2019, 2, 28, 23, 59, 0, 0, time.UTC),
      time.Date(2019, 2, 28, 23, 59, 59, 0, time.UTC),
    },
    Depth: "year",
    DisableDates: StringArr{ "we", "th" },
    Footer: `kendo.template($('#footer-template').html())`,
    Format: `yyyy/MM/dd`,
    Max: time.Date(2019, 2, 28, 23, 59, 59, 0, time.UTC),
    Messages: &KendoCalendarMessages{
      WeekColumnHeader: "W",
    },
    Min: time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
    Month: &KendoMonth{
      Content: `$('#cell-template').html()`,
      WeekNumber: `$('#week-template').html()`,
      Empty: `-`,
    },
    Selectable: "multiple",
    SelectDates: []time.Time{
      time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
      time.Date(2019, 2, 28, 0, 0, 0, 0, time.UTC),
    },
    WeekNumber: TRUE,
    Start: "year",
    Value: time.Now(),
  }
  fmt.Printf( "%s", javaScript.ToJavaScript() )

  // Output:
  // $("#calendar")kendoCalendar({
  // culture: "en-US",
  // dates: [
  // new Date(2018, 2, 28),
  // new Date(2019, 2, 28, 23, 59),
  // new Date(2019, 2, 28, 23, 59, 59)
  // ],
  // depth: "year",
  // disableDates: [
  // "we",
  // "th"
  // ],
  // footer: "kendo.template($('#footer-template').html())",
  // format: "yyyy/MM/dd",
  // max: new Date(2019, 2, 28, 23, 59, 59),
  // messages: {
  // weekColumnHeader: "W",
  // },
  // min: new Date(2018, 2, 28),
  // month: {
  // content: "$('#cell-template').html()",
  // weekNumber: "$('#week-template').html()",
  // empty: "-",
  // },
  // selectable: "multiple",
  // selectDates: [
  // new Date(2018, 2, 28),
  // new Date(2019, 2, 28)
  // ],
  // weekNumber: true,
  // start: "year",
  // value: new Date(2018, 3, 1, 12, 8, 44),
  // });
}