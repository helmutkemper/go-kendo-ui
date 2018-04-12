package telerik

import "fmt"

func ExampleGetTemplate() {
  customTemplate := CustomTemplate{
    Id: []byte( `ExposePort` ),
  }
  customTemplate.Footer   = customTemplate.GetFooterTemplateButton( []byte( `Add` ), []byte( `centerText` ) )
  customTemplate.NoData   = []byte( `<div>No ports found to add. Please add a port and a protocol before continuing.</div>` )
  customTemplate.Template = []byte( `<span>Port: #= Value #<br>Usually used in: #= ImageName #</span>` )

  html := KendoUiMultiSelect{
    Html: HtmlElementFormSelect{
      Global: HtmlGlobalAttributes{
        Id: "multiSelect",
      },
    },
    AutoClose: FALSE,
    AutoWidth: TRUE,
    ClearButton: FALSE,
    DataTextField: "productName",
    DataValueField: "productId",
    Delay: 100,
    Filter: FILTER_CONTAINS,
    // FixedGroupTemplate: "Fixed header: #: data #",
    FooterTemplate: JavaScript{
      Code: string( customTemplate.GetElementFooterTemplate() ),
    },
    // GroupTemplate: "Group template: #: data #",
    NoDataTemplate: JavaScript{
      Code: string( customTemplate.GetElementNoDataTemplate() ),
    },
    Placeholder: "Select...",
    // HeaderTemplate: "<div><h2>Fruits</h2></div>",
    ItemTemplate: JavaScript{
      Code: string( customTemplate.GetElementTemplate() ),
    },
    // TagMode: TAG_MODE_MULTIPLE,
    Value: []map[string]interface{}{
      { "productName": "Item 1", "productId": "1" },
      { "productName": "Item 2", "productId": "2" },
      { "productName": "Item 3", "productId": "3" },
    },
  }

  fmt.Printf( `<!DOCTYPE html>
<html>
<head>
    <base href="https://demos.telerik.com/kendo-ui/grid/index">
    <style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
    <title></title>
    <link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />

    <script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
    <script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
    <script>
    %s
    $(document).ready(function () {
    %s
    });
    </script>
    %s
    %s
    %s

</head>
<body>

<div id="example">
    %s
</div>

<style type="text/css">
    .customer-photo {
        display: inline-block;
        width: 32px;
        height: 32px;
        border-radius: 50%;
        background-size: 32px 35px;
        background-position: center center;
        vertical-align: middle;
        line-height: 32px;
        box-shadow: inset 0 0 1px #999, inset 0 0 10px rgba(0,0,0,.2);
        margin-left: 5px;
    }

    .customer-name {
        display: inline-block;
        vertical-align: middle;
        line-height: 32px;
        padding-left: 3px;
    }
</style>


</body>
</html>
`,
    customTemplate.GetFooterTemplateButtonOpenDialogJavaScript(),
    html.ToJavaScript(),
    customTemplate.GetTemplate(),
    customTemplate.GetFooterTemplate(),
    customTemplate.GetNoDataTemplate(),
    html.ToHtml(),
  )

  // Output:
  //
}

func ExampleIdea() {

  dialogWindow := HtmlElementScript{
    Global: HtmlGlobalAttributes{
      Id: "containerCreateTemplateExposedPortsAddNewPort",
    },
    Type: SCRIPT_TYPE_KENDO_TEMPLATE,
    Content: Content{

      HtmlElementDiv{
        Global: HtmlGlobalAttributes{
          Id: "spanCreateTemplateExposedPortsAddNewPort",
        },
        Content: Content{

          HtmlElementDiv{
            Content: Content{

              HtmlElementFormLabel{
                For: "ExposedPortsNumber",
                Content: Content{
                  "Port number",
                },
              },

              KendoUiNumericTextBox{
                Html: HtmlInputNumber{
                  Name: "ExposedPorts",
                  PlaceHolder: "",
                  AutoComplete: FALSE,
                  Required: TRUE,
                  // Pattern: "[^=]*",
                  Global: HtmlGlobalAttributes{
                    Id: "ExposedPortsNumber",
                    Class: "oneThirdSize",
                    Extra: map[string]interface{}{
                      "validationMessage": "Enter a {0}",
                    },
                  },
                },
                Format: "#",
              },

            },
          },

          HtmlElementDiv{
            Content: Content{

              HtmlElementFormLabel{
                For: "ExposedPortsProtocol",
                Content: Content{
                  "Port protocol",
                },
              },

              KendoUiComboBox{
                Html: HtmlElementFormSelect{
                  Global: HtmlGlobalAttributes{
                    Id: "ExposedPortsProtocol",
                    Class: "oneThirdSize",
                  },
                  Required: TRUE,
                  Options: []HtmlOptions{
                    {
                      Label: "Please, select one",
                      Key:   "",
                    },
                    {
                      Label: "TCP",
                      Key:   "TCP",
                    },
                    {
                      Label: "UDP",
                      Key:   "UDP",
                    },
                  },
                },

              },

            },
          },
        },
      },
    },
  }

  content := HtmlContent{
    Content: Content{

      HtmlElementScript{
        Global: HtmlGlobalAttributes{
          Id: "containerHostExposedPortsFooterTemplate",
        },
        Type:    SCRIPT_TYPE_KENDO_TEMPLATE,
        Content: Content{

          HtmlElementFormButton{
            Global: HtmlGlobalAttributes{
              Id: "buttonHostExposedPortsFooterTemplate",
              Class: "k-button k-primary centerText",
              OnClick: "containerHostExposedPortsAddNewPort()",
            },
            Content: Content{
              "Add new port and protocol",
            },
          },

        },
      },

      HtmlElementScript{
        Type: SCRIPT_TYPE_JAVASCRIPT,
        Content: Content{

          `function containerHostExposedPortsAddNewPort(){ $('#containerHostExposedPortsAddNewPort').data('kendoDialog').open(); }`,

        },
      },

      HtmlElementScript{
        Global: HtmlGlobalAttributes{
          Id: "containerHostExposedPortsTemplate",
        },
        Type:    SCRIPT_TYPE_KENDO_TEMPLATE,
        Content: Content{

          HtmlElementSpan{
            Content: Content{
              "containerHostExposedPortsTemplate",
            },
          },

        },
      },

      dialogWindow,

      HtmlElementScript{
        Global: HtmlGlobalAttributes{
          Id: "containerHostExposedPortsNoDataTemplate",
        },
        Type:    SCRIPT_TYPE_KENDO_TEMPLATE,
        Content: Content{

          HtmlElementDiv{
            Global: HtmlGlobalAttributes{
              Id: "divExposedPortsFooterTemplate",
            },
            Content: Content{
              "No ports found to add. Please add a port and a protocol before continuing.",
            },
          },

        },
      },

      KendoUiMultiSelect{
        Html: HtmlElementFormSelect{
          Global: HtmlGlobalAttributes{
            Id: "containerHostExposedPorts",
          },
        },
        Filter: FILTER_CONTAINS,
        Placeholder: "",
        ItemTemplate: JavaScript{
          Code: "kendo.template($('#containerHostExposedPortsTemplate').html())",
        },
        NoDataTemplate: JavaScript{
          Code: "kendo.template($('#containerHostExposedPortsNoDataTemplate').html())",
        },
        FooterTemplate: JavaScript{
          Code: "kendo.template($('#containerHostExposedPortsFooterTemplate').html())",
        },
        DataTextField: "Value",
        DataValueField: "Id",
        DataSource: KendoDataSource{
          Schema: KendoSchema{
            Model: KendoDataModel{
              Id: "Id",
              Fields: map[string]KendoField{
                "Id": {
                  Type: JAVASCRIPT_NUMBER,
                },
                "Value": {
                  Type: JAVASCRIPT_STRING,
                },
                "ImageName": {
                  Type: JAVASCRIPT_STRING,
                },
              },
            },
          },
        },
      },

      KendoUiDialog{
        Html: HtmlElementDiv{
          Global: HtmlGlobalAttributes{
            Id: "containerHostExposedPortsAddNewPort",
          },
        },
        Modal: TRUE,
        Visible: FALSE,
        Title: "Expose port from container",
        Width: 400,
        Content: JavaScript{
          Code: string( dialogWindow.ToKendoTemplate() ),
        },
        Actions: []KendoActions{
          {
            Text: "Close",
          },
          {
            Action:  JavaScript{
              Code: "function(input){ if(!$('#spanCreateTemplateExposedPortsAddNewPort').kendoValidator().data('kendoValidator').validate()){ return false; } return false; }",
            },
            Text:    "Add",
          },
          {
            Action:  JavaScript{
              Code: "function(input){ if(!$('#spanCreateTemplateExposedPortsAddNewPort').kendoValidator().data('kendoValidator').validate()){ return false; } }",
            },
            Primary: TRUE,
            Text:    "Add and close",
          },
        },
        EventOpen: JavaScript{
          Code: "function(){ " + string( dialogWindow.Content.ToJavaScript() ) + "}",
        },
      },

    },
  }

  fmt.Printf( "%s", content.ToJavaScript() )
  fmt.Printf( "%s", content.ToHtml() )

  // Output:
  //
}