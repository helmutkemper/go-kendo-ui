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

  dataSource := KendoDataSource{
    VarName: "exposedPortsDataSource",
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
  }

  dialogWindow := HtmlElementScript{
    Global: HtmlGlobalAttributes{
      Id: "containerCreateTemplateExposedPortsAddNewPort",
    },
    Type: SCRIPT_TYPE_KENDO_TEMPLATE,
    Content: Content{

      HtmlElementForm{
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
                  Name: "ExposedPortsNumber",
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
      },

    },
  }

  content := HtmlContent{
    Content: Content{

      dataSource,

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
        DataSource: dataSource,
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
              Code: `function(input){ 
                      if(!$('#spanCreateTemplateExposedPortsAddNewPort').kendoValidator().data('kendoValidator').validate())
                      { 
                        return false; 
                      }

                      dataSource.add({
                        Id: dataSource.total(),
                        Value: $('#ExposedPortsNumber').val() + '/' + $('#ExposedPortsProtocol').val(),
                        ImageName: 'fixme: containerConfigurationImageNameRef.text()'
                      });

                      dataSource.one('requestEnd', function(args) {
                        if (args.type !== 'create') {
                          return;
                        }

                        dataSource.one('sync', function() {
                          //$('#containerHostExposedPorts').value($('#containerHostExposedPorts').value().concat([containerHostExposedPortsItemsIdToAdd]));
                        });

                        $('#ExposedPortsProtocol').data('kendoComboBox').val('')
                        $('#ExposedPortsNumber').data('kendoNumericTextBox').value('');
                      });

                      dataSource.sync();
                      return false; 
                    }`,
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

  fmt.Printf( "%s\n\n", content.ToJavaScript() )
  fmt.Printf( "%s", content.ToHtml() )

  // Output:
  //
}
func ExampleSoUmTest() {
  el :=       HtmlElementForm{
    Global: HtmlGlobalAttributes{
      Id: "spanCreateTemplateExposedPortsAddNewPort",
    },
    Content: Content{

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigHostName",
            Content: Content{
              "Host Name",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigHostName",
              Class: "k-textbox",
            },
            Name: "HostName",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigDomainName",
            Content: Content{
              "Domain Name",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigDomainName",
              Class: "k-textbox",
            },
            Name: "DomainName",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigUser",
            Content: Content{
              "User",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigUser",
              Class: "k-textbox",
            },
            Name: "User",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigAttachStdIn",
            Content: Content{
              "Attach Std In",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigAttachStdIn",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "AttachStdIn",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigAttachStdOut",
            Content: Content{
              "Attach Std Out",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigAttachStdOut",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "AttachStdOut",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigAttachStdErr",
            Content: Content{
              "Host Attach Std Err",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigAttachStdErr",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "AttachStdErr",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigExposedPorts",
            Content: Content{
              "Exposed Ports",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigTry",
            Content: Content{
              "Try",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigTry",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "Try",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigOpenStdIn",
            Content: Content{
              "Open Std In",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigOpenStdIn",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "OpenStdIn",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigStdInOnce",
            Content: Content{
              "Std In Once",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigStdInOnce",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "StdInOnce",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigEnv",
            Content: Content{
              "Env",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigCmd",
            Content: Content{
              "Cmd",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigHealthCheck",
            Content: Content{
              "Health Check",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigArgsEscaped",
            Content: Content{
              "Args Escaped",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigArgsEscaped",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "ArgsEscaped",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigImage",
            Content: Content{
              "Image",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigHostImage",
              Class: "k-textbox",
            },
            Name: "Image",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigVolumes",
            Content: Content{
              "Volumes",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigWorkingDir",
            Content: Content{
              "Working Dir",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigWorkingDir",
              Class: "k-textbox",
            },
            Name: "WorkingDir",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigEntryPoint",
            Content: Content{
              "EntryPoint",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigNetworkDisabled",
            Content: Content{
              "Network Disabled",
            },
          },

          HtmlInputCheckBox{
            Global: HtmlGlobalAttributes{
              Id: "ConfigNetworkDisabled",
              Data: map[string]string{
                "role": "switch",
                "off-label": "No",
                "on-label": "Yes",
              },
            },
            Name: "NetworkDisabled",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigMacAddress",
            Content: Content{
              "Mac Address",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigMacAddress",
              Class: "k-textbox",
            },
            Name: "MacAddress",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigOnBuild",
            Content: Content{
              "On Build",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigLabels",
            Content: Content{
              "Labels",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigStopSignal",
            Content: Content{
              "Stop Signal",
            },
          },

          HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigStopSignal",
              Class: "k-textbox",
            },
            Name: "StopSignal",
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigStopTimeout",
            Content: Content{
              "Stop Timeout",
            },
          },

        },
      },

      HtmlElementDiv{
        Content: Content{

          HtmlElementFormLabel{
            For: "ConfigShell",
            Content: Content{
              "Shell",
            },
          },

        },
      },














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
              Name: "ExposedPortsNumber",
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
              Name: "ExposedPortsProtocol",
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
  }

  /*
  desejado:
  {
    "Docker": {
      "AttachStdin": true,
      "AttachStdout": true,
      "AttachStderr": true,
      "ExposedPorts": {
        "27018/tcp": {}
      },
      "OpenStdin": true,
      "Cmd": [
        "bash"
      ],
      "Image": "mongo:latest"
    },
    "Host": {
      "PortBindings": {
        "27017/tcp": [
          {
            "HostPort": "27018"
          }
        ]
      }
    },
    "Name": "mongo_test"
  }

  formato:
  {
    "root": [
      "Docker", "AttachStdin"
    ],
    "types": [
      {
        "name": "Docker",
        "type": "object",
        "attributes": [ "ExposedPorts" ]
      },
      {
        "name": "ExposedPortKey",
        "type": "concat",
        "attributes": [ "id:ExposedPortNumber", "/", "id:ExposedPortProtocol" ]
      },
      {
        "name": "ExposedPortValue",
        "type": "object"
      },
      {
        "name": "AttachStdin",
        "type": "boolean",
        "souce": {
          "value": "id:AttachStdin"
        }
      },
      {
        "name": "ExposedPorts",
        "type": "object",
        "source": {
          "key": "name:ExposedPortKey",
          "value": "name:ExposedPortValue"
        }
      }
    }
  ]
  */

  //var obj = make( map[string]interface{} )
  var key, jsCode string
  // fixme: mfalta um getName() ou algo parecido
  // fixme: KendoUiCalendar e KendoUiColorPalette devem ter name como obrigatórios
  fmt.Printf("var obj = {\n")
  for _, v := range el.MakeJavaScript(){
    switch converted := v.(type) {
    case *KendoUiAutoComplete:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoAutoComplete').value()`
    case *KendoUiButton:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoButton').value()`
    case *KendoUiCalendar:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoCalendar').value()`
    case *KendoUiColorPalette:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoColorPalette').value()`
    case *KendoUiColorPicker:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoColorPicker').value()`
    case *KendoUiComboBox:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoComboBox').value()`
    case *KendoUiNumericTextBox:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoNumericTextBox').value()`
    case *KendoUiDateInput:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoDateInput').value()`
    case *KendoUiDatePicker:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoDatePicker').value()`
    case *KendoUiDateTimePicker:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoDateTimePicker').value()`
    case *KendoUiDropDownList:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoDropDownList').value()`
    case *KendoUiMultiSelect:
      key = converted.Html.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').data('kendoMultiSelect').value()`
    case *HtmlElementFormSelect:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlElementFormTextArea:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputCheckBox:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputColor:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputDate:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputDateTimeLocal:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputEmail:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputFile:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputGeneric:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputHidden:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputImage:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputMonth:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputNumber:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputPassword:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputRadio:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputRange:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputSearch:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputTel:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputText:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputTime:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputUrl:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    case *HtmlInputWeek:
      key = converted.Name
      jsCode = `$('#` + string( converted.GetId() ) + `').value()`
    }
    //fmt.Printf( "%T\n", v )
    fmt.Printf( "  'id:%v': %s,\n", key, jsCode )
  }
  fmt.Printf( "};\n" )

  // Output:
  //
}