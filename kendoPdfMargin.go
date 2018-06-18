package telerik

type KendoPdfMargin struct {
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.margin#pdf.margin.top

  The bottom margin. Numbers are considered as "pt" units. (default: 0)
  */
  Top interface{} `jsObject:"top" jsType:"int,string"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.margin#pdf.margin.left

  The left margin. Numbers are considered as "pt" units. (default: 0)
  */
  Left interface{} `jsObject:"left" jsType:"int,string"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.margin#pdf.margin.right

  The right margin. Numbers are considered as "pt" units. (default: 0)
  */
  Right interface{} `jsObject:"right" jsType:"int,string"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.margin#pdf.margin.bottom

  The top margin. Numbers are considered as "pt" units. (default: 0)
  */
  Bottom interface{} `jsObject:"bottom" jsType:"int,string"`
}
