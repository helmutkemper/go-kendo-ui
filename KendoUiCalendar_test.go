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
