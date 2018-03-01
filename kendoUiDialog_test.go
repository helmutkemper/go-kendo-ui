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
    Modal: TRUE,
    Visible: TRUE,
    Title: "Configuration error",
    Content: "Please, set a valid environment var.",
    Width: 400,
    Actions: &[]KendoActions{
      {
        Action:  "function(){ setTimeout( function(){ containerConfigurationEnvVarNameRef.focus(); }, 1000); }",
        Primary: TRUE,
        Text:    "OK",
      },
    },
  }

  fmt.Printf( "%v\n", html.String() )
  fmt.Printf("<script>%v</script>\n", javaScript.String() )

  // Output:
  // <div  id="dialog" ></div>
  // <script>$("#dialog").kendoDialog({ actions: [{ text: "OK",action: function(){ setTimeout( function(){ containerConfigurationEnvVarNameRef.focus(); }, 1000); },primary: true, },],content: "Please, set a valid environment var.",modal: true,title: "Configuration error",visible: true,width: 400, });</script>
}
