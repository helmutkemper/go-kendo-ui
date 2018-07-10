package telerik

import "fmt"

func ExampleKendoUiPanelBar_ToHtml() {
  el := KendoUiPanelBar{
    Html: HtmlElementUl{
      Global: HtmlGlobalAttributes{
        Id: "PanelBar",
      },
    },
    Animation: KendoAnimationVertical{
      Collapse: KendoClose{
        Duration: 300,
        Effects: EFFECT_FADE_OUT,
      },
      Expand: KendoOpen{
        Duration: 300,
        Effects: EFFECT_FADE_IN,
      },
    },
    AutoBind: TRUE,
    ContentUrls: []string{
      "",
      "https://demos.telerik.com/kendo-ui/content/web/panelbar/ajax/ajaxContent1.html",
    },
    DataImageUrlField: "image",
    DataSpriteCssClassField: "sprite",
    DataTextField: []string{ "CategoryName", "ProductName" },
    DataUrlField: "LinksTo",
    ExpandMode: PANELBAR_EXPAND_MODE_SINGLE,
    LoadOnDemand: FALSE,
    Messages: PanelBarMessages{
      Loading: "Laden...",
      RequestFailed: "Sorry, error...",
      Retry: "Retry...",
    },
    Template: "#= item.text # (#= item.inStock #)",
  }

  fmt.Printf( "%s\n", el.ToHtml() )
  fmt.Printf( "%s\n", el.ToJavaScript() )

  // Output:
  // <ul id="PanelBar"></ul>
  // $("#PanelBar").kendoPanelBar({animation: { collapse: { effects: "fade:out",duration: 300,},expand: { effects: "fade:in",duration: 300,},},autoBind: true,contentUrls: ["","https://demos.telerik.com/kendo-ui/content/web/panelbar/ajax/ajaxContent1.html",],dataImageUrlField: "image",dataSpriteCssClassField: "sprite",dataTextField: ["CategoryName","ProductName",],dataUrlField: "LinksTo",expandMode: "single",loadOnDemand: false,messages: { loading: "Laden...",requestFailed: "Sorry, error...",retry: "Retry...",},template: #= item.text # (#= item.inStock #),});
}
