package telerik

type KendoMessages struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages.close  The title of the close button.
  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     messages:{
       close: "Close Me!"
     }
   });
   </script>
  */
  Close                                   string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages.promptInput  The title of the prompt input.
  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     messages:{
       promptInput: "Input!"
     }
   });
   </script>
  */
  PromptInput                             string

}
func(el *KendoMessages) IsSet() bool {
  return el != nil
}

