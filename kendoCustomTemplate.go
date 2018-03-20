package telerik

type CustomTemplate struct {
  Id                                      string                                  `jsObject:"-"`
  Template                                string                                  `jsObject:"-"`
  Footer                                  string                                  `jsObject:"-"`
  NoData                                  string                                  `jsObject:"-"`
}
func(el *CustomTemplate) GetTemplate() []byte {
  return []byte(`<script id="` + el.Id + `Template" type="text/x-kendo-tmpl">` + el.Template + `</script>`)
}
func(el *CustomTemplate) GetFooterTemplate() []byte {
  return []byte(`<script id="` + el.Id + `FooterTemplate" type="text/x-kendo-tmpl">` + el.Footer + `</script>`)
}
func(el *CustomTemplate) GetNoDataTemplate() []byte {
  return []byte(`<script id="` + el.Id + `NoDataTemplate" type="text/x-kendo-tmpl">` + el.NoData + `</script>`)
}
func(el *CustomTemplate) GetFooterTemplateButton(buttonLabel, buttonCss string) []byte {
  return []byte(`<button class="k-button k-primary ` + buttonCss + `" onclick="` + el.Id + `CustomButtonClick()">` + buttonLabel + `</button>`)
}
func(el *CustomTemplate) GetFooterTemplateButtonOpenDialog() []byte {
  return []byte(`function ` + el.Id + `CustomButtonClick(){
      containerHostExposedPortsAddNewPortDialogWindowRef.kendoDialog({
        title: "Expose port from container",
        content: kendo.template($("#containerCreateTemplateExposedPortsAddNewPort").html()),
        visible: false,
        modal: true,
        close: function(){

        },
        open: function(){
          $("#ExposedPortsNumber").kendoNumericTextBox({ format: "#" });
          containerConfigurationExposedPortsNumberRef = $("#ExposedPortsNumber").data("kendoNumericTextBox");

          $("#ExposedPortsProtocol").kendoDropDownList();

          setTimeout( function(){ containerConfigurationExposedPortsNumberRef.focus(); }, 1000)
        },
        actions: [
          {
            text: "Close"
          },
          {
            text: "Add",
            action: function(e){
              ` + el.Id + `CustomButtonOnActionAndClose();
              return false;
            }
          },
          {
            text: "Add and close",
            action: function(e){
              ` + el.Id + `CustomButtonOnActionAndClose();
            },
            primary: true
          }
        ]
      });
      containerHostExposedPortsAddNewPortDialogWindowRef.data("kendoDialog").open();
    }
    function ` + el.Id + `CustomButtonOnActionAndClose(){
      let imageName = containerConfigurationImageNameRef.text();

      // Procura por um nome de container
      if( imageName == "" ) {
        $("#dialog").kendoDialog({
          modal: true,
          visible: true,
          title: "Configuration error",
          content: "Please, select an image name first.",
          width: 400,
          actions: [
            { text: "OK", primary: true, action: function(){ containerConfigurationImageNameRef.open(); } }
          ]
        });
        return;
      }

      // Procura por uma porta válida
      if( containerConfigurationExposedPortsNumberRef.value() == null ){
        $("#dialog").kendoDialog({
          modal: true,
          visible: true,
          title: "Configuration error",
          content: "Please, select a valid port number.",
          width: 400,
          actions: [
            { text: "OK", primary: true, action: function(){ setTimeout( function(){ containerConfigurationExposedPortsNumberRef.focus(); }, 1000); } }
          ]
        });
        return;
      }

      let dataSource = containerHostExposedPortsRef.dataSource;
      containerHostExposedPortsItemsCount += 1;
      containerHostExposedPortsItemsIdToAdd = containerHostExposedPortsItemsCount;

      let dataInDataSource = dataSource.data();
      let pass = true;

      for (var i in dataInDataSource) {
        if( isNaN( i ) ){
          break
        }
        i = parseInt(i);
        if (dataInDataSource[i].Value === $("#ExposedPortsNumber").val() + "/" + $("#ExposedPortsProtocol").val()) {
          containerHostExposedPortsItemsIdToAdd = dataInDataSource[i].Id;
          pass = false;
          break;
        }
      }

      if( pass === true ) {
        dataSource.add({
          Id: containerHostExposedPortsItemsCount,
          Value: $("#ExposedPortsNumber").val() + "/" + $("#ExposedPortsProtocol").val(),
          ImageName: containerConfigurationImageNameRef.text()
        });
      }

      dataSource.one("requestEnd", function(args) {
        if (args.type !== "create") {
          return;
        }

        dataSource.one("sync", function() {
          containerHostExposedPortsRef.value(containerHostExposedPortsRef.value().concat([containerHostExposedPortsItemsIdToAdd]));
        });

        $("#ExposedPortsNumber").data("kendoNumericTextBox").value("");
      });

      dataSource.sync();
    }`)
}