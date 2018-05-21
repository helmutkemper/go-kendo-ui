package telerik

import (
  "bytes"
  "strconv"
)

type AceEditor struct {
  Html                                  HtmlElementDiv                          `jsObject:"-"`
  Theme                                 AceTheme
  Mode                                  AceMode
  Content                               string
  TabSize                               int
  TabSoft                               Boolean
  FontSize                              int
  WarpMode                              Boolean
  HighlightActiveLine                   Boolean
  ShowPrintMargin                       Boolean
  ReadOnly                              Boolean
}

func(el *AceEditor) ToJavaScript() []byte {
  var ret bytes.Buffer

  if el.Html.Global.Id == "" {
    el.Html.Global.Id = GetAutoId()
  }

  var name = el.GetName()

  ret.Write( []byte(` var `) )
  ret.Write( name )
  ret.Write( []byte(` = `) )
  ret.Write( []byte(`ace.edit("` + el.Html.Global.Id + `");`) )

  if el.Theme != 0 {
    ret.Write( name )
    ret.Write( []byte( `.setTheme("` + el.Theme.String() + `");` ) )
  }

  if el.Mode != 0 {
    ret.Write( name )
    ret.Write( []byte( `.session.setMode("` + el.Mode.String() + `");` ) )
  }

  if len( el.Content ) != 0 {
    ret.Write( name )
    ret.Write( []byte( `.setValue("` + el.Content + `");` ) )
  }

  if el.TabSoft != 0 {
    ret.Write( name )
    ret.Write( []byte( `.getSession().setUseSoftTabs(` + el.TabSoft.String() + `);` ) )
  }

  if el.TabSize != 0 {
    ret.Write( name )
    ret.Write( []byte( `.getSession().setTabSize(` + strconv.Itoa( el.TabSize ) + `);` ) )
  }

  if el.FontSize != 0 {
    ret.Write( []byte( `document.getElementById("` ) )
    ret.Write( name )
    ret.Write( []byte( `").style.fontSize="` + strconv.Itoa( el.FontSize ) + `px";` ) )
  }

  if el.WarpMode != 0 {
    ret.Write( name )
    ret.Write( []byte( `.getSession().setUseWrapMode(` + el.WarpMode.String() + `);` ) )
  }

  if el.HighlightActiveLine != 0 {
    ret.Write( name )
    ret.Write( []byte( `.setHighlightActiveLine(` + el.HighlightActiveLine.String() + `);` ) )
  }

  if el.ShowPrintMargin != 0 {
    ret.Write( name )
    ret.Write( []byte( `.setShowPrintMargin(` + el.ShowPrintMargin.String() + `);` ) )
  }

  if el.ReadOnly != 0 {
    ret.Write( name )
    ret.Write( []byte( `.setReadOnly(` + el.ReadOnly.String() + `);` ) )
  }

  ret.Write( []byte{ 0x0A } )

  return ret.Bytes()
}

func(el *AceEditor) ToHtml() []byte {
  return el.Html.ToHtml()
}
func(el *AceEditor) GetId() []byte{
  if el.Html.Global.Id == "" {
    el.Html.Global.Id = GetAutoId()
  }
  return []byte( el.Html.Global.Id )
}
func(el *AceEditor) GetName() []byte{
  if el.Html.Name == "" {
    el.Html.Name = GetAutoId()
  }
  return []byte( el.Html.Name )
}
