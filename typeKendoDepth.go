package telerik

type KendoTimeDepth int

var kendoTimeDepths = [...]string {
  "",
  "month",
  "year",
  "decade",
  "century",
}

func (el KendoTimeDepth) String() string {
  return kendoTimeDepths[el]
}

const (
  // Shows the days of the month.
  TIME_DEPTH_MONTH KendoTimeDepth = iota + 1

  // Shows the months of the year.
  TIME_DEPTH_YEAR

  // Shows the years of the decade.
  TIME_DEPTH_DECADE

  // Shows the decades from the century.
  TIME_DEPTH_CENTURY
)
