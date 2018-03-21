package telerik

func CustomTemplate_Demo() {
  tmpl := CustomTemplate{}
  tmpl.Id = `Port`
  tmpl.Template = `<span>Port: #= Value #<br>Usually used in: #= ImageName #</span>`
  tmpl.Footer = tmpl.GetFooterTemplateButtonAsString(`Add new port and protocol`, `centerText`)
  tmpl.NoData = `<div>No ports found to add. Please add a port and a protocol before continuing.</div>`
  tmpl.Dialog = KendoUiDialog{
    Title: `Environments vars from container`,
    Content: &JavaScript{
      Code: `kendo.template($("#containerCreateTemplateEnvVarAddNewEnvVar").html())`,
    },
    Visible: FALSE,
    Modal: TRUE,
    EventOpen: &JavaScript{
      Code: ``,
    },
    Actions: &[]KendoActions{
      {
        Text: `Close`,
      },
      {
        Text: `Add`,
        Action: &JavaScript{
          Code: `function(e){
              containerHostEnvAddNewEnvVarFunction();
              return false;
            }`,
          },
        },
        {
          Text: `Add and close`,
          Action: &JavaScript{
            Code: `function(e){
              containerHostEnvAddNewEnvVarFunction();
            }`,
          },
          Primary: TRUE,
        },
      },
    }
}

