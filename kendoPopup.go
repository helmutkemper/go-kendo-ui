package telerik

type KendoPopup struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.appendTo  Defines a jQuery selector that will be used to find a container element, where the popup will be appended to.
  Example - append the popup to a specific element
   <div id="container">
       <input id="dropdownlist" />
   </div>
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: [
       { id: 1, name: "Apples" },
       { id: 2, name: "Oranges" }
     ],
     dataTextField: "name",
     dataValueField: "id",
     popup: {
       appendTo: $("#container")
     }
   });
   </script>
  */
  AppendTo                                string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.origin  Specifies how to position the popup element based on anchor point. The value is space separated "y" plus "x" position.
  The available "y" positions are:
  The available "x" positions are:
  Example - append the popup to a specific element
   <div id="container">
       <input id="dropdownlist" />
   </div>
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: [
       { id: 1, name: "Apples" },
       { id: 2, name: "Oranges" }
     ],
     dataTextField: "name",
     dataValueField: "id",
     popup: {
       origin: "top left"
     }
   });
   </script>
  */
  Origin                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.position  Specifies which point of the popup element to attach to the anchor's origin point. The value is space separated "y" plus "x" position.
  The available "y" positions are:
  The available "x" positions are:
  Example - append the popup to a specific element
   <div id="container">
       <input id="dropdownlist" />
   </div>
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: [
       { id: 1, name: "Apples" },
       { id: 2, name: "Oranges" }
     ],
     dataTextField: "name",
     dataValueField: "id",
     popup: {
       origin: "top left"
     }
   });
   </script>
  */
  Position                                string

}
func(el *KendoPopup) IsSet() bool {
  return el != nil
}

