package telerik

import (
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoGridFilterableOperators struct {
  // @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.enums
  //
  // The texts of the filter operators displayed for columns which have their values option set.
  //
  // Important:
  //
  // Omitting an operator will exclude it from the DropDownList with the available operators.
  //
  //    Example - set enum operators
  //
  //    <div id="grid"></div>
  //    <script>
  //      $("#grid").kendoGrid({
  //        columns: [
  //          { field: "productName" },
  //          { field: "category", values: [
  //              {text: "Beverages", value: 1 },
  //              {text: "Food", value: 2 }
  //            ]
  //          }
  //        ],
  //        dataSource: [
  //          { productName: "Tea", category: 1 },
  //          { productName: "Ham", category: 2 }
  //        ],
  //        filterable: {
  //          operators: {
  //            enums: {
  //              eq: "Equal to",
  //              neq: "Not equal to"
  //            }
  //          }
  //        }
  //      });
  //    </script>
  Enums KendoGridFilterableEnums `jsObject:"enums"`

  *ToJavaScriptConverter
}
func(el *KendoGridFilterableOperators) ToJavaScript() []byte {
  element := reflect.ValueOf(el).Elem()
  ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoGridFilterableOperators.Error: %v", err.Error() )
    return []byte{}
  }

  return ret
}