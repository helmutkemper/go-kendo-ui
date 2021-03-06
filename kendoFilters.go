package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type kendoFilters struct {
	Field    string        `jsObject:"field"`
	Operator KendoOperator `jsObject:"operator"`
	Value    string        `jsObject:"value"`

	*ToJavaScriptConverter
}

func (el *kendoFilters) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoFilters.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
