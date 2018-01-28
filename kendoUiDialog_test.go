package telerik

import "fmt"

func ExampleKendoUiDialog_String(){
  html := HtmlElementDiv{
    Global: HtmlGlobalAttributes{
      Id: "dialog",
    },
  }
  javaScript := KendoUiDialog{
    HtmlId: "dialog",
    Actions: &[]KendoActions{
      {
        Action:  "function(el){ return false; }",
        Primary: TRUE,
        Text:    "Cancel",
      },
      {
        Primary: FALSE,
        Text:    "OK",
      },
    },
  }

  fmt.Printf( "%v\n", html.String() )
  fmt.Printf("<script>%v</script>\n", javaScript.String() )

  // Output:
  // <div Id="dialog" ></div>
  // <script>$("#dialog").kendoDialog({ actions: [{ Text: "Cancel",Action: function(el){ return false; },Primary: true, },{ Text: "OK",Primary: false, },], });</script>
}
