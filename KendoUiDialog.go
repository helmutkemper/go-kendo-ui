package telerik

type KendoUiDialog struct{
  HtmlId                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions

  A collection of objects containing text, action and primary attributes used to specify the dialog buttons.

  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
         actions: [{
             text: "OK",
             action: function(e){
                 // e.sender is a reference to the dialog widget object
                 // OK action was clicked
                 // Returning false will prevent the closing of the dialog
                 return false;
             },
             primary: true
         },{
             text: "Cancel"
         }]
       });
   </script>
  */

  Actions                                 *KendoActions

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-animation

  A collection of {Animation} objects, used to change default animations. A value of <b><u>false</u></b> will disable all animations in the widget.
  <b><u>animation:true</u></b> is not a valid configuration.

  Example - disable animation
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     animation: false
   });
   </script>
  */

  Animation                               *KendoAnimation

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-buttonLayout

  Specifies the possible layout of the action buttons in the <strong>Dialog</strong>.
  Possible values are:

  Example
   <div id="dialog"></div>
   <script>
       $("#dialog").kendoDialog({
           buttonLayout: "normal"
       });
   </script>
  */

  ButtonLayout                            string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-closable

  Specifies whether a close button should be rendered at the top corner of the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     closable: true
   });
   </script>
  */

  Closable                                Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-content

  Specifies the content of a <strong>Dialog</strong>.

  Example - fetch content from the server
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     content: "<em>Dialog content</em>"
   });
   </script>
  */

  Content                                 string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-height

  Specifies height of the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     height: 400
   });
   </script>
  */

  Height                                  int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-maxHeight

  The maximum height (in pixels) that may be achieved by resizing the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     maxHeight: 300
   });
   </script>
  */

  MaxHeight                               int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-maxWidth

  The maximum width (in pixels) that may be achieved by resizing the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     maxWidth: 300
   });
   </script>
  */

  MaxWidth                                int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages

  Defines the text of the labels that are shown within the dialog. Used primarily for localization.

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

  Messages                                *KendoMessages

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-minHeight

  The minimum height (in pixels) that may be achieved by resizing the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     minHeight: 100
   });
   </script>
  */

  MinHeight                               int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-minWidth

  The minimum width (in pixels) that may be achieved by resizing the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     minWidth: 100
   });
   </script>
  */

  MinWidth                                int

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-modal

  Specifies whether the dialog should show a modal overlay over the page.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     modal: true
   });
   </script>
  */

  Modal                                   Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-title

  The text in the dialog title bar. If <b><u>false</u></b>, the dialog will be displayed without a title bar.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     title: "Customer details"
   });
   </script>
  */

  Title                                   string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-visible

  Specifies whether the dialog will be initially visible.

  Example - show a dialog after one second delay
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     visible: false
   });
   setTimeout(function() {
     $("#dialog").data("kendoDialog").open();
   }, 1000);
   </script>
  */

  Visible                                 Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-width

  Specifies width of the dialog.

  Example
   <div id="dialog"></div>
   <script>
   $("#dialog").kendoDialog({
     width: 400
   });
   </script>
  */

  Width                                   int
}
func(el *KendoUiDialog) IsSet() bool {
  return el != nil
}
