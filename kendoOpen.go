package telerik

type KendoOpen struct{
  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open.effects  The effect(s) to use when playing the open animation. Multiple effects should be separated with a space.
  <a href="/kendo-ui/api/javascript/effects/common">Complete list of available animations</a>
  
  */
  Effects                                 string

  /*
  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.open.duration  The duration of the open animation in milliseconds.
  
  */
  Duration                                int

}
func(el *KendoOpen) IsSet() bool {
  return el != nil
}

