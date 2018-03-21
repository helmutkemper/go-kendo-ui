package telerik

import (
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoUiDialog struct{
  Html                                    HtmlElementDiv                          `jsObject:"-"`

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
  Actions                                 *[]KendoActions                         `jsObject:"actions"`

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
  Animation                               interface{}                             `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

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
  ButtonLayout                            KendoButtonLayout                       `jsObject:"buttonLayout"`

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
  Closable                                Boolean                                 `jsObject:"closable"`

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
  Content                                 string                                  `jsObject:"content"`

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
  Height                                  int                                     `jsObject:"height"`

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
  MaxHeight                               int                                     `jsObject:"maxHeight"`

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
  MaxWidth                                int                                     `jsObject:"maxWidth"`

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
  Messages                                *KendoMessages                          `jsObject:"messages"`

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
  MinHeight                               int                                     `jsObject:"minHeight"`

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
  MinWidth                                int                                     `jsObject:"minWidth"`

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
  Modal                                   Boolean                                 `jsObject:"modal"`

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
  Title                                   string                                  `jsObject:"title"`

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
  Visible                                 Boolean                                 `jsObject:"visible"`

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
  Width                                   int                                     `jsObject:"width"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/close#close

  Triggered when a Dialog is closed (by a user or through the close() method).

  Event Data: e.userTriggered - Boolean
  Indicates whether the close action has been triggered by the user (by clicking the close button or hitting the escape key). When the close method has been called, this field is false.

  Example - subscribe to the "close" event during initialization
  <div id="dialog"></div>
  <script>
  $("#dialog").kendoDialog({
    close: function(e) {
      // close animation has finished playing
    }
  });
  </script>

  Example - subscribe to the "close" event after initialization
  <div id="dialog"></div>
  <script>
    function dialog_close(e) {
      // close animation has finished playing
    }
    $("#dialog").kendoDialog();
    var dialog = $("#dialog").data("kendoDialog");
    dialog.bind("close", dialog_close);
  </script>
  */
  EventClose                              *JavaScript                             `jsObject:"close"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/hide#hide

  Triggered when a Dialog has finished its closing animation.

  Example - subscribe to the "hide" event during initialization
  <div id="dialog"></div>
  <script>
  $("#dialog").kendoDialog({
    hide: function() {
      // close animation is about to finish
    }
  });
  </script>

  Example - subscribe to the "hide" event after initialization
  <div id="dialog"></div>
  <script>
    function dialog_hide() {
      // close animation will start soon
    }
    $("#dialog").kendoDialog();
    var dialog = $("#dialog").data("kendoDialog");
    dialog.bind("hide", dialog_hide);
  </script>
  */
  EventHide                               *JavaScript                             `jsObject:"hide"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/initopen#initOpen

  Triggered when a Dialog is opened for the first time (i.e. the open() method is called).

  Example - subscribe to the "initOpen" event during initialization
  <div id="dialog"></div>
  <script>
  $("#dialog").kendoDialog({
    initOpen: function() {
      // open animation will start soon
    }
  });
  </script>

  Example - subscribe to the "initOpen" event after initialization
  <div id="dialog" style="display: none;"></div>
  <script>
    function dialog_initOpen() {
      // open animation will start soon
    }
    $("#dialog").kendoDialog();
    var dialog = $("#dialog").data("kendoDialog");
    dialog.bind("initOpen", dialog_initOpen);
    dialog.open();
  </script>
  */
  EventInitOpen                           *JavaScript                             `jsObject:"initOpen"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/open#open

  Triggered when a Dialog is opened (i.e. the open() method is called).

  Example - subscribe to the "open" event during initialization
  <div id="dialog"></div>
  <script>
  $("#dialog").kendoDialog({
    open: function() {
      // open animation will start soon
    }
  });
  </script>

  Example - subscribe to the "open" event after initialization
  <div id="dialog"></div>
  <script>
    function dialog_open() {
      // open animation will start soon
    }
    $("#dialog").kendoDialog();
    var dialog = $("#dialog").data("kendoDialog");
    dialog.bind("open", dialog_open);
  </script>
  */
  EventOpen                               *JavaScript                             `jsObject:"open"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/show#show

  Triggered when a Dialog has finished its opening animation.

  Example - subscribe to the "show" event during initialization
  <div id="dialog"></div>
  <script>
  $("#dialog").kendoDialog({
    show: function() {
      // open animation has finished playing
    }
  });
  </script>

  Example - subscribe to the "show" event after initialization
  <div id="dialog"></div>
  <script>
    function dialog_show() {
      // open animation has finished playing
    }
    $("#dialog").kendoDialog();
    var dialog = $("#dialog").data("kendoDialog");
    dialog.bind("show", dialog_show);
  </script>
  */
  EventShow                               *JavaScript                             `jsObject:"show"`

  *ToJavaScriptConverter
}
func(el *KendoUiDialog) ToJavaScript() []byte {
  var ret bytes.Buffer
  if el.Html.Global.Id == "" {
    log.Critical("kendoDialog not have a html id for mount JavaScript code.")
    return []byte{}
  }

  element := reflect.ValueOf(el).Elem()
  data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "kendoDialog.Error: %v", err.Error() )
    return []byte{}
  }

  ret.Write( []byte(`$("#` + el.Html.Global.Id + `").kendoDialog({`) )
  ret.Write( data )
  ret.Write( []byte(`});`) )

  return ret.Bytes()
}
func(el *KendoUiDialog) ToHtml() []byte{
  return el.Html.ToHtml()
}