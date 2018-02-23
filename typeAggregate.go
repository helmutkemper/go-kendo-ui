package telerik

type TypeAggregate int

func (el TypeAggregate) IsSet() bool {
  if el != 0 {
    return true
  }
  return false
}

func (el TypeAggregate) String() string {
  return typeAggregates[el]
}

var typeAggregates = [...]string{
  "",
  "average",
  "count",
  "max",
  "min",
  "sum",
}

const (
  //Only for Number.
  AGGREGATE_AVERAGE TypeAggregate = 1
  //String, Number and Date.
  AGGREGATE_COUNT
  //Number and Date.
  AGGREGATE_MAX
  //Number and Date.
  AGGREGATE_MIN
  //Only for Number.
  AGGREGATE_SUM
)