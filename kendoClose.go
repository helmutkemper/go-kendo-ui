package telerik

type KendoClose struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close.effects  The effect(s) to use when playing the close animation. Multiple effects should be separated with a space.
  <a href="/kendo-ui/api/javascript/effects/common">Complete list of available animations</a>
  
  */
  Effects                                 string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close.duration  The duration of the close animation in milliseconds.
  
  */
  Duration                                int

}
func(el *KendoClose) IsSet() bool {
  return el != nil
}

