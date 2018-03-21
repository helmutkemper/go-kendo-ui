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
      Code: `kendo.template($("#` + tmpl.Id + `Template").html())`,
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
              ` + tmpl.Id + `CustomValueCheck();
              return false;
            }`,
          },
        },
        {
          Text: `Add and close`,
          Action: &JavaScript{
            Code: `function(e){
              ` + tmpl.Id + `CustomValueCheck();
            }`,
          },
          Primary: TRUE,
        },
      },
    }
}

