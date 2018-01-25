package telerik

type StringArr []string

func(el StringArr) ToAttr(name string) string {
  var ret string

  for k, v := range el{
    if k != 0 {
      ret += `, `
    }
    ret += v
  }

  return ` ` + name + `="` + ret + `" `
}
