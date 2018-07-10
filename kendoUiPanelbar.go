package telerik

import (
  "bytes"
  "reflect"
  log "github.com/helmutkemper/seelog"
)

type KendoUiPanelBar struct {
  Html HtmlElementUl `jsObject:"-"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/animation

  A collection of visual animations used when PanelBar items are expand or collapsed through user interactions. Setting
  this option to false will disable all animations.

  Important:

  animation:true is not a valid configuration.
  */
  //    Example
  //
  //    <ul id="panelbar">
  //      <li>Item 1
  //          <ul>
  //              <li>Sub Item 1</li>
  //              <li>Sub Item 2</li>
  //              <li>Sub Item 3</li>
  //          </ul>
  //      </li>
  //      <li>Item 2
  //          <ul>
  //              <li>Sub Item 1</li>
  //              <li>Sub Item 2</li>
  //              <li>Sub Item 3</li>
  //          </ul>
  //      </li>
  //    </ul>
  //    <script>
  //        $("#panelbar").kendoPanelBar({
  //            animation: {
  //                // fade-out closing items over 1000 milliseconds
  //                collapse: {
  //                    duration: 1000,
  //                    effects: "fadeOut"
  //                },
  //               // fade-in and expand opening items over 500 milliseconds
  //               expand: {
  //                   duration: 500,
  //                   effects: "expandVertical fadeIn"
  //               }
  //           }
  //        });
  //    </script>
  Animation interface{} `jsObject:"animation" jsType:"KendoAnimationVertical,Boolean"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/autobind

  If set to false the widget will not bind to the data source during initialization. In this case data binding will
  occur when the change event of the data source is fired. By default the widget will bind to the data source specified
  in the configuration. (default: true)

  Important:

  Setting autoBind to false is useful when multiple widgets are bound to the same data source. Disabling automatic
  binding ensures that the shared data source does not make more than one request to the remote service.
  */
  //    Example - disable automatic binding
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    var dataSource = new kendo.data.HierarchicalDataSource({
  //      data: [ { text: "Jane Doe" }, { text: "John Doe" }]
  //    });
  //    $("#panelbar").kendoPanelBar({
  //      autoBind: false,
  //      dataSource: dataSource
  //    });
  //    dataSource.read(); // "read()" will fire the "change" event of the dataSource and the widget will be bound
  //    </script>
  AutoBind Boolean `jsObject:"autoBind"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/contenturls

  Sets an array with the URLs from which the PanelBar items content to be loaded from. If only specific items should be
  loaded via Ajax, then you should set the URLs to the corresponding positions in the array and set the other elements
  to null.
  */
  //    Example - specify that the second item should be loaded remotely
  //
  //    <ul id="panelbar">
  //        <li>Item 1
  //          <div>Content 1</div>
  //        </li>
  //        <li>
  //            Ajax Item
  //            <div></div>
  //        </li>
  //    </ul>
  //
  //    <script>
  //        $("#panelbar").kendoPanelBar({
  //            contentUrls: [
  //              null,
  //              "https://demos.telerik.com/kendo-ui/content/web/panelbar/ajax/ajaxContent1.html"
  //            ]
  //        });
  //    </script>
  ContentUrls []string `jsObject:"contentUrls"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/dataimageurlfield

  Sets the field of the data item that provides the image URL of the PanelBar nodes. (default: null)
  */
  //    Example - specify custom image URL field
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    var items = [
  //      { text: "Baseball", image: "https://demos.telerik.com/kendo-ui/content/shared/icons/sports/baseball.png" },
  //      { text: "Golf", image: "https://demos.telerik.com/kendo-ui/content/shared/icons/sports/golf.png" }
  //    ];
  //    $("#panelbar").kendoPanelBar({
  //      dataImageUrlField: "image",
  //      dataSource: items
  //    });
  //    </script>
  DataImageUrlField string `jsObject:"dataImageUrlField"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/datasource

  The data source of the widget which is used render nodes. Can be a JavaScript object which represents a valid data
  source configuration, a JavaScript array or an existing kendo.data.HierarchicalDataSource instance.

  If the dataSource option is set to a JavaScript object or array the widget will initialize a new
  kendo.data.HierarchicalDataSource instance using that value as data source configuration.

  If the dataSource option is an existing kendo.data.HierarchicalDataSource instance the widget will use that instance
  and will not initialize a new one.
  */
  //    Example - set dataSource as a JavaScript object
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    $("#panelbar").kendoTreeView({
  //      dataSource: {
  //        data: [
  //          { text: "foo", items: [
  //            { text: "bar" }
  //          ] }
  //        ]
  //      }
  //    });
  //    </script>
  //
  //
  //    Example - set dataSource as a JavaScript array
  //
  //    <ul id="panelbar"></ul>
  //    <script>
  //      $("#panelbar").kendoPanelBar({
  //          dataSource: [
  //              {
  //                  text: "Item 1 (link)",
  //                  cssClass: "myClass",                            // Add custom CSS class to the item, optional, added 2012 Q3 SP1.
  //                  url: "http://www.kendoui.com/"                  // link URL if navigation is needed (optional)
  //              },
  //              {
  //                  text: "<b>Item 2</b>",
  //                  encoded: false,                                 // Allows use of HTML for item text
  //                  content: "text"                                 // content within an item
  //              },
  //              {
  //                  text: "Item 3",
  //                  // content URL to load within an item
  //                  contentUrl: "https://demos.telerik.com/kendo-ui/content/web/panelbar/ajax/ajaxContent1.html"
  //              },
  //              {
  //                  text: "Item 4",
  //                  // item image URL, optional
  //                  imageUrl: "https://demos.telerik.com/kendo-ui/content/shared/icons/sports/baseball.png",
  //                  expanded: true,                                 // item is rendered expanded
  //                  items: [{                                       // Sub item collection.
  //                      text: "Sub Item 1"
  //                  },
  //                  {
  //                      text: "Sub Item 2"
  //                  }]
  //              },
  //              {
  //                  text: "Item 5"
  //              }
  //          ]
  //      });
  //    </script>
  //
  //
  //    Example - set dataSource as an existing kendo.data.HierarchicalDataSource instance
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    var dataSource = new kendo.data.HierarchicalDataSource({
  //      transport: {
  //        read: {
  //          url: "https://demos.telerik.com/kendo-ui/service/Employees",
  //          dataType: "jsonp"
  //        }
  //      },
  //      schema: {
  //        model: {
  //          id: "EmployeeId",
  //          hasChildren: "HasEmployees"
  //        }
  //      }
  //    });
  //
  //    $("#panelbar").kendoPanelBar({
  //      dataSource: dataSource,
  //      dataTextField: "FullName"
  //    });
  //    </script>
  DataSource interface{} `jsObject:"dataSource" jsType:"KendoDataSource,string,*map[string]interface {},[]string"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/dataspritecssclassfield

  Sets the field of the data item that provides the sprite CSS class of the nodes. If an array, each level uses the
  field that is at the same index in the array, or the last item in the array. (default: null)
  */
  //    Example
  //
  //    <style>
  //        #panelbar .k-sprite {
  //            background-image: url("https://demos.telerik.com/kendo-ui/content/shared/styles/flags.png");
  //        }
  //    </style>
  //    <div id="panelbar"></div>
  //    <script>
  //    var items = [
  //      { text: "Brazil", sprite: "brazilFlag" },
  //      { text: "India", sprite: "indiaFlag" }
  //    ];
  //    $("#panelbar").kendoPanelBar({
  //      dataSpriteCssClassField: "sprite",
  //      dataSource: items
  //    });
  //    </script>
  DataSpriteCssClassField string `jsObject:"dataSpriteCssClassField"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/datatextfield

  Sets the field of the data item that provides the text content of the nodes. If an array, each level uses the field
  that is at the same index in the array, or the last item in the array.
  */
  //    Example
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    var items = [
  //      { ProductName: "Tea", items: [
  //        { ProductName: "Green Tea" },
  //        { ProductName: "Black Tea" }
  //      ] },
  //      { ProductName: "Coffee" }
  //    ];
  //    $("#panelbar").kendoPanelBar({
  //      dataTextField: "ProductName",
  //      dataSource: items
  //    });
  //    </script>
  //
  //
  //    Example - using different fields on different levels
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    var items = [
  //      { CategoryName: "Tea", items: [
  //        { ProductName: "Green Tea" },
  //        { ProductName: "Black Tea" }
  //      ] },
  //      { CategoryName: "Coffee" }
  //    ];
  //    $("#panelbar").kendoPanelBar({
  //      dataTextField: [ "CategoryName", "ProductName" ],
  //      dataSource: items
  //    });
  //    </script>
  DataTextField []string `jsObject:"dataTextField"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/dataurlfield

  Sets the field of the data item that provides the link URL of the nodes. (default: null)
  */
  //    Example
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    var items = [
  //      { text: "Tea", LinksTo: "http://tea.example.com" },
  //      { text: "Coffee", LinksTo: "http://coffee.example.com" }
  //    ];
  //    $("#panelbar").kendoPanelBar({
  //      dataUrlField: "LinksTo",
  //      dataSource: items
  //    });
  //    </script>
  DataUrlField string `jsObject:"dataUrlField"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/expandmode

  Specifies how the PanelBar items are displayed when opened and closed. The following values are available:
  (default: "multiple")

  "single" - Display one item at a time when an item is opened; opening an item will close the previously opened item.

  "multiple" - Display multiple values at one time; opening an item has no visual impact on any other items in the
  PanelBar.
  */
  //    Example
  //
  //    <ul id="panelbar">
  //        <li>Item 1
  //            <ul>
  //                <li>Sub Item 1</li>
  //                <li>Sub Item 2</li>
  //                <li>Sub Item 3</li>
  //            </ul>
  //        </li>
  //        <li>Item 2
  //            <ul>
  //                <li>Sub Item 1</li>
  //                <li>Sub Item 2</li>
  //                <li>Sub Item 3</li>
  //            </ul>
  //        </li>
  //    </ul>
  //    <script>
  //        $("#panelbar").kendoPanelBar({
  //            expandMode: "single"
  //        });
  //    </script>
  ExpandMode TypeExpandMode `jsObject:"expandMode"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/loadondemand

  Indicates whether the child DataSources should be fetched lazily when parent groups get expanded. Setting this to
  false causes all child DataSources to be loaded at initialization time.
  */
  //    Example - force lazy loading of sublevels
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    $("#panelbar").kendoPanelBar({
  //      loadOnDemand: true,
  //      dataSource: [
  //        { text: "foo", items: [
  //          { text: "bar" }
  //        ] }
  //      ]
  //    });
  //    </script>
  LoadOnDemand Boolean `jsObject:"loadOnDemand"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/messages

  The text messages displayed in the widget. Use it to customize or localize the messages.
  */
  //    Example - customize TreeView messages
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    $("#panelbar").kendoPanelBar({
  //      dataSource: {
  //        transport: {
  //          read: function(options) {
  //            // request always fails after 1s
  //            setTimeout(function() {
  //              options.error({});
  //            }, 1000);
  //          }
  //        }
  //      },
  //      messages: {
  //        retry: "Wiederholen",
  //        requestFailed: "Anforderung fehlgeschlagen.",
  //        loading: "Laden..."
  //      }
  //    });
  //    </script>
  Messages PanelBarMessages `jsObject:"messages"`

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/panelbar/configuration/template

  Template for rendering each node.
  */
  //    Example - customize TreeView messages
  //
  //    <div id="panelbar"></div>
  //    <script>
  //    $("#panelbar").kendoPanelBar({
  //      template: "#= item.text # (#= item.inStock #)",
  //      dataSource: [
  //        { text: "foo", inStock: 7, items: [
  //          { text: "bar", inStock: 2 },
  //          { text: "baz", inStock: 5 }
  //        ] }
  //      ]
  //    });
  //    </script>
  Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

  *ToJavaScriptConverter
}
func(el *KendoUiPanelBar) ToJavaScript() []byte {
  var ret bytes.Buffer

  if el.Html.Global.Id == "" {
    el.Html.Global.Id = GetAutoId()
  }

  element := reflect.ValueOf(el).Elem()
  data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
  if err != nil {
    log.Criticalf( "KendoUiPanelBar.Error: %v", err.Error() )
    return []byte{}
  }

  ret.Write( []byte(`$("#` + el.Html.Global.Id + `").kendoPanelBar({`) )
  ret.Write( data )
  ret.Write( []byte(`});`) )
  ret.Write( []byte{ 0x0A } )

  return ret.Bytes()
}
func(el *KendoUiPanelBar) ToHtml() []byte{
  return el.Html.ToHtml()
}
func(el *KendoUiPanelBar) GetId() []byte{
  if el.Html.Global.Id == "" {
    el.Html.Global.Id = GetAutoId()
  }
  return []byte( el.Html.Global.Id )
}
func(el *KendoUiPanelBar) GetName() []byte{
  if el.Html.Name == "" {
    el.Html.Name = GetAutoId()
  }
  return []byte( el.Html.Name )
}