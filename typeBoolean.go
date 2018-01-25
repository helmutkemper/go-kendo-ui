package telerik

type Boolean int

func (el Boolean) IsSet() bool {
  if el != 0 {
    return true
  }
  return false
}

func (el Boolean) String() string {
  if el == 1 {
    return "true"
  }

  return "false"
}

func (el Boolean) onOff() string {
  if el == 1 {
    return "on"
  }

  return "off"
}

func(el Boolean) ToAttr(name string) string {
  var ret string
  if el != 0 {
    ret = ` ` + name + `=` + el.String() + ` `
  }

  return ret
}

func(el Boolean) ToAttrOnOff(name string) string {
  var ret string
  if el != 0 {
    ret = ` ` + name + `="` + el.onOff() + `" `
  }

  return ret
}

func(el Boolean) ToAttrSet(name string) string {
  var ret string
  if el == 1 {
    ret = ` ` + name + ` `
  }

  return ret
}

const (
  FALSE Boolean = -1
  NOT_SET = 0
  TRUE = 1
)