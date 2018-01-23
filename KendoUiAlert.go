package telerik

type KendoUiAlert struct{
  HtmlId                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/alert#configuration-messages

  Defines the text of the labels that are shown within the alert dialog. Used primarily for localization.

  Example
   <div id="alert"></div>
   <script>
   $("#alert").kendoAlert({
     messages:{
       okText: "OK"
     }
   }).data("kendoAlert").open();
   </script>
  */

  Messages                                interface{}
}
func(el *KendoUiAlert) IsSet() bool {
  return el != nil
}

