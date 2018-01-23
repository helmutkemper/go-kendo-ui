package telerik

type KendoAnimation struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close  
  Example - configure the close animation
   <input id="dropdownlist" />
   <script>
   $("#dropdownlist").kendoDropDownList({
     dataSource: ["One", "Two"],
     animation: {
      close: {
        effects: "zoom:out",
        duration: 300
      }
     }
   });
   </script>
  */
  Close                                   *KendoClose

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open  The animation played when the suggestion popup is opened.
  Example - configure the open animation
   <input id="dropdownlist" />
   <script>
   $("#dropdownlist").kendoDropDownList({
     animation: {
      open: {
        effects: "zoom:in",
        duration: 300
      }
     }
   });
   </script>
  */
  Open                                    *KendoOpen

}
func(el *KendoAnimation) IsSet() bool {
  return el != nil
}

