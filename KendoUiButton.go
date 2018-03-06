package telerik

import (
  "fmt"
  "html/template"
  "bytes"
)

type KendoUiButton struct{
  HtmlId                                  string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-enable

  Indicates whether the <strong>Button</strong> should be enabled or disabled. By default, it is enabled, unless a <b><u>disabled="disabled"</u></b> attribute is detected.

  Example
   <button id="button" type="button">Foo</button>
   <script>
   $("#button").kendoButton({
       enable: false
   });
   </script>
  */

  Enable                                  Boolean

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-icon

  Defines a name of an existing icon in the Kendo UI theme sprite. The icon will be applied as background image of a <b><u>span</u></b> element inside the <strong>Button</strong>. The <b><u>span</u></b> element can be added automatically by the widget, or an existing element can be used, if it has a <b><u>k-icon</u></b> CSS class applied. For a list of available icon names, please refer to the <a href="http://demos.telerik.com/kendo-ui/web/styling/icons.html">Icons demo</a>.

  Example
   <button id="button" type="button">Cancel</button>
   <script>
   $("#button").kendoButton({
       icon: "cancel"
   });
   </script>
  */

  Icon                                    string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-imageUrl

  Defines a URL, which will be used for an <b><u>img</u></b> element inside the Button. The URL can be relative or absolute. In case it is relative, it will be evaluated with relation to the web page URL.
  The <b><u>img</u></b> element can be added automatically by the widget, or an existing element can be used, if it has a <b><u>k-image</u></b> CSS class applied.

  Example
   <button id="button" type="button">Edit</button>
   <script>
   $("#button").kendoButton({
       imageUrl: "/images/edit-icon.gif"
   });
   </script>
  */

  ImageUrl                                string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-spriteCssClass

  Defines a CSS class (or multiple classes separated by spaces), which will be used for applying a background image to a <b><u>span</u></b> element inside the <strong>Button</strong>. In case you want to use an icon from the Kendo UI theme sprite background image, it is easier to use the <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-icon"><b><u>icon</u></b> property</a>.
  The <b><u>span</u></b> element can be added automatically by the widget, or an existing element can be used, if it has a <b><u>k-sprite</u></b> CSS class applied.

  Example
   <button id="button" type="button">Edit</button>
   <script>
   $("#button").kendoButton({
       spriteCssClass: "myEditIcon"
   });
   </script>
  */

  SpriteCssClass                          string
}
func(el *KendoUiButton) IsSet() bool {
  return el != nil
}
func(el *KendoUiButton) String() string {
  var buffer bytes.Buffer
  tmpl := template.Must(template.New("").Funcs(template.FuncMap{
    "safeHTML": func(s interface{}) template.HTML {
      return template.HTML(fmt.Sprint(s))
    },
  }).Parse(GetTemplate()))
  err := tmpl.ExecuteTemplate(&buffer, "KendoUiButton", *(el))
  if err != nil {
    fmt.Println(err.Error())
  }
  
  return buffer.String()
}

