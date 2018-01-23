package telerik

type KendoActions struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions.text  The text to be shown in the action's button.
  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
         }]
       });
   </script>
  */
  Text                                    string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions.action  The callback function to be called after pressing the action button.
  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
             action: function(e){
                 // e.sender is a reference to the dialog widget object
                 alert("OK action was clicked");
                 // Returning false will prevent the closing of the dialog
                 return true;
             },
         }]
       });
   </script>
  */
  Action                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions.primary  A boolean property indicating whether the action button will be decorated as primary button or not.
  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
             primary: true
         }]
       });
   </script>
  */
  Primary                                 Boolean

}
func(el *KendoActions) IsSet() bool {
  return el != nil
}

