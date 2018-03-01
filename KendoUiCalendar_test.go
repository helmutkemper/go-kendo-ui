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
    Value: time.Now(),
    Dates: []time.Time{
      time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
      time.Date(2019, 2, 28, 23, 59, 0, 0, time.UTC),
    },
  }
  javaScript.ToJavaScript()

  // Output:
  //
}